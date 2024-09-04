package user

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/define"
	"server/model"
	"server/mysqlModel"
	"server/pb"
	"server/pkg/manage"
	sysDefined "server/pkg/sysConst"
	"server/tool"
	"time"
)

func Enter(c *gin.Context, req *pb.UserControllerEnter, res *pb.UserEnterResponse) {
	tokenString := req.Token
	// 解析并验证JWT令牌
	tokenObj, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用正确的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(manage.GetConfigManage().Viper.GetString("token_secret_key")), nil
	})

	if err != nil {
		fmt.Println("token解析异常", err)
		tool.SetGameError(c, define.EnterTokenFait)
		return
	}
	// 检查Token是否有效
	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok || !tokenObj.Valid {
		fmt.Println("token验证异常 !")
		tool.SetGameError(c, define.EnterTokenFait)
		return
	}

	sId := int(claims["sId"].(float64))
	exp := int64(claims["exp"].(float64))
	accountId := int(claims["accountId"].(float64))
	if sId != manage.GetConfigManage().Viper.GetInt("server.id") {
		fmt.Println("Token数据异常!")
		tool.SetGameError(c, define.EnterTokenFait)
		return
	}
	if exp <= time.Now().Unix() {
		fmt.Println("Token过期!")
		tool.SetGameError(c, define.EnterTokenFait)
		return
	}
	var uSql mysqlModel.User
	result := manage.GetMysqlManage().Client.Where("account_id = ?", accountId).Find(&uSql)
	if result.Error != nil {
		fmt.Println("读取数据库异常!")
		tool.SetGameError(c, define.EnterTokenFait)
		return
	}
	var user *model.User
	if uSql.UserId == 0 {
		//从mysql中找不到玩家信息
		uId := manage.GetNextUserId()
		uSql.AccountId = accountId
		uSql.UserId = int(uId)
		uSql.Status = 0
		//创建玩家mysql的数据
		createResult := manage.GetMysqlManage().Client.Create(&uSql)
		if createResult.Error != nil {
			fmt.Println("写入数据库异常!")
			tool.SetGameError(c, define.EnterCreateUserFait)
			return
		}
		//创建玩家mongoDo的数据
		user, err = initMongoUser(c, uId)
		if err != nil {
			tool.SetGameError(c, define.EnterCreateUserFait)
			return
		}
		user = model.NewUser()
	} else {
		//从mysql中找到了玩家就读取mongo的数据给客户端
		user, ok = manage.GetUser(c, int64(uSql.UserId))
		if !ok {
			user, err = initMongoUser(c, int64(uSql.UserId))
			if err != nil {
				tool.SetGameError(c, define.EnterCreateUserFait)
				return
			}
		}

	}
	pbUser := &pb.User{}
	tool.StructToPb(user, pbUser)
	res.User = pbUser
	c.Set(sysDefined.ActionUser, user)
}

func initMongoUser(c *gin.Context, uId int64) (*model.User, error) {
	user := model.NewUser()
	user.Id = uId
	user.Name = fmt.Sprintf("user_%d", user.Id)
	collection := manage.GetMongoManage().GetUserDb().Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

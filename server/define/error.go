package define

const (
	ProtoDecodeFait = iota
	ActionFuncFait
	UserDataFait
	RouteError
	RequestOften
	SendToClientFait
	EnterTokenFait
	EnterCreateUserFait
	UserNoLogin
)

var ErrorMap = map[int]string{
	ProtoDecodeFait:     "proto数据异常",
	UserDataFait:        "用户数据异常",
	ActionFuncFait:      "业务方法异常",
	RouteError:          "请求路由异常",
	RequestOften:        "系统忙碌中请稍后",
	EnterTokenFait:      "token解析异常",
	EnterCreateUserFait: "用户创建失败",
	UserNoLogin:         "请先登陆",
}

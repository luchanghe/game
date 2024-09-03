package manage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
	"time"
)

type MysqlManage struct {
	Client *gorm.DB
}

var mysqlManageOnce sync.Once
var mysqlManageCache *MysqlManage

func GetMysqlManage() *MysqlManage {
	mysqlManageOnce.Do(func() {
		dsn := strings.Join([]string{
			GetConfigManage().GetString("mysql.user"),
			":",
			GetConfigManage().GetString("mysql.password"),
			"@tcp(",
			GetConfigManage().GetString("mysql.host"),
			":",
			GetConfigManage().GetString("mysql.port"),
			")/",
			GetConfigManage().GetString("mysql.db_name"),
			"?charset=utf8mb4&parseTime=True&loc=Local",
		}, "")

		mysqlManageCache = &MysqlManage{}
		var err error
		mysqlManageCache.Client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("连接数据库失败: %v", err)
		}

		rdb, err := mysqlManageCache.Client.DB()
		if err != nil {
			log.Fatalf("获取底层数据库连接失败: %v", err)
		}
		// 配置连接池
		rdb.SetMaxIdleConns(GetConfigManage().GetInt("mysql.max_free_conn"))                                   // 设置最大空闲连接数
		rdb.SetMaxOpenConns(GetConfigManage().GetInt("mysql.max_open_conn"))                                   // 设置最大打开连接数
		rdb.SetConnMaxIdleTime(time.Duration(GetConfigManage().GetInt("mysql.max_free_second")) * time.Second) // 设置连接最大空闲时间
		rdb.SetConnMaxLifetime(time.Duration(GetConfigManage().GetInt("mysql.max_conn_second")) * time.Second) // 设置连接最大生命周期
	})
	return mysqlManageCache
}

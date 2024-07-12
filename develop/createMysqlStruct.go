package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

//go:generate go run createMysqlStruct.go
func main() {
	createMysqlStruct()
}

type Column struct {
	Field   string
	Type    string
	Comment string
}

func createMysqlStruct() {
	dsn := "root:@tcp(127.0.0.1:3306)/game_s1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
		return
	}
	var columns []Column
	db.Raw("SHOW FULL COLUMNS FROM user").Scan(&columns)
	var buf = strings.Builder{}
	buf.WriteString("type User struct ")
	for _, column := range columns {
		fmt.Printf("Column: %s, Type: %s, Comment: %s\n",
			column.Field, column.Type, column.Comment)
	}
}

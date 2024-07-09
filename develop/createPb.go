package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:generate go run createPb.go
func main() {
	createPb()
}

func createPb() {
	log.Println("开始生成pb")
	path, _ := os.Getwd()
	outPath := filepath.Clean(path + "/../pb")
	protoPath := filepath.Clean(path + "/../proto/")
	err := filepath.Walk(protoPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == ".proto" {
				///opt/homebrew/bin/protoc --proto_path=/Users/zhangshaojie/go/src/game/proto --go_out=/Users/zhangshaojie/go/src/game/pb  --go_opt=Mroute.proto=../pb  route.proto
				c := exec.Command("protoc", "--proto_path="+protoPath, "--go_out="+outPath, "--go_opt=M"+info.Name()+"=../pb", info.Name())
				log.Println("执行生成pb命令:", c.String())
				err := c.Run()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal("pb生成异常", err)
		return
	}
	log.Println("pb生成结束")
}

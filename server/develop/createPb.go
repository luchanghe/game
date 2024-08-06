package develop

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CreatePb() {
	log.Println("开始生成pb")
	path, _ := os.Getwd()
	outPath := path
	protoPath := filepath.Clean(path + "/../proto/")
	err := filepath.Walk(protoPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == ".proto" {
				c := exec.Command("protoc", "--proto_path="+protoPath, "--go_out="+outPath, "--go_opt=Mbase.proto=./pb", "--go_opt=M"+info.Name()+"=./pb", info.Name())
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
		log.Println("pb生成异常", err)
		return
	}
	log.Println("pb生成结束")
}

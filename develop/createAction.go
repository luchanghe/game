package main

import (
	"fmt"
	"game/pb"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

//go:generate go run createAction.go
func main() {
	createPbMap()
}

var doc = `
package server

import (
	"errors"
%s
	"game/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func doAction(c *gin.Context, result *Data, reqRoute uint32) (proto.Message, error) {
	switch reqRoute {
	%s
	default:
		return nil, errors.New("异常的路由枚举")
	}
}
`
var docCaseHasReq = `
	case uint32(__routeEnum__):
			req := &pb.__routeReq__{}
			res := &pb.__routeRes__{}
			err := proto.Unmarshal(result.Proto, req)
			if err != nil {
				return nil, err
			}
			__funcBar__.__funcName__(c, req, res)
			return res,nil`

var docCaseNoReq = `
	case uint32(__routeEnum__):
			res := &pb.__routeRes__{}
			__funcBar__.__funcName__(c, res)
			return res,nil`
var funcDocHasReq = `
package __funcBar__

import (
	"game/pb"
	"github.com/gin-gonic/gin"
)

func __funcName__(c *gin.Context, req *pb.__routeReq__, res *pb.__routeRes__) {
}`

var funcDocNoReq = `
package __funcBar__
import (
	"game/pb"
	"github.com/gin-gonic/gin"
)
func __funcName__(c *gin.Context, res *pb.__routeRes__) {
}`

type replaceData struct {
	routeEnum string
	routeReq  string
	routeRes  string
	funcBar   string
	funcName  string
}

func createPbMap() {
	log.Print("开始生成doAction")
	slice := make([]int32, 0, len(pb.RouteMap_name))
	for i := range pb.RouteMap_name {
		slice = append(slice, i)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	var strBuf strings.Builder
	var importBuf strings.Builder
	encountered := make(map[string]any)
	path, _ := os.Getwd()
	for _, v := range slice {
		if v%2 != 0 {
			csName := pb.RouteMap_name[v]
			scName, ok := pb.RouteMap_name[v+1]
			if !ok {
				scName = "SC_DefaultResponse"
			}
			csFix := strings.Split(csName, "_")                                      // [CS,AbcController,getList]
			scFix := strings.Split(scName, "_")                                      // [SC,AbcResponse]
			upModName := strings.TrimSuffix(csFix[1], "Controller")                  // Abc
			loModName := string(unicode.ToLower(rune(upModName[0]))) + upModName[1:] // abc
			funBar := loModName
			funcName := string(unicode.ToUpper(rune(csFix[2][0]))) + csFix[2][1:]
			reqName := csFix[1] + funcName
			resName := scFix[1]
			pbFileName := funBar
			pbFilePath := filepath.Clean(path + "/../pb/" + pbFileName + ".pb.go")
			pbFile, err := os.OpenFile(pbFilePath, os.O_RDONLY, 0644)
			var docCase string
			var funcCase string
			if err != nil {
				if os.IsNotExist(err) {
					docCase = docCaseNoReq
					funcCase = funcDocNoReq
				} else {
					log.Fatalln(err)
					return
				}
			} else {
				pbContent, err := io.ReadAll(pbFile)
				if err != nil {
					log.Fatalln(err)
					return
				}
				if strings.Contains(string(pbContent), reqName) {
					docCase = docCaseHasReq
					funcCase = funcDocHasReq
				} else {
					docCase = docCaseNoReq
					funcCase = funcDocNoReq
				}
			}
			r := &replaceData{
				routeEnum: "pb.RouteMap_" + csName,
				routeReq:  reqName,
				routeRes:  resName,
				funcBar:   funBar,
				funcName:  funcName,
			}
			strBuf.WriteString(replaceDoc(docCase, r))
			if _, ok := encountered[funBar]; !ok {
				encountered[funBar] = struct{}{}
				importBuf.WriteString("	\"game/action/" + funBar + "\"\n")
			}
			err = createDir(filepath.Clean(path + "/../action/" + funBar))
			if err != nil {
				log.Fatalln(err)
				return
			}
			err = createAndWriteFile(filepath.Clean(path+"/../action/"+funBar+"/"+csFix[2]+".go"), replaceDoc(funcCase, r))
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}
	code := fmt.Sprintf(doc, importBuf.String(), strBuf.String())
	doActionPath := filepath.Clean(path + "/../pkg/server/doAction.go")
	file, err := os.OpenFile(doActionPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("写入doAction出错", err)
		}
	}(file)
	_, err = file.WriteString(code)
	if err != nil {
		log.Fatalln("写入doAction出错", err)
		return
	}
	log.Println("doAction文件创建完成")
	c := exec.Command("go", "fmt", doActionPath)
	err = c.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func replaceDoc(doc string, r *replaceData) string {
	doc = strings.ReplaceAll(doc, "__routeEnum__", r.routeEnum)
	doc = strings.ReplaceAll(doc, "__routeReq__", r.routeReq)
	doc = strings.ReplaceAll(doc, "__routeRes__", r.routeRes)
	doc = strings.ReplaceAll(doc, "__funcBar__", r.funcBar)
	doc = strings.ReplaceAll(doc, "__funcName__", r.funcName)
	return doc
}

func createAndWriteFile(path string, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		defer func() {
			err := file.Close()
			if err != nil {
				return
			}
		}()
		if err != nil {
			return err
		}
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
		c := exec.Command("go", "fmt", path)
		err = c.Run()
		if err != nil {
			return err
		}
		log.Println("创建方法文件:", path)
	}
	return nil
}

func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

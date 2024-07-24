package develop

import (
	"game/pb"
	"game/tool"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

func CreatePbJson() {
	log.Println("开始生成pb")
	// 先反射pb包生成所有pb结构体的名称
	pkgPath := "game/pb"
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax,
	}
	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		panic(err)
	}
	structNames := make(map[string]struct{})
	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if obj, ok := obj.(*types.TypeName); ok {
				if _, ok := obj.Type().Underlying().(*types.Struct); ok {
					structNames[obj.Name()] = struct{}{}
				}
			}
		}
	}
	//因为map是无序的，转为切片使用
	slice := make([]int32, 0, len(pb.RouteMap_name))
	for i := range pb.RouteMap_name {
		slice = append(slice, i)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	path, _ := os.Getwd()
	var routes []Route
	importNames := make(map[string]struct{})
	for _, v := range slice {
		if v%2 != 0 {
			csName := pb.RouteMap_name[v]
			scName, ok := pb.RouteMap_name[v+1]
			if !ok {
				scName = "SC_DefaultResponse"
			}
			csFix := strings.Split(csName, "_")                   // [CS,AbcController,getList]
			scFix := strings.Split(scName, "_")                   // [SC,AbcResponse]
			modName := strings.TrimSuffix(csFix[1], "Controller") // Abc
			loModName := tool.FirstToLower(modName)               // abc
			funBar := loModName
			funcName := strings.Join([]string{string(unicode.ToUpper(rune(csFix[2][0]))), csFix[2][1:]}, "")
			reqName := csFix[1] + funcName
			resName := scFix[1]
			if _, ok := structNames[reqName]; !ok {
				reqName = ""
			}
			if _, ok := structNames[resName]; !ok {
				resName = "DefaultResponse"
			}
			route := Route{
				Route:    "pb.RouteMap_" + csName,
				Request:  reqName,
				Response: resName,
				BarName:  funBar,
				FuncName: funcName,
			}
			routes = append(routes, route)
			importNames[funBar] = struct{}{}
			//模块目录不存在时进行创建
			modDirPath := filepath.Clean(strings.Join([]string{path, "/action/", funBar}, ""))
			if _, err := os.Stat(modDirPath); os.IsNotExist(err) {
				err := os.MkdirAll(modDirPath, 0755)
				if err != nil {
					panic(err)
				}
			}
			// 方法文件不存在时进行创建
			funcFilePath := filepath.Clean(strings.Join([]string{path, "/action/", funBar, "/" + csFix[2], ".go"}, ""))
			if _, err := os.Stat(funcFilePath); os.IsNotExist(err) {
				tmpl := template.Must(template.New("func").Parse(templateFuncText))
				file, err := os.Create(funcFilePath)
				if err != nil {
					panic(err)
				}
				err = tmpl.Execute(file, route)
				if err != nil {
					panic(err)
				}
				log.Println("生成方法文件", funcFilePath)
				err = file.Close()
				if err != nil {
					panic(err)
				}
			}
			err = tool.FmtGoCode(funcFilePath)
			if err != nil {
				panic(err)
			}
		}
	}
	data := struct {
		Routes      []Route
		ImportNames map[string]struct{}
	}{
		Routes:      routes,
		ImportNames: importNames,
	}
	tmpl := template.Must(template.New("doAction").Parse(templateText))
	doActionPath := filepath.Clean(strings.Join([]string{path, "/pkg/server/doAction.go"}, ""))
	file, err := os.Create(doActionPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
	err = tool.FmtGoCode(doActionPath)
	if err != nil {
		panic(err)
	}
}

package develop

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"strings"
)

func CreateModelToPb() {
	pkgPath := "server/model"
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax,
	}
	var protoDef strings.Builder
	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		protoDef.WriteString("syntax = \"proto3\";\n\npackage pb;\n\n")
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if obj, ok := obj.(*types.TypeName); ok {
				if structType, ok := obj.Type().Underlying().(*types.Struct); ok {
					protoDef.WriteString(fmt.Sprintf("message %s {\n", obj.Name()))
					for i := 0; i < structType.NumFields(); i++ {
						field := structType.Field(i)
						fieldType := field.Type()
						protoType := getProtobufType(fieldType)
						protoDef.WriteString(fmt.Sprintf("  %s %s = %d;\n", protoType, field.Name(), i+1))
					}
					protoDef.WriteString("}\n\n")
				}
			}
		}
	}
	path, _ := os.Getwd()
	filePath := filepath.Clean(strings.Join([]string{path, "/../proto/base.proto"}, ""))
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.WriteString(protoDef.String())
	if err != nil {
		panic("model转为base.proto异常")
		return
	}
}

func getProtobufType(typ types.Type) string {
	switch t := typ.(type) {
	case *types.Basic:
		return getBasicProtobufType(t.Kind())
	case *types.Pointer:
		return getProtobufType(t.Elem())
	case *types.Slice:
		return fmt.Sprintf("repeated %s", getProtobufType(t.Elem()))
	case *types.Map:
		return fmt.Sprintf("map<%s, %s>", getProtobufType(t.Key()), getProtobufType(t.Elem()))
	case *types.Named:
		return t.Obj().Name()
	default:
		return "unknown"
	}
}

func getBasicProtobufType(kind types.BasicKind) string {
	switch kind {
	case types.Bool:
		return "bool"
	case types.Int, types.Int8, types.Int16, types.Int32:
		return "int32"
	case types.Int64:
		return "int64"
	case types.Uint, types.Uint8, types.Uint16, types.Uint32:
		return "uint32"
	case types.Uint64:
		return "uint64"
	case types.Float32:
		return "float"
	case types.Float64:
		return "double"
	case types.String:
		return "string"
	default:
		return "unknown"
	}
}

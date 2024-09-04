package tool

import (
	"reflect"
)

func DeepCopy(src interface{}) interface{} {
	// 获取源对象的值和类型
	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	// 如果源对象是指针，获取指针指向的元素
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
		srcType = srcType.Elem()
	}

	// 创建目标对象的实例
	dstVal := reflect.New(srcType).Elem()

	// 调用辅助函数递归复制
	deepCopyRecursive(srcVal, dstVal)

	return dstVal.Addr().Interface()
}

func deepCopyRecursive(srcVal, dstVal reflect.Value) {
	switch srcVal.Kind() {
	case reflect.Ptr:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.New(srcVal.Type().Elem()))
			deepCopyRecursive(srcVal.Elem(), dstVal.Elem())
		}
	case reflect.Interface:
		if !srcVal.IsNil() {
			newVal := reflect.New(srcVal.Elem().Type()).Elem()
			deepCopyRecursive(srcVal.Elem(), newVal)
			dstVal.Set(newVal)
		}
	case reflect.Struct:
		for i := 0; i < srcVal.NumField(); i++ {
			deepCopyRecursive(srcVal.Field(i), dstVal.Field(i))
		}
	case reflect.Slice:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.MakeSlice(srcVal.Type(), srcVal.Len(), srcVal.Cap()))
			for i := 0; i < srcVal.Len(); i++ {
				deepCopyRecursive(srcVal.Index(i), dstVal.Index(i))
			}
		}
	case reflect.Map:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.MakeMap(srcVal.Type()))
			for _, key := range srcVal.MapKeys() {
				newVal := reflect.New(srcVal.MapIndex(key).Type()).Elem()
				deepCopyRecursive(srcVal.MapIndex(key), newVal)
				dstVal.SetMapIndex(key, newVal)
			}
		}
	default:
		dstVal.Set(srcVal)
	}
}

func StructToPb(src interface{}, dest interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	destVal := reflect.ValueOf(dest).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)
		if destField.IsValid() && destField.CanSet() {
			if destField.Type() == srcField.Type {
				destField.Set(srcVal.Field(i))
			}
		}
	}

}

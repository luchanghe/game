package tool

import (
	"fmt"
	"reflect"
)

func InitStruct(i interface{}) {
	refVal := reflect.ValueOf(i).Elem()
	for i := 0; i < refVal.NumField(); i++ {
		field := refVal.Field(i)
		switch field.Kind() {
		case reflect.Map:
			field.Set(reflect.MakeMap(field.Type()))
		case reflect.Slice:
			field.Set(reflect.MakeSlice(field.Type(), 0, 0))
		case reflect.Struct:
			v := reflect.ValueOf(field)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			} else {
				fmt.Println("Expected a pointer to a struct")
				return
			}
			InitStruct(v)
		}
	}
}

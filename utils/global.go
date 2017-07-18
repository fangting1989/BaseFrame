package utils

import (
	"fmt"
	"reflect"
)

type Pageobject struct {
	Pagenum  int
	Pagesize int
}

type ResultList struct {
	Errid       int
	Errmsg      string
	Data        interface{}
	Recordcount int64
}

type ResultObject struct {
	Errid  int
	Errmsg string
	Data   interface{}
}

func StructInfo(obj interface{}) string {
	val := reflect.ValueOf(obj).Elem()
	// fmt.Println("type:", t.Name())

	// if t.Kind() != reflect.Struct {
	// 	fmt.Println("err:type invalid")
	// }
	// s := t.Elem()
	// typeOfT := s.Type() //把s.Type()返回的Type对象复制给typeofT，typeofT也是一个反射
	for k := 0; k < val.NumField(); k++ {

		valueField := val.Field(k)
		typeField := val.Type().Field(k)
		tag := typeField.Tag

		// var aa = 123
		// f := val.Field(k)                                                                                                               //迭代s的各个域，注意每个域仍然是反射。
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name")) //提取了每个域的名字

		// fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	}
	return "123"
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

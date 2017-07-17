package utils

import (
	"fmt"
	"reflect"
)

type Pageobject struct {
	pagenum  int64
	pagesize int64
	// pagenum  int64 `xorm:"BIGINT(20)"`
	// pagesize int64 `xorm:"BIGINT(20)"`
}

func StructInfo(obj interface{}) string {
	t := reflect.ValueOf(obj)
	// fmt.Println("type:", t.Name())

	if t.Kind() != reflect.Struct {
		fmt.Println("err:type invalid")
	}
	// t.Elem().NumField()

	// v := reflect.ValueOf(o)
	// u := T{Id: 1001, Name: "xxx" /*, addr:"xxx"*/}
	// t := reflect.TypeOf(u)
	// v := reflect.ValueOf(u)
	s := t.Elem()
	typeOfT := s.Type() //把s.Type()返回的Type对象复制给typeofT，typeofT也是一个反射
	for k := 0; k < s.NumField(); k++ {
		// var aa = 123
		f := s.Field(k) //迭代s的各个域，注意每个域仍然是反射。
		fmt.Printf("%d: %s %s = %v\n", k,
			typeOfT.Field(k).Name, f.Type(), f.Interface()) //提取了每个域的名字

		// fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	}
	return "123"
}

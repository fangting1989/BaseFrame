package utils

import "fmt"

type Pageobject struct {
	pagenum  int64
	pagesize int64
}

func RangStruct(v interface{}) string {
	// u := T{Id: 1001, Name: "xxx" /*, addr:"xxx"*/}
	// t := reflect.TypeOf(u)
	// v := reflect.ValueOf(u)
	for k := 0; k < t.NumFiled(); k++ {
		fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	}
	return "123"
}

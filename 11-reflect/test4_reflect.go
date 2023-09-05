// @Author huzejun 2023/9/4 19:11:00
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("type : ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
}

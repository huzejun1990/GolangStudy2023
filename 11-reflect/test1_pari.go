// @Author huzejun 2023/9/4 17:36:00
package main

import "fmt"

func main() {
	var a string

	//pair<statictype:string, value:"dream">
	a = "dream"

	//pair<type:string,value:"dream">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}

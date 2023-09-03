// @Author huzejun 2023/8/18 18:58:00
package lib2

import "fmt"

// 当前lib2包提供的API
func Lib2Test() {
	fmt.Println("lib2Test()...")
}

func init() {
	fmt.Println("lib2.init()...")
}

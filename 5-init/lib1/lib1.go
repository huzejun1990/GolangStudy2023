// @Author huzejun 2023/8/18 18:58:00
package lib1

import "fmt"

// 当前lib1包提供的API
func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1.init()...")
}

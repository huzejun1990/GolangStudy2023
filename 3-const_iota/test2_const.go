// @Author huzejun 2023/8/17 19:43:00
package main

import "fmt"

// const 来定义枚举类型
const (
	//可以在const{}添加一个关键字 iota,每行的iota都会累加1,第一行的iota的默认值是0
	BEIGING  = 10 * iota //iota = 0
	SHANGHAI             // iota = 1
	SHENZHEN             // iota = 2

)

const (
	a, b = iota + 1, iota + 2 //iota
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	//常量(只读属性)
	const length int = 10

	fmt.Println("length =", length)

	//length = 100 //常量是不允许修改的.

	fmt.Println("BEIJING = ", BEIGING)
	fmt.Println("SHANGHAI", SHANGHAI)
	fmt.Println("SHENZHEN", SHENZHEN)

}

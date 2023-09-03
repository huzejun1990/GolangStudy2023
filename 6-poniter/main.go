// @Author huzejun 2023/8/22 1:43:00
package main

import "fmt"

//func swap(a int ,b int)  {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp = main;
	*pa = *pb  // main::a = main::b
	*pb = temp

}

func main() {
	var a int = 10
	var b int = 20

	// swap
	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)

}

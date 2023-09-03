// @Author huzejun 2023/8/28 17:40:00
package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化，默认值
	//slice1 := []int{1,2,3}

	//声明一个slice1是一个切片，但是并没有给slice分配空间
	var slice1 []int
	//slice1 = make([]int,3) //开辟3个空间，默认值是0

	//声明slice1 是一个切片，同时能slice分配空间，3个空间，初始化值是0
	//var slice1 []int = make([]int,3)

	//声明slice1是一个切片，同时给slice1分配空间，3个空间，初始化值是0:通过:= 推导出slice是一个切片
	//slice1 := make([]int,3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个slice1是否为0
	if slice1 != nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}

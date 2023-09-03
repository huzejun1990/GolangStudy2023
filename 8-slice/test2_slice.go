// @Author huzejun 2023/8/27 0:10:00
package main

import "fmt"

func print2Array(myArray []int) {
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100

}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	print2Array(myArray)

	fmt.Println("===========")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}

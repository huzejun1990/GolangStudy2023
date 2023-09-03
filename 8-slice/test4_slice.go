// @Author huzejun 2023/8/28 18:55:00
package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	//fmt.Printf("len = %d,slice = %v\n",len(numbers),cap(numbers),numbers)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1，numbers len = 4,[0,0,0,1],cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2，numbers len = 5, [0,0,0,1,2],cap = 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的slice 追加元素，
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, cap = %d,slice = %v\n", len(numbers), numbers)

	fmt.Println("---------------")
	var number2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number2), cap(number2), number2)
	number2 = append(number2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number2), cap(number2), number2)

}

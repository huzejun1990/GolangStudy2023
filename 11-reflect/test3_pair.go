// @Author huzejun 2023/9/4 18:07:00
package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<type:Book, value:book{}地址>
	b := &Book{}

	//r:pair<type:, value:>
	var r Reader
	// r: pair<type:Book, value:book{}地址>
	r = b

	r.ReadBook()

	var w Write
	//w:
	w = r.(Write) //此处的断言为什么会成功？ 因为w r 具体的type是一致的

	w.WriteBook()

}

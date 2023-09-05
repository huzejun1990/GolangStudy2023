// @Author huzejun 2023/9/4 17:39:00
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty: pair<type:*os.File,"/dev/tty"文件描述符>
	file := "F:/dev/ttb.txt"
	//file := "F:/dev/ttt"
	tty, err := os.OpenFile(file, os.O_RDWR, 0)
	//tty, err := os.OpenFile("F:/dev/tty.txt",os.O_RDWR,0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//r: pair<type: , value>
	var r io.Reader
	//r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	//w: pair<type: ,value:>
	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("HELLO THIS is a test123!!!\n"))

}

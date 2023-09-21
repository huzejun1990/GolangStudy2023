// @Author huzejun 2023/9/14 20:39:00
package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

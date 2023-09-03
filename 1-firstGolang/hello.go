// @Author huzejun 2023/8/16 17:59:00
package main //程序的包名
import (
	"fmt"
	"time"
)

// main函数
func main() {
	//golang中的表达式，加";", 和不加 都可以，建议不加
	fmt.Println(" hello GO!")

	time.Sleep(1 * time.Second)
}

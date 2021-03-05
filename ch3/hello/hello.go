package main

import (
	"fmt"
	// 命名导入
	myfmt "github.com/dizhu-roc/go-in-action/ch3/mylib/fmt"
)

func main() {
	fmt.Println("Hello World!!!")

	x := myfmt.Sub(1, 2)
	fmt.Println(x)
}

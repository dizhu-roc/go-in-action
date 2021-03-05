package main

import (
	_ "github.com/dizhu-roc/go-in-action/ch2/simple/matchers"
	"github.com/dizhu-roc/go-in-action/ch2/simple/search"
	"log"
	"os"
)

// 调用之前，将日志设置到标准输出
// init 在 main 之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("COVID-19")
}

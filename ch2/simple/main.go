package main

import (
	`log`
	`os`

	`github.com/dizhu-roc/go-in-action/ch2/simple/search`
)

// 调用之前，将日志设置到标准输出
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

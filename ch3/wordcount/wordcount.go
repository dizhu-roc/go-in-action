package main

import (
	"fmt"
	"github.com/dizhu-roc/go-in-action/ch3/words"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := words.CountWord(text)
	fmt.Printf("总共有 %d 个单词\n", count)
}

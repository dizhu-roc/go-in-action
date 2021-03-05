package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine来执行这个函数
// 并发的执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜搜
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入隧道
	for _, result := range searchResults {
		results <- result
	}
}

// Display 从每个单独的 goroutine 接收到结果后，在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直阻塞，直到有结果写入
	// 一旦通道被关闭，for 循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

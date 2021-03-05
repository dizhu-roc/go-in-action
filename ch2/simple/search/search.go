package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射
// 包级别的变量
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 WaitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待的处理
	// 每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 用于搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控所有工作是否都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知 Display 函数
		// 可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果，并且
	// 在最后一个结果显示完成后返回
	Display(results)
}

// Register 注册匹配器，提供给后面的程序调用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered.")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

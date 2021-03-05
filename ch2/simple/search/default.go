package search

// 默认的匹配器
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是指向一个 Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 不需要检查错误
	return feeds, err
}

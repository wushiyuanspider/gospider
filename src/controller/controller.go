package controller

import (
	"gospider/src/configure"
	"gospider/src/fetch"
	"fmt"
)

var (
	// 获取一个全局的spider
	spider *configure.Spider
	// 全局错误变量
	err error
)

// 初始化环境
func Init(filepath string, msg bool) {
	spider, err = configure.NewSpider(filepath)
	if err != nil {
		println("Miss an error, please check!\nError: ", err)
		return
	}
	if msg {
		fmt.Println("Name:      ", spider.Name)
		fmt.Println("StartURL:  ", spider.StartURL)
		fmt.Println("Depth:     ", spider.Depth)
	}
}

// 爬取网络内容的入口函数
func Run() {
	var S fetch.Searcher
	S.GetHtmlByUrl("http://blog.csdn.net/mybc724")
	S.GetURLsFromPage(spider)
	data, _ := S.GetDataFromPage("article", spider)
	fmt.Println(data)
	/* 测试getURLsFromPage
	for k, v := range S.urls {
		fmt.Printf("name: %s\n", k)
		for x, y := range v {
			fmt.Printf("\t%d: %s\n", x + 1, y)
		}
	}*/
}

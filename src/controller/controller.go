package controller

import (
	"gospider/src/configure"
	"gospider/src/fetch"
	"fmt"
)

var (
	// 获取一个全局的spider
	spider *configure.Spider
	// 获取一个全局的Searcher(无需初始化)
	searcher *fetch.Searcher
	// 全局错误变量
	err error
	// URL队列
	QURL *queue
	// 获取到的URL的种类
	urlGroup []string
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
	// 初始化URL队列
	QURL = NewQueue()

	urlGroup = make([]string, spider.NumURLGroup())
}

// 爬取网络内容的入口函数
func Run() {
	start()

	// 获取当前页面的URL
	err = searcher.GetURLsFromPage(spider)
	if err != nil {}
	urlGroup = searcher.URLGroupNames()
	fmt.Println(urlGroup)
	currGroup := prepareURL()
	fmt.Println(currGroup)

}

func start() {
	// 将初始URL放入队列中
	QURL.Put(spider.StartURL)
	// 获得StartURL的HTML
	searcher.GetHtmlByUrl(QURL.Get())

}

// 每一次调用，都将一类URL写入队列中，并将类名返回
func prepareURL() string {
	var urls []string
	for i, v := range urlGroup {
		if v != "" {
			// 实际每次只处理一组
			urls = searcher.Urls[v]
			for _, url := range urls {
				QURL.Put(url)
			}
			urlGroup[i] = ""
			return v
		}
	}
	return ""
}

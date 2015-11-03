package controller

import (
	"gospider/src/configure"
	"gospider/src/fetch"
	"fmt"
)

var (
	// 获取一个全局的spider
	spider *configure.Spider
	// 获取一个全局的Searcher
	searcher *fetch.Searcher
	// 全局错误变量
	err error
	// URL队列
	QURL *queue
	// 已经匹配到的深度
	count int = 0
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
		fmt.Println("Depth:     ", spider.Count)
	}
	// 初始化URL队列
	QURL = NewQueue()

	//urlGroup = make([]string, spider.NumURLGroup())
	searcher = new(fetch.Searcher)
}

// 爬取网络内容的入口函数
func Run() {
	start()
	
	name, url := QURL.Get()
	for ; name != "" && url != "" && count <= spider.Count; name, url = QURL.Get() {
		searcher.GetHtmlByUrl(url)
		searcher.GetURLsFromPage(spider)
		prepareURL()
		data, _ := searcher.GetDataFromPage(name, spider)
		count++
		// 打印抓取到的结果
		fmt.Println(data["info"][0][1])
	}

	fmt.Println("urls: ", QURL.Len_urls())
	fmt.Println("used_urls: ", QURL.Len_used())
}

func start() {
	// 将初始URL无需放入队列中
	// 因为规则没有匹配到的话，则后面就不会再出现了
	// 获得StartURL的HTML
	searcher.GetHtmlByUrl(spider.StartURL)
	// 获取首页上的URL
	searcher.GetURLsFromPage(spider)
	// 将获取到的url加入队列中去
	prepareURL()
}

// 将当前页面中的url写入队列中去
func prepareURL() {
	for name, urls := range searcher.Urls {
		for _, url := range urls {
			QURL.Put(name, url)
		}
	}
	
}

// [test]
// 输出测试数据
func PrintKeyData(data fetch.KeyData) {
	for _, v := range data {
		for _, x := range v {
			fmt.Println(x[1])
		}
	}
}

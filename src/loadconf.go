package main

import (
	"fmt"
	C "github.com/Unknwon/goconfig"
	"regexp"
	"time"
)

type Spider struct {
	name           string
	startURL       string
	depth          int
	urls           map[string]*regexp.Regexp
	contents       map[string]map[string]*regexp.Regexp
	matchedURL     int
	matchedContent int
	startTime      time.Time
	endTime        time.Time
}

// 创建一个新的Macker
func NewSpider(filepath string) (*Spider, error) {
	spider := new(Spider)
	// 加载配置文件
	conf, err := C.LoadConfigFile(filepath)
	if err != nil {
		return nil, err
	}

	spider.name, err = conf.GetValue("core", "name")
    if err != nil {
        return nil, err
    }
    
	spider.startURL, err = conf.GetValue("core", "startURL")
    if err != nil {
        return nil, err
    }
    
	spider.depth, err = conf.Int("core", "depth")
    if err != nil {
        return nil, err
    }

	// 获得待匹配URL的正则表达式列表
	urls, err := conf.GetSection("url")
	if err != nil {
		return nil, err
	}

    // 初始化外层的map
	spider.urls = make(map[string]*regexp.Regexp, len(urls))
	spider.contents = make(map[string]map[string]*regexp.Regexp, len(urls))
	
	// 编译用于匹配URL的正则表达式
	for k, v := range urls {
		re_url, err := regexp.Compile(v)
		if err != nil {
			fmt.Println("Compile ", v, " error: ", err)
		} else {
			// 将URL保存，并在contents中创建对应的存储map
			spider.urls[k] = re_url
			// 获取每个URL对应的待匹配的正则表达式
			sub, err := conf.GetSection("url." + k)
			if err != nil {
				return nil, fmt.Errorf("Don't find url." + k)
			}
			// 初始化里层的map
			spider.contents[k] = make(map[string]*regexp.Regexp, len(k))
			for x, y := range sub {
				re_con, err := regexp.Compile(y)
				if err != nil {
					fmt.Println("Compile ", y, "error:", err)
				} else {
					spider.contents[k][x] = re_con
				}
			} 
		}
	}

	// 为空则什么也匹配不到，没有意义
	if len(spider.urls) == 0 {
		return nil, fmt.Errorf("Nothing to fetch!")
	}

	return spider, nil
}
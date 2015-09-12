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
	urls           []*regexp.Regexp
	contents       []*regexp.Regexp
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
    
	spider.startURL, err = conf.GetValue("core", "url")
    if err != nil {
        return nil, err
    }
    
	spider.depth, err = conf.Int("core", "depth")
    if err != nil {
        return nil, err
    }

	// 获得待匹配URL的正则表达式列表
	urls, err := conf.GetSection("urls")
	if err != nil {
		return nil, err
	}

    // 获得待匹配内容的正则表达式列表
	contents, err := conf.GetSection("contents")
	if err != nil {
		return nil, err
	}

    // 编译用于匹配URL的正则表达式
	spider.urls = make([]*regexp.Regexp, len(urls))
	for _, tmp := range urls {
		re, err := regexp.Compile(tmp)
		if err != nil {
			fmt.Println("Compile ", tmp, " error: ", err)
		} else {
			spider.urls = append(spider.urls, re)
		}
	}

    // 编译用于匹配页面内容的正则表达式
	spider.contents = make([]*regexp.Regexp, len(contents))
	for _, tmp := range contents {
		re, err := regexp.Compile(tmp)
		if err != nil {
			fmt.Println("Compile ", tmp, " error: ", err)
		} else {
			spider.contents = append(spider.contents, re)
		}
	}

	// 检查用于匹配内容的slice是否为空
	// 为空则什么也匹配不到，没有意义
	if len(spider.contents) == 0 {
		return nil, fmt.Errorf("Nothing to fetch!")
	}

	return spider, nil
}
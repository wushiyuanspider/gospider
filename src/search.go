package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

// 用于当前页面爬取到的数据
type KeyData map[string][][]string

// 对应的字段都是保存当前页的值
type Searcher struct {
	// 当前页面的源代码
	html string
	// 当前页面每条URL规则所匹配到的全部URL
	urls map[string][]string
	// 当前页面的深度
	depth int
}

// 爬取网络内容的入口函数
func Run(spider *Spider) {
	var S Searcher
	S.getHtmlByUrl("http://blog.csdn.net/mybc724")
	S.getURLsFromPage(spider)
	data, _ := S.getDataFromPage("article", spider)
	fmt.Println(data)
	/* 测试getURLsFromPage
	for k, v := range S.urls {
		fmt.Printf("name: %s\n", k)
		for x, y := range v {
			fmt.Printf("\t%d: %s\n", x + 1, y)
		}
	}*/
}

// 获得指定URL的网页源代码
func (s *Searcher)getHtmlByUrl(url string) error {
	if url == "" {
		return fmt.Errorf("URL不能为空！")
	}
	
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	// 暂时只处理200的情况
	if res.StatusCode != 200 {
		return fmt.Errorf("Can't deal this code: ", res.StatusCode)
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	s.html = string(body)
	return nil
	
}

// 从页面中获取指定的URL
func (s *Searcher) getURLsFromPage(spider *Spider) error {
	if s.html == "" {
		return fmt.Errorf("Please get Html source first!")
	}
	urlnames := spider.GetURLName()
	if len(urlnames) == 0 {
		return fmt.Errorf("URL is empty!")
	}
	
	s.urls = make(map[string][]string, len(urlnames))
	for _, name := range urlnames {
		// 忽略slice中可能产生的气泡
		if name == "" {
			continue
		}
		re := spider.GetURLByName(name)
		if re == nil {
			return  fmt.Errorf("There is nothing in ", name)
		}
		s.urls[name] = re.FindAllString(s.html, -1)
	}
	
	return nil
}

func (s *Searcher) getDataFromPage(urlName string, spider *Spider) (KeyData, error) {
	if urlName == "" {
		return nil, fmt.Errorf("URL's name can't empty!")
	}
	// contentName is a slice
	contentNames := spider.GetContentNames(urlName)
	
	// 初始化用于保存结果的map
	data := make(KeyData, len(contentNames))
	for _, name := range contentNames {
		// 获取每一个名字对应的正则表达式
		re := spider.GetContentValue(urlName, name)
		// 将匹配到的内容匹配到对应的名称下
		data[name] = re.FindAllStringSubmatch(s.html, -1)
	}
	
	return data, nil
}
package fetch

import (
	"fmt"
	"gospider/src/configure"
	"io/ioutil"
	"net/http"
)

// 用于当前页面爬取到的数据
type KeyData map[string][][]string

// 对应的字段都是保存当前页的值
type Searcher struct {
	// 当前页面的源代码
	Html string
	// 当前页面每条URL规则所匹配到的全部URL
	Urls map[string][]string
	// 当前页面的深度
	Depth int
}

// 获得指定URL的网页源代码
func (s *Searcher) GetHtmlByUrl(url string) error {
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

	s.Html = string(body)
	return nil

}

// 从页面中获取指定的URL
func (s *Searcher) GetURLsFromPage(spider *configure.Spider) error {
	if s.Html == "" {
		return fmt.Errorf("Please get Html source first!")
	}
	urlnames := spider.GetURLName()
	if len(urlnames) == 0 {
		return fmt.Errorf("URL is empty!")
	}

	s.Urls = make(map[string][]string, len(urlnames))
	for _, name := range urlnames {
		// 忽略slice中可能产生的气泡
		if name == "" {
			continue
		}
		re := spider.GetURLByName(name)
		if re == nil {
			return fmt.Errorf("There is nothing in ", name)
		}
		s.Urls[name] = re.FindAllString(s.Html, -1)
	}

	return nil
}

func (s *Searcher) GetDataFromPage(urlName string, spider *configure.Spider) (KeyData, error) {
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
		data[name] = re.FindAllStringSubmatch(s.Html, -1)
	}

	return data, nil
}

// 返回URL分组
func (s *Searcher) URLGroupNames() []string {
	groups := make([]string, len(s.Urls))
	i := 0
	for k, _ := range s.Urls {
		groups[i] = k
		i++
	}

	return groups
}

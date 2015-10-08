package configure

import (
	"fmt"
	C "github.com/Unknwon/goconfig"
	"regexp"
	"time"
)

type Spider struct {
	Name           string
	StartURL       string
	Depth          int
	urls           map[string]*regexp.Regexp
	contents       map[string]map[string]*regexp.Regexp
	MatchedURL     int
	MatchedContent int
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

	spider.Name, err = conf.GetValue("core", "name")
	if err != nil {
		return nil, err
	}

	spider.StartURL, err = conf.GetValue("core", "startURL")
	if err != nil {
		return nil, err
	}

	spider.Depth, err = conf.Int("core", "depth")
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

func (s *Spider) GetURLName() []string {
	// 注意这里创建了指定长度的空slice
	// 千万不能使用append向slice中追加元素
	// 因为append是向slice末尾添加元素，而不会像空slice中填元素
	names := make([]string, len(s.urls))
	var i int = 0
	for key, _ := range s.urls {
		names[i] = key
		i++
	}

	return names
}

func (s *Spider) GetURLByName(name string) *regexp.Regexp {
	return s.urls[name]
}

// 给定一个URL Name，返回待搜索内容正则表达式的名字
func (s *Spider) GetContentNames(urlName string) []string {
	// con is map
	con := s.contents[urlName]
	if con == nil {
		return nil
	}
	// 创建指定URL Name对应的content slice
	names := make([]string, len(con))
	var i int = 0
	for name, _ := range con {
		names[i] = name
		i++
	}

	return names
}

func (s *Spider) GetContentValue(urlName, contentName string) *regexp.Regexp {
	return s.contents[urlName][contentName]
}

// 返回URL有多少组
func (s *Spider) NumURLGroup() int { return len(s.urls) }

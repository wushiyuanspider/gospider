package controller

import "testing"
import "fmt"

func Test_Queue(t *testing.T) {
	Q := NewQueue()

	if Q.Len_urls() == 0 && Q.Len_used() == 0 {
		t.Log("  NewQueue success!")
	} else {
		t.Error("  NewQueue error!")
	}

	Q.Put("a", "aaa")
	if Q.IsExist("aaa") && Q.Len_urls() == 1 && Q.Len_used() == 0 {
		t.Log("  Put and IsExist success!")
	} else {
		t.Error("  Put or IsExist error!")
	}

	_, v := Q.Get()
	if  v == "aaa" && Q.Len_urls() == 0 && Q.Len_used() == 1 {
		t.Log("  Get success!")
	} else {
		t.Error("  Get error!")
		fmt.Println(Q.IsExist("aaa"))
		fmt.Println(v)
		fmt.Println(Q.Len_urls())
	}

}

func PrintQueue(Q *queue) {
	urls := Q.urls
	used := Q.used

	for url := urls.Front(); url != nil; url = url.Next() {
		if u, ok := url.Value.(*URL); ok {
			fmt.Print(u.name, " -- ", u.url, " | ")
		}
	}

	fmt.Println()

	for used_url := used.Front(); used_url != nil; used_url = used_url.Next() {
		fmt.Print(used_url)
	}
}

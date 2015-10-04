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

	Q.Put("abc")
	if Q.IsExist("abc") && Q.Len_urls() == 1 && Q.Len_used() == 0 {
		t.Log("  Put and IsExist success!")
	} else {
		t.Error("  Put or IsExist error!")
	}

	v := Q.Get()
	if  v == "abc" && Q.Len_urls() == 0 && Q.Len_used() == 1 {
		t.Log("  Get success!")
	} else {
		t.Error("  Get error!")
		fmt.Println(Q.IsExist("abc"))
		fmt.Println(v)
		fmt.Println(Q.Len_urls())
	}

}

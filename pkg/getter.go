package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func GetJson(url string, ch chan map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("ERROR ")
		fmt.Println(err)
	}
	var target map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil || target == nil || len(target) <= 0 {
		return
	}
	ch <- target
}

//
//innerChan := make(chan struct{}, 2)
//for _, u := range urls {
//innerChan <- struct{}{}
//go func(chan struct{}) {
//resp, err := http.Get(u)
//
//if err != nil {
//fmt.Print("ERROR ")
//fmt.Println(err)
//}
//
//var target map[string]interface{}
//err = json.NewDecoder(resp.Body).Decode(&target)
//
//if err != nil || target == nil || len(target) <= 0 {
//<-innerChan
//return
//}
//ch <- target
//<-innerChan
//}(innerChan)
//}
//close(ch)

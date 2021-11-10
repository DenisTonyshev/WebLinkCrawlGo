package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetterInterface interface {
	Get(urls []string, ch chan map[string]interface{})
}

func GetJson(urls []string, ch chan map[string]interface{}) {
	for _, u := range urls {
		resp, err := http.Get(u)

		if err != nil {
			fmt.Print("ERROR ")
			fmt.Println(err)
		}

		var target map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&target)

		if err != nil || target == nil || len(target) <= 0 {
			continue
		}
		ch <- target
	}
	close(ch)
}

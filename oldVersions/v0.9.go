package oldVersions
//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/url"
//	"strings"
//	"sync"
//)
//
//var CrawledURLs = make(map[string]bool)
//var mu sync.Mutex
//
//func getJson(urls []string, ch chan map[string]interface{}) {
//	for _, u := range urls {
//			resp, err := http.Get(u)
//			if err != nil {
//			}
//			var target map[string]interface{}
//			err = json.NewDecoder(resp.Body).Decode(&target)
//			if err != nil {
//			}
//			ch <- target
//	}
//	close(ch)
//}
//
//func parseJsons(ch chan map[string]interface{}) {
//	for {
//		val := <-ch
//		var urls []string
//		urls = parseJson(val, urls, 20)
//		if urls != nil {
//			go getJson(urls, ch)
//		}
//		return
//	}
//}
//
//func parseJson(res map[string]interface{}, urls []string, depth int, ) (urls1 []string) {
//	if depth <= 0 {
//		return
//	}
//	for _, v := range res {
//		if jsonMap, ok := v.(map[string]interface{}); ok {
//			u1 := parseJson(jsonMap, urls, depth-1)
//			urls1 = append(urls1, u1...)
//			if urls1 == nil {
//				continue
//			}
//		}
//
//		if arrayMap, ok := v.([]interface{}); ok {
//			for _, value := range arrayMap {
//				if arrayMap, ok := value.(map[string]interface{}); ok {
//					u1 := parseJson(arrayMap, urls, depth-1)
//					urls1 = append(urls1, u1...)
//				}
//			}
//		}
//
//		var strVal string
//		var ok bool
//
//		if strVal, ok = v.(string); !ok {
//			continue
//		}
//
//		_, err := url.ParseRequestURI(strVal)
//		if (err != nil || !strings.HasPrefix(strVal, "https")) || strings.HasSuffix(strVal, "png") || strings.HasSuffix(strVal, "jpeg") {
//			continue
//		}
//
//		mu.Lock()
//		if !CrawledURLs[strVal] {
//			CrawledURLs[strVal] = true
//		} else {
//			mu.Unlock()
//			return
//		}
//		mu.Unlock()
//
//		urls1 = append(urls1, strVal)
//	}
//
//	return urls1
//}
//
//func Start() {
//	urls := []string{"https://l1-1.anzu.io/ios/1/baadc6f86730c7fcfbe5f5c3"}
//	ch := make(chan map[string]interface{}, 2)
//	go getJson(urls, ch)
//	parseJsons(ch)
//}
//
//func main() {
//	Start()
//	for k, v := range CrawledURLs {
//		fmt.Println("KEY "+k, v)
//	}
//}
//
//// 0 -> pkg (JSON {a:url, b: url}) -> parse (URL [url1, url2]) -> pkg (JSON {a:url, b: url})

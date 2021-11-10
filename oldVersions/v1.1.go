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
//var StartingUrl = "https://l1-1.anzu.io/ios/1/baadc6f86730c7fcfbe5f5c3"
//
////var StartingUrl = "https://dynamic-l1.anzuinfra.com/config/ios/1/baadc6f86730c7fcfbe5f5c3?"
//
//func main() {
//	CrawledURLs[StartingUrl] = true
//
//	crawlMap := getCrawlMap(StartingUrl)
//	crawlUrlPlain(crawlMap, 4)
//	for k, v := range CrawledURLs {
//		fmt.Println("KEY "+k, v)
//	}
//}
//
//func crawlUrl(crawlMap map[string]interface{}, depth int) {
//	var mainCH = make(chan bool)
//
//	var RecursiveCrawl func(url string, depth int, ch chan bool)
//	RecursiveCrawl = func(url string, depth int, ch chan bool) {
//		defer func() { ch <- true }()
//		if depth <= 0 {
//			return
//		}
//		mu.Lock()
//		collectToVisitedMap(url, depth-1)
//		mu.Unlock()
//
//		crawlMap = getCrawlMap(url)
//
//		innerCH := make(chan bool)
//
//		for _, v := range crawlMap {
//			if v == nil {
//				continue
//			}
//			if jsonMap, ok := v.(map[string]interface{}); ok {
//				url = collectUrl(jsonMap)
//				go RecursiveCrawl(url, depth-1, innerCH)
//			}
//		}
//		// receive all child URL goroutines here
//		for range url {
//			<-innerCH
//		}
//		return
//
//	}
//
//	go RecursiveCrawl(url, depth, mainCH)
//	<-mainCH
//	return
//}
//
//func crawlUrlPlain(crawlMap map[string]interface{}, depth int) {
//
//
//	if depth <= 0 {
//		return
//	}
//
//	for _, v := range crawlMap {
//		if v == nil {
//			continue
//		}
//		if jsonMap, ok := v.(map[string]interface{}); ok {
//			crawlUrlPlain(jsonMap, depth-1)
//		}
//
//		if arrayMap, ok := v.([]interface{}); ok {
//			for _, value := range arrayMap {
//				if arrayMap, ok := value.(map[string]interface{}); ok {
//					crawlUrlPlain(arrayMap, depth-1)
//				}
//			}
//		}
//
//		strVal := collectUrl(v)
//		//mu.Lock()
//		collectToVisitedMap(strVal, depth-1)
//		//mu.Unlock()
//	}
//}
//func collectToVisitedMap(strVal string, depth int) {
//	if strVal == "" {
//		return
//	}
//	if !CrawledURLs[strVal] {
//		CrawledURLs[strVal] = true
//		if !strings.HasSuffix(strVal, "png") && !strings.HasSuffix(strVal, "jpeg") {
//			newCrawlMap := getCrawlMap(strVal)
//			if newCrawlMap != nil {
//				crawlUrlPlain(newCrawlMap, depth)
//			}
//		}
//	} else {
//		return
//	}
//}
//func collectUrl(v interface{}) string {
//	var strVal string
//	var ok bool
//
//	if strVal, ok = v.(string); ok {
//		_, err := url.ParseRequestURI(strVal)
//		if err != nil || !strings.HasPrefix(strVal, "https") {
//			return ""
//		}
//		return strVal
//	}
//	return ""
//}
//func getCrawlMap(url string) map[string]interface{} {
//	resp, err := http.Get(url)
//	if err != nil {
//	}
//	var target map[string]interface{}
//	err = json.NewDecoder(resp.Body).Decode(&target)
//	if err != nil {
//	}
//	return target
//}

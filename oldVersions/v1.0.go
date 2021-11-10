package oldVersions
//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//	"sync"
//)
//
//var CrawledURLs = make(map[string]bool)
//var ALLURLS []string
//var mu sync.Mutex
//var StartingUrl = "https://l1-1.anzu.io/ios/1/baadc6f86730c7fcfbe5f5c3"
//var maxRecursionDepth = 5
//
////var StartingUrl = "https://dynamic-l1.anzuinfra.com/config/ios/1/baadc6f86730c7fcfbe5f5c3?"
//
//func main() {
//	CrawledURLs[StartingUrl] = true
//
//	crawlUrl(StartingUrl, maxRecursionDepth)
//	fmt.Println(len(CrawledURLs))
//
//	//for k, v := range CrawledURLs {
//	//	fmt.Println("KEY "+k, v)
//	//}
//}
//
//func crawlUrl(url string, depth int) {
//	var mainCH = make(chan bool, 2)
//
//	var RecursiveCrawl func(url string, depth int, ch chan bool)
//	RecursiveCrawl = func(url string, depth int, ch chan bool) {
//		defer func() { ch <- true }()
//		if depth <= 0 {
//			return
//		}
//		urls := collectUrls(url, depth)
//		log.Println(urls)
//		if len(urls) <= 0 {
//			return
//		}
//
//		innerCH := make(chan bool)
//		for _, u := range urls {
//			mu.Lock()
//			collectToVisitedMap(u)
//			mu.Unlock()
//			go RecursiveCrawl(u, depth-1, innerCH)
//		}
//
//		for range urls {
//			<-innerCH
//		}
//		return
//	}
//
//	go RecursiveCrawl(url, depth-1, mainCH)
//
//	<-mainCH
//	return
//}
//
//func collectUrls(url string, depth int) []string {
//	var urls []string
//	crawlMap := getCrawlMap(url)
//	if crawlMap == nil {
//		return nil
//	}
//	for _, v := range crawlMap {
//		if v == nil {
//			continue
//		}
//
//		parseInternalLinks(v, depth)
//
//		if len(ALLURLS) > 0 {
//			urls = append(urls, ALLURLS...)
//		}
//	}
//
//	return ALLURLS
//}
//
//func parseInternalLinks(v interface{}, depth int) {
//	if depth <= 0 {
//		return
//	}
//
//	if jsonMap, ok := v.(map[string]interface{}); ok {
//		for _, v := range jsonMap {
//			extractUrl(v)
//			parseInternalLinks(v, depth-1)
//		}
//	}
//
//	if arrayMap, ok := v.([]interface{}); ok {
//		for _, value := range arrayMap {
//			if arrayMap, ok := value.(map[string]interface{}); ok {
//				extractUrl(v)
//				parseInternalLinks(arrayMap, depth-1)
//			}
//		}
//	}
//}
//
//func collectToVisitedMap(strVal string) {
//	if strVal != "" && !CrawledURLs[strVal] {
//		CrawledURLs[strVal] = true
//		if !strings.HasSuffix(strVal, "png") && !strings.HasSuffix(strVal, "jpeg") {
//		}
//	}
//}
//
//func extractUrl(v interface{}) {
//	var strVal string
//	var ok bool
//
//	if strVal, ok = v.(string); ok {
//		_, err := url.ParseRequestURI(strVal)
//		if err == nil && strings.HasPrefix(strVal, "https") && strVal != "" {
//
//			ALLURLS = append(ALLURLS, strVal)
//		}
//	}
//}
//
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

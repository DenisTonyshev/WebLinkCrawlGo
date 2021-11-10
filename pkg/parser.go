package pkg

type Parser interface {
	ParseJsons(ch chan map[string]interface{}, tempCrawlMap map[string]bool)
}

var CrawlDepth = 3

func ParseJsons(ch chan map[string]interface{}, tempCrawlMap map[string]bool) {
	for chanVal := range ch {
		urls := parseJson(chanVal, CrawlDepth, tempCrawlMap)
		if urls != nil {
			Run(urls)
		}
	}
}

func parseJson(res map[string]interface{}, depth int, tempCrawlMap map[string]bool) (parsedUrls []string) {
	for _, v := range res {
		searchInternalLinks(v, depth, tempCrawlMap)
	}
	return CollectToArray(tempCrawlMap)
}

func searchInternalLinks(v interface{}, depth int, tempCrawlMap map[string]bool) {
	if depth <= 0 {
		return
	}

	if jsonMap, ok := v.(map[string]interface{}); ok {
		for _, v := range jsonMap {
			searchInternalLinks(v, depth-1, tempCrawlMap)
		}
	}

	if arrayMap, ok := v.([]interface{}); ok {
		for _, value := range arrayMap {
			if arrayMap, ok := value.(map[string]interface{}); ok {
				searchInternalLinks(arrayMap, depth-1, tempCrawlMap)
			}
		}
	}
	CheckUrl(v, tempCrawlMap)
}

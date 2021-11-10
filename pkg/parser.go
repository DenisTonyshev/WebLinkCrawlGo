package pkg

var CrawlDepth = 3

func ParseJsons(ch chan map[string]interface{}, saver FileSaver) {
	for value := range ch {
		urls := parseJson(value, CrawlDepth, saver)
		if urls != nil {
			Run(urls, saver)
		}
	}
}

func parseJson(res map[string]interface{}, depth int, saver FileSaver) (parsedUrls []string) {
	for _, v := range res {
		searchInternalLinks(v, depth, saver)
	}
	return CollectToArray(saver.tempMap)
}

func searchInternalLinks(v interface{}, depth int, saver FileSaver) {
	if depth <= 0 {
		return
	}

	if jsonMap, ok := v.(map[string]interface{}); ok {
		for _, v := range jsonMap {
			searchInternalLinks(v, depth-1, saver)
		}
	}

	if arrayMap, ok := v.([]interface{}); ok {
		for _, value := range arrayMap {
			if arrayMap, ok := value.(map[string]interface{}); ok {
				searchInternalLinks(arrayMap, depth-1, saver)
			}
		}
	}
	CheckUrl(v, saver)
}

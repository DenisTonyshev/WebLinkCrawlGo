package pkg

type Crawler interface {
	Run()
}

var ConcurrentRequestsNmb = 2

func Run(urls []string) {
	tempCrawlMap := make(map[string]bool)
	mainCH := make(chan map[string]interface{}, ConcurrentRequestsNmb)
	go GetJson(urls, mainCH)
	ParseJsons(mainCH, tempCrawlMap)
}

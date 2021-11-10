package pkg

import "sync"

var ConcurrentRequestsNmb = 1000

func Run(urls []string, saver FileSaver) {
	saver.tempMap = make(map[string]bool)
	mainCH := make(chan map[string]interface{}, ConcurrentRequestsNmb)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go GetJson(url, mainCH, &wg)
	}
	go func() {
		wg.Wait()
		close(mainCH)
	}()
	ParseJsons(mainCH, saver)
}

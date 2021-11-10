package pkg

type Save interface {
	SaveUrls(tempCrawlMap map[string]bool) []string
}

func SaveToMap(value string, tempMap map[string]bool, mainMap map[string]bool) {
	if !mainMap[value] {
		mainMap[value] = true
		tempMap[value] = true
	} else {
		return
	}
}

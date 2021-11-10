package pkg

import (
	"net/url"
	"strings"
	"sync"
)

var mu sync.Mutex
var MainMap = make(map[string]bool)

type UrlSchema interface {
	CheckUrl(v interface{}, tempCrawlMap map[string]bool)
}
type Saver struct {
	save Save
}

func CheckUrl(v interface{}, tempCrawlMap map[string]bool) {
	var strVal string
	var ok bool

	if strVal, ok = v.(string); ok {
		_, err := url.ParseRequestURI(strVal)
		if err != nil || !strings.HasPrefix(strVal, "https") {
			return
		}

		mu.Lock()
		SaveToMap(strVal, tempCrawlMap, MainMap)
		mu.Unlock()
	}
}

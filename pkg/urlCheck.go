package pkg

import (
	"net/url"
	"strings"
	"sync"
)

var mu sync.Mutex

func CheckUrl(v interface{}, tempCrawlMap map[string]bool) {
	var strVal string
	var ok bool

	if strVal, ok = v.(string); ok {
		_, err := url.ParseRequestURI(strVal)
		if err != nil || !strings.HasPrefix(strVal, "https") {
			return
		}

		mu.Lock()
		err = fileSaver.Save(strVal)
		if err != nil {
			return
		}
		mu.Unlock()
	}
}

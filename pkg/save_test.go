package pkg

import "testing"

func TestSave(t *testing.T) {
	testMainMap := make(map[string]bool)
	testTempMap := make(map[string]bool)
	testValue := "test1.com"
	SaveToMap(testValue, testTempMap, testMainMap)
	if len(testMainMap) == 0 {
		t.Errorf("Empty")
	}
	if testMainMap[testValue] != true {
		t.Errorf("wrong values saved")
	}

}

//func SaveToMap(value string, tempMap map[string]bool) {
//	if !CrawledURLs[value] {
//		CrawledURLs[value] = true
//		tempMap[value] = true
//	} else {
//		return
//	}
//}

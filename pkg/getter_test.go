package pkg

import (
	"testing"
)

func TestGetJson(t *testing.T) {
	testCH := make(chan map[string]interface{}, 1)
	testArr := []string{"https://l1-1.anzu.io/ios/"}

	go GetJson(testArr, testCH)

	for val := range testCH {
		if val == nil{
			t.Errorf("Empty or Wrong format")
		}
	}
}

package pkg

import "testing"

func TestCollect(t *testing.T) {
	testMap := map[string]bool{
		"https://test1.com": true,
		"https://test2.com": true,
		"https://test3.com": true,
	}
	testArr := CollectToArray(testMap)

	if len(testArr) != 3 {
		t.Errorf("Collected Wrong")
	}

}

package pkg

import "fmt"

func PrintToConsole(mainMap map[string]bool) {
	for k, _ := range mainMap {
		fmt.Println(k)
	}
	fmt.Println(len(mainMap))
}

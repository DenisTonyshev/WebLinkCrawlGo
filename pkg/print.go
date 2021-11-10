package pkg

import "fmt"

type Print interface {
	PrintToConsole(crawledMap map[string]bool)
}

func PrintToConsole(crawledMap map[string]bool) {
	for k, _ := range crawledMap {
		fmt.Println(k)
	}
	fmt.Println(len(crawledMap))
}

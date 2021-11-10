package main

import (
	"awesomeProject/pkg"
)

func main() {
	saver := pkg.FileSaver{}
	saver.ResultMap = make(map[string]bool)
	StartingUrl := []string{"https://l1-1.anzu.io/ios/1/baadc6f86730c7fcfbe5f5c3"}
	pkg.Run(StartingUrl, saver)
	pkg.PrintToConsole(saver.ResultMap)
}

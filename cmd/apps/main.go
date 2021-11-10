package main

import (
	"awesomeProject/pkg"
)

func main() {
	StartingUrl := []string{"https://l1-1.anzu.io/ios/1/baadc6f86730c7fcfbe5f5c3"}
	pkg.Run(StartingUrl)
	pkg.PrintToConsole(pkg.MainMap)
}

package main

import (
	"fmt"
	"./crawler"
)

func main() {
	
	rawHTML := crawler.FetchHTML("http://google.com")
	fmt.Printf(string(rawHTML) + "\n")

}

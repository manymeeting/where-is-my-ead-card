package main

import (
	"fmt"
	"./crawler"
	"./parser"
)

func main() {
	
	rawHTML := crawler.FetchHTML("https://www.checkuscis.com/")
	fmt.Printf(string(rawHTML) + "\n")

	htmlparser.ParseCheckUSCIS(rawHTML)

}

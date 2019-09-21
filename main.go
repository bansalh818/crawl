// file: crawl.go
package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/bansalh818/crawler/parser"
	"github.com/bansalh818/crawler/utils"
)

func main() {
	var siteURL string
	var depth int
	var err error
	var visitedURL utils.Set

	flag.Parse()
	args := flag.Args()

	// Fetch the URL and depth
	// default URL : monzo.com and depth : 1
	switch len(args) {
	case 1:
		siteURL = args[0]
		depth = 1
	case 2:
		siteURL = args[0]
		depth, err = strconv.Atoi(args[1])
		if err != nil {
			//log
			return
		}
	default:
		siteURL = "https://google.com"
		depth = 1
	}

	// Parse the site URL
	URL, err := utils.Parse(siteURL)
	if err != nil {
		// Log failed to parse site url
		return
	}

	//Keeping track of time taken by the progam.
	start := time.Now()

	site := utils.Site{URL: URL}

	visitedURL.Add(siteURL)

	utils.Waitlist.Add(1)
	go parser.GetLinks(&site, &visitedURL, depth)

	utils.Waitlist.Wait()

	//utils.PrintSite(&site, 0)

	elapsed := time.Since(start)
	//Log the output
	fmt.Println(elapsed)
	fmt.Println(visitedURL.Size())
	fmt.Println(depth)
}

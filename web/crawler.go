package web

import (
	"github.com/bansalh818/crawler/parser"
	"github.com/bansalh818/crawler/utils"
)

//
type Crawler interface {
	Crawl(baseURL string, depth int) (*Site, error)
}

//
type CrawlerImpl struct {
}

// Crawl s
func (c *CrawlerImpl) Crawl(baseURL string, depth int) (*parser.Site, error) {
	var visitedURL utils.Set

	// Parse the site URL
	URL, err := utils.Parse(baseURL)
	if err != nil {
		// Log failed to parse site url
		return nil, err
	}

	site := utils.Site{URL: URL}

	visitedURL.Add(baseURL)

	//utils.Waitlist.Add(1)
	go parser.ExtractLinks(&site, &visitedURL, depth)

	//utils.Waitlist.Wait()
}

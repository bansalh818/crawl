package parser

import (
	"net/http"
	"sync"

	"github.com/bansalh818/crawler/utils"

	"golang.org/x/net/html"
	//"github.com/sirupsen/logrus"
)

type Tree struct {
	Head  *Site
	Depth int
}

func (*Tree) ExtractLinks(visitedURL *utils.Set) {

}

// GetLinks s
func GetLinks(headSite *utils.Site, visitedURL *utils.Set, depth int) {
	defer utils.Waitlist.Done()

	var localURLs utils.Set

	// Max depth reached
	if depth <= 0 {
		//return
	}

	// Get page html
	resp, err := http.Get((*headSite).URL.String())
	if err != nil {
		//Log error
		return
	}

	// Create channel to add links to Site
	links := make(chan *utils.Site)
	var linkswg sync.WaitGroup

	linkswg.Add(1)
	defer linkswg.Done()

	utils.Waitlist.Add(1)
	// Close link channels when parsing finishes
	go func() {
		defer utils.Waitlist.Done()
		linkswg.Wait()
		close(links)
	}()

	utils.Waitlist.Add(1)
	// Add link to site
	go func() {
		defer utils.Waitlist.Done()
		for link := range links {
			(*headSite).AddLink(link)
		}
	}()

	tokenizer := html.NewTokenizer(resp.Body)
	defer resp.Body.Close()
	for {
		// get the next token type
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			// get the token
			token := tokenizer.Token()
			switch token.DataAtom.String() {
			case "a", "link": //link tags
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						if !localURLs.Has(attr.Val) {
							localURLs.Add(attr.Val)
							linkswg.Add(1)
							go ParseLink(attr.Val, headSite, links, visitedURL, &linkswg, depth)
						}
					}
				}
			}
		}
	}
}

// ParseLink a
func ParseLink(link string, headSite *utils.Site, result chan *utils.Site, visitedURL *utils.Set, linkswg *sync.WaitGroup, depth int) {
	defer (*linkswg).Done()

	tempURL, err := utils.Parse(link)
	if err != nil {
		// Log error
		return
	}

	newURL, err := utils.ResolveURL(tempURL, headSite.URL)
	if err != nil {
		// Log
		return
	}

	// Check for external links
	if !utils.IsSameHost(newURL, headSite.URL) {
		return
	}

	// Check whether we have seen the URL before
	if visitedURL.Has(newURL.String()) {
		// Add link to the page but do not crawl
		newSite := utils.Site{URL: newURL, Links: nil}
		result <- &newSite
		return
	}
	visitedURL.Add(newURL.String())

	// Add new url as site
	newSite := utils.Site{URL: newURL}
	utils.Waitlist.Add(1)
	go GetLinks(&newSite, visitedURL, depth-1)

	result <- &newSite
	return
}

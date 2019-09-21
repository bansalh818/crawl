package parser

import "net/url"

// Site
type Site struct {
	URL   *url.URL
	Links []*Site
}

// AddLink
func (s *Site) AddLink(link *Site) {
	s.Links = append(s.Links, link)
}

func (s *Site) PrintSite() {

}

package utils

import "net/url"

// // Gethost returns the hostname of the url
// func Gethost(rawurl string) (string, error) {
// 	targetURL, err := url.Parse(rawurl)
// 	if err != nil {
// 		return "", err
// 	} else {
// 		return targetURL.Host, nil
// 	}
// }

// Parse add
func Parse(rawurl string) (*url.URL, error) {
	targetURL, err := url.Parse(rawurl)
	return targetURL, err
}

// IsSameHost add
func IsSameHost(rawURL, siteURL *url.URL) bool {
	if rawURL.Host == siteURL.Host {
		return true
	}
	return false
}

// ResolveURL resolves the rawurl
func ResolveURL(rawURL, siteURL *url.URL) (*url.URL, error) {
	newURL := siteURL.ResolveReference(rawURL)
	newURL.Fragment = ""
	return newURL, nil
}

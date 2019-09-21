package utils

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetBody t
func GetBody(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	if checks(resp) {
		return resp, nil
	}
	return nil, err
}

func checks(resp *http.Response) bool {
	if resp.StatusCode != http.StatusOK {
		return false
	}

	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		return false
	}

	return true

}

func IsTokenError(tokenType html.TokenType) bool {
	if tokenType == html.ErrorToken {
		return false
	}
	return true
}

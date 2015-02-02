package crawl

import (
	"fmt"
	"net/url"
	"strings"
)

func normalize(rawURL string) (*url.URL, error) {
	if rawURL == "" {
		return nil, fmt.Errorf("crawl: empty URLs are invalid")
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	// If no host was determined, split on the first slash and set
	// that as host
	if parsed.Host == "" {
		parts := strings.SplitN(parsed.Path, "/", 2)
		// There will always be at least one part
		parsed.Host = parts[0]
		if len(parts) > 1 && parts[1] != "" {
			parsed.Path = "/" + parts[1]
		} else {
			parsed.Path = ""
		}
	}
	// Assume http if not given
	if parsed.Scheme == "" {
		parsed.Scheme = "http"
	}
	return parsed, nil
}

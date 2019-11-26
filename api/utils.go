package api

import (
	"net/url"
	"strings"
)

// TrimMultiline - trims multiline
func TrimMultiline(multi string) string {
	res := ""
	for _, line := range strings.Split(multi, "\n") {
		res += strings.Trim(line, "\t")
	}
	return res
}

// GetConfHeaders resolves headers from config overrides
func GetConfHeaders(conf *Conf) map[string]string {
	headers := map[string]string{}
	if conf != nil {
		headers = conf.Headers
	}
	return headers
}

// GetRelativeURL out of an absolute one
func GetRelativeURL(absURL string) string {
	url, _ := url.Parse(absURL)
	return url.Path
}

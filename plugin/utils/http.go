package utils

import (
	"fmt"

	"github.com/navidrome/navidrome/plugins/pdk/go/host"
)

func DoGetRequest(endpoint string) ([]byte, error) {
	acceptLanguage := ConfigSearchLanguage()
	userAgent := ConfigUserAgent()
	httpAcceptHeader := ConfigSearchHTTPAcceptHeader()
	shazamCookie := ConfigShazamCookie()

	resp, err := host.HTTPSend(host.HTTPRequest{
		Method: "GET",
		URL:    endpoint,
		Headers: map[string]string{
			"Accept":          httpAcceptHeader,
			"Accept-Language": acceptLanguage,
			"User-Agent":      userAgent,
			"Cookie":          shazamCookie,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Shazam endpoint %s: %w", endpoint, err)
	}

	if resp.StatusCode != HTTPStatusOK {
		return resp.Body, fmt.Errorf("error code %d returned from Shazam for endpoint %s", resp.StatusCode, endpoint)
	}
	return resp.Body, nil
}

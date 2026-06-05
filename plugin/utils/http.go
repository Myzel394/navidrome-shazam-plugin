package utils

import (
	"github.com/extism/go-pdk"
)

const (
	DefaultLang    = "en-US"
	DefaultCountry = "GB"
)

// DoGetRequest sends an HTTP GET and returns the raw body bytes, or nil on error.
func DoGetRequest(endpoint string) []byte {
	req := pdk.NewHTTPRequest(pdk.MethodGet, endpoint)
	req.SetHeader("Accept", "*/*")
	req.SetHeader("Accept-Language", DefaultLang)
	req.SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3 Mobile/15E148 Safari/604.1")

	resp := req.Send()
	if resp.Status() != 200 {
		return nil
	}
	return resp.Body()
}

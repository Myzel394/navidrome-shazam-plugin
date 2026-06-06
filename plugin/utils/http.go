package utils

import (
	"github.com/extism/go-pdk"
)

func DoGetRequest(endpoint string) []byte {
	acceptLanguage := ConfigSearchLanguage()
	userAgent := ConfigUserAgent()
	httpAcceptHeader := ConfigSearchHTTPAcceptHeader()
	shazamCookie := ConfigShazamCookie()

	req := pdk.NewHTTPRequest(pdk.MethodGet, endpoint)
	req.SetHeader("Accept", httpAcceptHeader)
	req.SetHeader("Accept-Language", acceptLanguage)
	req.SetHeader("User-Agent", userAgent)
	req.SetHeader("Cookie", shazamCookie)

	resp := req.Send()
	if resp.Status() != 200 {
		return nil
	}
	return resp.Body()
}

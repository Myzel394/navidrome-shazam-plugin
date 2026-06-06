package utils

import (
	"fmt"

	"github.com/navidrome/navidrome/plugins/pdk/go/pdk"
)

func DoGetRequest(endpoint string) ([]byte, error) {
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
		return resp.Body(), fmt.Errorf("navidrome-shazam-plugin: error code %d returned from Shazam for endpoint %s", resp.Status(), endpoint)
	}
	return resp.Body(), nil
}

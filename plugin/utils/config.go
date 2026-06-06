package utils

import (
	"strconv"

	"github.com/extism/go-pdk"
)

// ConfigSearchCountry reads the configured Apple Music storefront country
// from Navidrome's plugin config, defaulting to "GB".
func ConfigSearchCountry() string {
	v, ok := pdk.GetConfig("shazam_search_country")
	if !ok || v == "" {
		return "US"
	}
	return v
}

func ConfigSearchLanguage() string {
	v, ok := pdk.GetConfig("shazam_language")

	if !ok || v == "" {
		return "en"
	}

	return v
}

func ConfigShazamCookie() string {
	v, ok := pdk.GetConfig("shazam_cookie")

	if !ok || v == "" {
		return ""
	}

	return v
}

func ConfigSearchHTTPAcceptHeader() string {
	v, ok := pdk.GetConfig("shazam_http_accept")

	if !ok || v == "" {
		return "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	}

	return v
}

func ConfigSearchLimit() int {
	v, ok := pdk.GetConfig("shazam_search_limit")

	if !ok || v == "" {
		return 10
	}

	limit, err := strconv.Atoi(v)
	if err != nil {
		pdk.Log(pdk.LogError, "shazam lyrics: failed to parse shazam_search_limit config value, defaulting to 10")
		return 10
	}

	return limit
}

func ConfigSearchLevenshteinThreshold() float64 {
	v, ok := pdk.GetConfig("shazam_search_levenshtein_threshold")

	if !ok || v == "" {
		return 0.85
	}

	threshold, err := strconv.ParseFloat(v, 64)
	if err != nil {
		pdk.Log(pdk.LogError, "shazam lyrics: failed to parse shazam_search_levenshtein_threshold config value, defaulting to 0.85")
		return 0.85
	}

	return threshold
}

func ConfigUserAgent() string {
	v, ok := pdk.GetConfig("shazam_user_agent")

	if !ok || v == "" {
		return "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3 Mobile/15E148 Safari/604.1"
	}

	return v
}

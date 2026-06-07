package utils

import (
	"strconv"

	"github.com/navidrome/navidrome/plugins/pdk/go/pdk"
)

func ConfigSearchCountry() string {
	v, ok := pdk.GetConfig(ConfigKeySearchCountry)
	if !ok || v == "" {
		return DefaultSearchCountry
	}
	return v
}

func ConfigSearchLanguage() string {
	v, ok := pdk.GetConfig(ConfigKeySearchLanguage)

	if !ok || v == "" {
		return DefaultSearchLanguage
	}

	return v
}

func ConfigShazamCookie() string {
	v, ok := pdk.GetConfig(ConfigKeyCookie)

	if !ok || v == "" {
		return ""
	}

	return v
}

func ConfigSearchHTTPAcceptHeader() string {
	v, ok := pdk.GetConfig(ConfigKeyHTTPAccept)

	if !ok || v == "" {
		return DefaultHTTPAccept
	}

	return v
}

func ConfigSearchLimit() int {
	v, ok := pdk.GetConfig(ConfigKeySearchLimit)

	if !ok || v == "" {
		return DefaultSearchLimit
	}

	limit, err := strconv.Atoi(v)
	if err != nil {
		LogErrorf("failed to parse %s config value, defaulting to %d", ConfigKeySearchLimit, DefaultSearchLimit)
		return DefaultSearchLimit
	}

	return limit
}

func ConfigSearchLevenshteinThreshold() float64 {
	v, ok := pdk.GetConfig(ConfigKeyLevenshteinThreshold)

	if !ok || v == "" {
		return DefaultLevenshteinThreshold
	}

	threshold, err := strconv.ParseFloat(v, 64)
	if err != nil {
		LogErrorf("failed to parse %s config value, defaulting to %v", ConfigKeyLevenshteinThreshold, DefaultLevenshteinThreshold)
		return DefaultLevenshteinThreshold
	}

	return threshold
}

func ConfigUserAgent() string {
	v, ok := pdk.GetConfig(ConfigKeyUserAgent)

	if !ok || v == "" {
		return DefaultUserAgent
	}

	return v
}

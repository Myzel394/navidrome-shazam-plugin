package shazam

import (
	"github.com/extism/go-pdk"
)

const configKeySearchCountry = "shazam_search_country"

// configSearchCountry reads the configured Apple Music storefront country
// from Navidrome's plugin config, defaulting to "GB".
func configSearchCountry() string {
	v, ok := pdk.GetConfig(configKeySearchCountry)
	if !ok || v == "" {
		return "US"
	}
	return v
}

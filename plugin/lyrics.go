package main

import (
	"github.com/Myzel394/navidrome-lyrics-scrape-plugin/plugin/providers/shazam"
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

func (p *plugin) GetLyrics(input lyrics.GetLyricsRequest) (lyrics.GetLyricsResponse, error) {
	return shazam.FetchLyrics(input)
}

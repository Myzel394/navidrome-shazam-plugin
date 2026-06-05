package main

import (
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

func (p *plugin) GetLyrics(input lyrics.GetLyricsRequest) (lyrics.GetLyricsResponse, error) {
	return lyrics.GetLyricsResponse{
		Lyrics: []lyrics.LyricsText{
			{
				Lang: "en",
				Text: "These are the lyrics for " + input.Track.Title + " by " + input.Track.Artist,
			},
		},
	}, nil
}


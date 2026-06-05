package shazam

import (
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

type Song struct {
	ID     string
	Artist string
	Title  string
}

func FetchLyrics(input lyrics.GetLyricsRequest) (lyrics.GetLyricsResponse, error) {
	track := searchForTrack(input)

	if track == nil {
		return lyrics.GetLyricsResponse{}, nil
	}

	result, err := fetchLyricsForTrack(track)

	return result, err
}

package shazam

import (
	"fmt"

	"github.com/Myzel394/navidrome-shazam-plugin/plugin/utils"
	"github.com/extism/go-pdk"
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

func fetchLyricsForTrack(track *Song) (lyrics.GetLyricsResponse, error) {
	// Normalize already stripped brackets, so "Never Gonna Give You Up (2022 Remaster)"
	// becomes "never gonna give you up" → slug: "never-gonna-give-you-up"
	slug := slugify(track.Title)

	endpoint := fmt.Sprintf(
		"https://www.shazam.com/song/%s/%s",
		track.ID, slug,
	)

	body := utils.DoGetRequest(endpoint)
	if body == nil {
		return lyrics.GetLyricsResponse{}, fmt.Errorf("shazam lyrics: failed to do shazam fetchLyrics request for track ID %s", track.ID)
	}

	text, err := extractLyricsFromHTML(string(body))
	if err != nil {
		return lyrics.GetLyricsResponse{}, err
	}

	pdk.Log(pdk.LogInfo, fmt.Sprintf("shazam lyrics: found lyrics for track ID %s", track.ID))

	return lyrics.GetLyricsResponse{
		Lyrics: []lyrics.LyricsText{
			{
				Text: text,
			},
		},
	}, nil
}

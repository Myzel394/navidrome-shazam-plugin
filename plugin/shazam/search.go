package shazam

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Myzel394/navidrome-shazam-plugin/plugin/utils"
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

type searchHit struct {
	ID         string `json:"id"`
	Attributes struct {
		ArtistName string `json:"artistName"`
		Name       string `json:"name"`
	} `json:"attributes"`
}

type searchResponse struct {
	Results struct {
		Songs struct {
			Data []searchHit `json:"data"`
		} `json:"songs"`
	} `json:"results"`
}

func searchForTrack(input lyrics.GetLyricsRequest) (*Song, error) {
	normArtist := normalize(input.Track.Artist)
	normTitle := normalize(input.Track.Title)

	// Primary: artist + track query.
	query := normArtist + " " + normTitle
	country := utils.ConfigSearchCountry()
	searchLimit := utils.ConfigSearchLimit()
	endpoint := fmt.Sprintf(
		"https://www.shazam.com/services/amapi/v1/catalog/%s/search?term=%s&types=songs&limit=%d",
		country, url.QueryEscape(query), searchLimit,
	)

	body, err := utils.DoGetRequest(endpoint)
	if err != nil || body == nil {
		return nil, fmt.Errorf("navidrome-shazam-plugin: failed to do shazam search request for query %s; Error: %v", query, err)
	}

	var result searchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("navidrome-shazam-plugin: failed to parse shazam search response for query %s", query)
	}

	if best := pickBestMatch(result.Results.Songs.Data, normArtist, normTitle); best != nil {
		return best, nil
	}

	// Fallback: search track-only (sometimes short queries rank better).
	endpointFallback := fmt.Sprintf(
		"https://www.shazam.com/services/amapi/v1/catalog/%s/search?term=%s&types=songs&limit=%d",
		country, url.QueryEscape(normTitle), searchLimit,
	)
	bodyFallback, err := utils.DoGetRequest(endpointFallback)
	if err != nil || bodyFallback == nil {
		return nil, fmt.Errorf("navidrome-shazam-plugin: failed to do shazam search request for fallback query %s; Error: %v", normTitle, err)
	}

	var resultFallback searchResponse
	if err := json.Unmarshal(bodyFallback, &resultFallback); err != nil {
		return nil, fmt.Errorf("navidrome-shazam-plugin: failed to parse shazam search response for fallback query %s", normTitle)
	}

	return pickBestMatch(resultFallback.Results.Songs.Data, normArtist, normTitle), nil
}

// pickBestMatch scores every hit against the normalized artist/title and
// returns the Song with the highest combined Levenshtein ratio, provided
// both individual ratios meet the minimum threshold.
func pickBestMatch(hits []searchHit, normArtist, normTitle string) *Song {
	matchThreshold := utils.ConfigSearchLevenshteinThreshold()

	var bestSong *Song
	var bestScore float64

	for _, song := range hits {
		hitArtist := normalize(song.Attributes.ArtistName)
		hitTitle := normalize(song.Attributes.Name)

		artistRatio := levenshteinRatio(normArtist, hitArtist)
		titleRatio := levenshteinRatio(normTitle, hitTitle)

		if artistRatio < matchThreshold || titleRatio < matchThreshold {
			continue
		}

		score := (artistRatio + titleRatio) / 2
		if score > bestScore {
			bestScore = score
			bestSong = &Song{
				ID:     song.ID,
				Artist: hitArtist,
				Title:  hitTitle,
			}
		}
	}

	return bestSong
}

package lyrics

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	lrclibUrl = "https://lrclib.net/api/search?"
)

var (
	lyricsCache   = make(map[string][]string)
	lyricsCacheMu sync.RWMutex
)

type GetLyricsRequest struct {
	TrackName  string
	ArtistName string
	AlbumName  string
}

func GetLyrics(r GetLyricsRequest) ([]string, error) {
	start := time.Now()

	// Create cache key
	key := r.TrackName + "|" + r.ArtistName + "|" + r.AlbumName

	// Check cache
	lyricsCacheMu.RLock()
	if lyrics, ok := lyricsCache[key]; ok {
		lyricsCacheMu.RUnlock()
		log.Printf("GetLyrics took %s (cached)", time.Since(start))
		return lyrics, nil
	}
	lyricsCacheMu.RUnlock()

	// Fetch from API
	url := lrclibUrl + url.Values{
		"track_name":  {r.TrackName},
		"artist_name": {r.ArtistName},
		"album_name":  {r.AlbumName},
	}.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response []response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	var lyrics []string
	if len(response) > 0 && response[0].PlainLyrics != "" {
		lyrics = strings.Split(response[0].PlainLyrics, "\n")
	} else {
		lyrics = []string{}
	}

	// Store in cache
	lyricsCacheMu.Lock()
	lyricsCache[key] = lyrics
	lyricsCacheMu.Unlock()

	log.Printf("GetLyrics took %s (fetched)", time.Since(start))
	return lyrics, nil
}

type response struct {
	PlainLyrics string `json:"plainLyrics"`
}

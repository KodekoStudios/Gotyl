package spotify

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	client     *spotify.Client
	clientOnce sync.Once

	// Caches
	searchCache   = make(map[string]*spotify.FullTrack)
	searchCacheMu sync.RWMutex

	labelCache   = make(map[spotify.ID]string)
	labelCacheMu sync.RWMutex
)

// GetClient returns a singleton Spotify client.
func GetClient() *spotify.Client {
	clientOnce.Do(func() {
		ctx := context.Background()

		config := &clientcredentials.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			TokenURL:     spotifyauth.TokenURL,
		}

		token, err := config.Token(ctx)
		if err != nil {
			panic("couldn't get Spotify token: " + err.Error())
		}

		httpClient := spotifyauth.New().Client(ctx, token)
		client = spotify.New(httpClient)
	})

	return client
}

// SearchTrack searches for a track and returns the first result (cached).
func SearchTrack(ctx context.Context, query string) (*spotify.FullTrack, error) {
	start := time.Now()

	// Check cache
	searchCacheMu.RLock()
	if track, ok := searchCache[query]; ok {
		searchCacheMu.RUnlock()
		log.Printf("SearchTrack took %s (cached)", time.Since(start))
		return track, nil
	}
	searchCacheMu.RUnlock()

	// Fetch from API
	results, err := GetClient().Search(ctx, query, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}

	if len(results.Tracks.Tracks) == 0 {
		log.Printf("SearchTrack took %s (no results)", time.Since(start))
		return nil, nil
	}

	track := &results.Tracks.Tracks[0]

	// Store in cache
	searchCacheMu.Lock()
	searchCache[query] = track
	searchCacheMu.Unlock()

	log.Printf("SearchTrack took %s (fetched)", time.Since(start))
	return track, nil
}

// GetAlbumLabel returns the copyright label for a track's album (cached).
func GetAlbumLabel(ctx context.Context, albumID spotify.ID) string {
	start := time.Now()

	// Check cache
	labelCacheMu.RLock()
	if label, ok := labelCache[albumID]; ok {
		labelCacheMu.RUnlock()
		log.Printf("GetAlbumLabel took %s (cached)", time.Since(start))
		return label
	}
	labelCacheMu.RUnlock()

	// Fetch from API
	album, err := GetClient().GetAlbum(ctx, albumID)
	if err != nil || album == nil {
		log.Printf("GetAlbumLabel took %s (error)", time.Since(start))
		return ""
	}

	label := ""
	if len(album.Copyrights) > 0 {
		label = album.Copyrights[0].Text
	}

	// Store in cache
	labelCacheMu.Lock()
	labelCache[albumID] = label
	labelCacheMu.Unlock()

	log.Printf("GetAlbumLabel took %s (fetched)", time.Since(start))
	return label
}

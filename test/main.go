package main

import (
	"context"
	"log"
	"os"
	"time"

	"gostyl/internal/card"
	"gostyl/internal/lyrics"
	"gostyl/internal/palette"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("missing Spotify credentials")
	}

	ctx := context.Background()

	// Spotify client
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	results, err := client.Search(
		ctx,
		"dance alone the vanished people",
		spotify.SearchTypeTrack,
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(results.Tracks.Tracks) == 0 {
		log.Fatal("no tracks found")
	}

	track := results.Tracks.Tracks[0]

	album, err := client.GetAlbum(
		ctx,
		track.ID,
	)
	// Use album label if available
	var label string
	if album != nil && album.Copyrights != nil && len(album.Copyrights) > 0 {
		label = album.Copyrights[0].Text
	} else {
		label = track.Artists[0].Name
	}

	lyrics, err := lyrics.GetLyrics(lyrics.GetLyricsRequest{
		TrackName:  track.Name,
		ArtistName: track.Artists[0].Name,
		AlbumName:  track.Album.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	cardData := &card.Card{
		URI:           string(track.URI),
		Theme:         palette.Light,
		Title:         track.Name,
		Artist:        track.Artists[0].Name,
		Duration:      card.FormatDuration(int(track.Duration)),
		ReleaseDate:   track.Album.ReleaseDateTime().Format("January 2, 2006"),
		Label:         label,
		Lyrics:        lyrics,
		CoverImage:    track.Album.Images[0].URL,
		ShowScannable: true,
	}

	// Cold render
	log.Println("──── Cold render (empty cache) ────")
	start := time.Now()

	pngCold, err := card.RenderPNG(cardData)
	if err != nil {
		log.Fatal(err)
	}

	coldTime := time.Since(start)
	log.Printf("Cold render took: %s", coldTime)

	// Warm render
	log.Println("──── Warm render (cached) ────")
	start = time.Now()

	pngWarm, err := card.RenderPNG(cardData)
	if err != nil {
		log.Fatal(err)
	}

	warmTime := time.Since(start)
	log.Printf("Warm render took: %s", warmTime)

	// Summary
	log.Println("──── Summary ────")
	log.Printf("Cold:  %s", coldTime)
	log.Printf("Warm:  %s", warmTime)
	log.Printf("Speedup: %.2fx", float64(coldTime)/float64(warmTime))

	// Write output (warm result)
	if err := os.WriteFile("out.png", pngWarm, 0644); err != nil {
		log.Fatal(err)
	}

	_ = pngCold

	log.Println("Rendered out.png successfully")
}

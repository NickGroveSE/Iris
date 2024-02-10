package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const (
	redirectURI     = "http://localhost:5000/callback"
	charsetForState = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lengthOfState   = 16
)

var (
	auth  *spotifyauth.Authenticator
	ch    = make(chan *spotify.Client)
	state string
)

type Track struct {
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	ImgURL     string `json:"imgurl"`
	SpotifyURL string `json:"spotifyurl"`
	// Glimpses   []string `json:"glimpses"`
}

type Connection struct {
	Tracks []*Track `json:"tracks"`
}

func main() {

	// Gin Init with CORS Middleware
	app := gin.Default()
	var client *spotify.Client

	app.Use(cors.Default())
	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.GET("/redirect", getRedirectURL)
	app.GET("/callback", completeAuth)
	app.GET("/music", func(c *gin.Context) {

		var tracks []Track
		var options []spotify.RequestOption
		options = append(options, spotify.Timerange("short_term"))
		options = append(options, spotify.Limit(10))

		topTracks, err := client.CurrentUsersTopTracks(context.Background(), options...)
		if err != nil {
			log.Println("Fatal at Tracks")
			log.Fatal(err)
		}

		for _, track := range topTracks.Tracks {

			tracks = append(tracks, Track{
				Name:       track.Name,
				Artist:     track.Artists[0].Name,
				ImgURL:     track.Album.Images[0].URL,
				SpotifyURL: track.ExternalURLs["spotify"],
			})

		}

		c.JSON(http.StatusOK, gin.H{"data": tracks})
	})

	go func() {
		client = <-ch

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			log.Println("Fatal at Client")
			log.Fatal(err)
		}

		fmt.Println("You are logged in as:", user.ID)
	}()

	// Link for
	// url := auth.AuthURL(state)
	// fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	app.Run(":5000")

}

func getRedirectURL(c *gin.Context) {

	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead),
		spotifyauth.WithClientID(""),
		spotifyauth.WithClientSecret(""))

	state = randomString()

	c.JSON(http.StatusOK, gin.H{"response": auth.AuthURL(state)})

}

func completeAuth(c *gin.Context) {
	tok, err := auth.Token(c, state, c.Request)
	if err != nil {
		log.Println("Fatal at Token")
		http.Error(c.Writer, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := c.Request.FormValue("state"); st != state {
		log.Println("Fatal at State")
		http.NotFound(c.Writer, c.Request)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	callbackClient := spotify.New(auth.Client(c, tok))
	user, err := callbackClient.CurrentUser(context.Background())
	if err != nil {
		log.Println("Fatal at Client")
		log.Fatal(err)
	}

	frontendURL := "http://localhost:5173/" + user.ID

	c.Redirect(http.StatusMovedPermanently, frontendURL)
	ch <- callbackClient
}

func randomString() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lengthOfState)
	for i := range b {
		b[i] = charsetForState[seededRand.Intn(len(charsetForState))]
	}
	return string(b)
}

// func handleArtistNames() string {

// }

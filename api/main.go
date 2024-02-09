package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:5000/callback"

var (
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead),
		spotifyauth.WithClientID(""),
		spotifyauth.WithClientSecret(""))
	ch    = make(chan *spotify.Client)
	state = "12345"
)

func main() {

	// Gin Init with CORS Middleware
	app := gin.Default()
	app.Use(cors.Default())

	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.GET("/redirect", getRedirectURL)
	app.GET("/callback", completeAuth)
	app.GET("/music", getMusicData)

	go func() {
		client := <-ch

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			log.Println("Fatal at Client")
			log.Fatal(err)
		}

		// topTracks, err := client.CurrentUsersTopTracks(context.Background(), spotify.Timerange("long_term"))

		fmt.Println("You are logged in as:", user.ID)
		// fmt.Println("Top Tracks: ", topTracks.Tracks)
	}()

	// Link for
	// url := auth.AuthURL(state)
	// fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	app.Run(":5000")

}

func getRedirectURL(c *gin.Context) {

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

	client := spotify.New(auth.Client(c, tok))
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Println("Fatal at Client")
		log.Fatal(err)
	}

	values := c.Request.URL.Query()

	frontendURL := "http://localhost:5173/" + user.ID + "?code=" + values.Get("code") + "&state=" + values.Get("state")

	c.Redirect(http.StatusMovedPermanently, frontendURL)
	ch <- client
}

func getMusicData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"response": "music data"})

}

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

const redirectURI = "http://localhost:5000/callback"
const charsetForState = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const lengthOfState = 16

var (
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead),
		spotifyauth.WithClientID(""),
		spotifyauth.WithClientSecret(""))
	ch                    = make(chan *spotify.Client)
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	state                 = randomString()
	client     *spotify.Client
)

func main() {

	// Gin Init with CORS Middleware
	app := gin.Default()

	app.Use(cors.Default())

	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.GET("/redirect", getRedirectURL)
	app.GET("/callback", completeAuth)
	app.GET("/music", func(c *gin.Context) {

		topTracks, err := client.CurrentUsersTopTracks(context.Background(), spotify.Timerange("long_term"))
		if err != nil {
			log.Println("Fatal at Tracks")
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": topTracks.Tracks})
	})

	go func() {
		client = <-ch

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
	b := make([]byte, lengthOfState)
	for i := range b {
		b[i] = charsetForState[seededRand.Intn(len(charsetForState))]
	}
	return string(b)
}

package main

import (
	"api/router"
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
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate),
		spotifyauth.WithClientID(""),
		spotifyauth.WithClientSecret(""))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {

	app := gin.Default()

	app.Use(cors.Default())

	router.SetupRouter(app)

	app.GET("/callback", CompleteAuth)

	go func() {
		app.Run(":5000")
	}()

	// first start an HTTP server
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Got request for:", r.URL.String())
	// })
	// go func() {
	// 	err := http.ListenAndServe(":5000", nil)
	// 	if err != nil {
	// 		log.Println("Fatal at Server")
	// 		log.Fatal(err)
	// 	}
	// }()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Println("Fatal at Client")
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

}

func CompleteAuth(c *gin.Context) {
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

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(c, tok))
	fmt.Fprintf(c.Writer, "Login Completed!")
	ch <- client
}

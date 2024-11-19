package main

import (
	"context"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"sort"
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
	auth   *spotifyauth.Authenticator
	ch     = make(chan *spotify.Client)
	state  string
	client *spotify.Client
)

type Track struct {
	Name       string      `json:"name"`
	Artist     string      `json:"artist"`
	ImgURL     string      `json:"imgurl"`
	SpotifyURL string      `json:"spotifyurl"`
	Glimpses   [5][3]uint8 `json:"glimpses"`
}

type Connection struct {
	Tracks []*Track `json:"tracks"`
}

func main() {

	// Gin Init with CORS Middleware
	app := gin.Default()

	app.Use(cors.Default())
	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.GET("/redirect", getRedirectURL)
	app.GET("/callback", completeAuth)
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "healthy"})
	})
	app.GET("/music", func(c *gin.Context) {

		var tracks []Track
		var options []spotify.RequestOption
		var query = c.Request.URL.Query()

		options = append(options, spotify.Timerange(spotify.Range(query["timerange"][0])))
		options = append(options, spotify.Limit(10))

		// fmt.Println(query["timerange"][0])

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
				Glimpses:   getImageColors(track.Album.Images[0].URL),
			})

		}

		c.JSON(http.StatusOK, gin.H{"data": tracks})
	})

	// Link for
	// url := auth.AuthURL(state)
	// fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	app.Run(":5000")

}

func getRedirectURL(c *gin.Context) {

	go clientChannel()

	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead),
		spotifyauth.WithClientID("004a411911e54982b702e657f22c64b2"),
		spotifyauth.WithClientSecret("a50ff8b0428d4cf88e093fd7b50b26c9"))

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

	frontendURL := "https://f219-2603-6010-5303-4994-fd4a-d2f8-c0ae-7b30.ngrok-free.app/" + user.ID

	c.Redirect(http.StatusMovedPermanently, frontendURL)
	ch <- callbackClient
}

func clientChannel() {

	client = <-ch

	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Println("Fatal at Client")
		log.Fatal(err)
	}

	fmt.Println("You are logged in as:", user.ID)

}

func randomString() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lengthOfState)
	for i := range b {
		b[i] = charsetForState[seededRand.Intn(len(charsetForState))]
	}
	return string(b)
}

func getImageColors(path string) [5][3]uint8 {

	response, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal(err)
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := img.Bounds()

	colorOccurrences := make(map[[3]uint8]int)

	for y := 0; y < r.Max.Y; y++ {
		for x := 0; x < r.Max.X; x++ {

			pixel := img.At(x, y)
			color := color.NRGBAModel.Convert(pixel).(color.NRGBA)

			// Originally a lot of colors within the album covers would be various shades of similar colors so
			// the results of the color analysis would not show this. These operations are to protect against that,
			// grouping the occurrences of colors into a shade they are very similar to.
			newR := color.R - (color.R % 15)
			newG := color.G - (color.G % 15)
			newB := color.B - (color.B % 15)
			adjacentShade := [3]uint8{newR, newG, newB}

			value, exists := colorOccurrences[adjacentShade]

			if exists {
				colorOccurrences[adjacentShade] = value + 1
			} else {
				colorOccurrences[adjacentShade] = 1
			}

		}
	}

	keys := make([][3]uint8, 0, len(colorOccurrences))

	for key := range colorOccurrences {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return colorOccurrences[keys[i]] < colorOccurrences[keys[j]]
	})

	return [5][3]uint8{
		keys[len(keys)-1],
		keys[len(keys)-2],
		keys[len(keys)-3],
		keys[len(keys)-4],
		keys[len(keys)-5],
	}

	// fmt.Println("Top Colors")
	// fmt.Println(keys[len(keys)-1])
	// fmt.Println(keys[len(keys)-2])
	// fmt.Println(keys[len(keys)-3])
	// fmt.Println(keys[len(keys)-4])
	// fmt.Println(keys[len(keys)-5])
	// fmt.Println()

}

// func handleArtistNames() string {

// }

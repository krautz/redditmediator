package main

import (
	"fmt"
	"net/http"
	"os"
	"redditmediator/globals"
	"redditmediator/requester"
	"redditmediator/responder"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// get user, password, app id and app secret from environment
	globals.USERNAME = os.Getenv("REDDIT_MEDIATOR_USER_USERNAME")
	password := os.Getenv("REDDIT_MEDIATOR_USER_PASSWORD")
	appID := os.Getenv("REDDIT_MEDIATOR_APP_ID")
	appSecret := os.Getenv("REDDIT_MEDIATOR_APP_SECRET")
	if globals.USERNAME == "" || password == "" || appID == "" || appSecret == "" {
		panic("Please set all environment variables: REDDIT_MEDIATOR_USER_USERNAME, REDDIT_MEDIATOR_USER_PASSWORD, REDDIT_MEDIATOR_APP_ID, REDDIT_MEDIATOR_APP_SECRET")
	}

	// print progression
	fmt.Println("Requesting session token...")

	// authentice user
	globals.TOKEN = requester.GetToken(
		globals.USERNAME,
		password,
		appID,
		appSecret,
	)

	// print token
	fmt.Println("Token:", globals.TOKEN)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// print progression
	fmt.Println("Requesting user's sub reddits...")

	// get user's sub reddits
	globals.SUB_REDDITS = requester.GetSubReddits(
		globals.USERNAME,
		globals.TOKEN,
	)

	// print sub reddits
	for _, subReddit := range globals.SUB_REDDITS {
		fmt.Println("Id:", subReddit.Id)
		fmt.Println("Name:", subReddit.Name)
		fmt.Println("SubReddit:", subReddit.Display_name_prefixed)
		fmt.Println("Title:", subReddit.Title)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	// create mux router and create routes
	router := mux.NewRouter()
	router.HandleFunc("/posts/hot", responder.GET_Posts_Hot)

	// build and start server
	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	server.ListenAndServe()
}

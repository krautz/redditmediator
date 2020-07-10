package main

import (
	"bufio"
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
	// get user, password, app id and app secret from io
	fmt.Println("Insert user, password, appID and appSecret")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	globals.USERNAME = scanner.Text()
	scanner.Scan()
	password := scanner.Text()
	scanner.Scan()
	appID := scanner.Text()
	scanner.Scan()
	appSecret := scanner.Text()

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

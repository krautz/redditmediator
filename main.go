package main

import (
	"bufio"
	"fmt"
	"os"
	"redditmediator/globals"
	"redditmediator/requester"
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

	// get sub reddits' posts
	postControll := make(map[string]requester.PostControll)
	posts := requester.GetPosts(
		globals.USERNAME,
		globals.SUB_REDDITS,
		postControll,
		1,
		"hot",
		globals.TOKEN,
	)

	// print posts
	for _, post := range posts {
		fmt.Println("Id:", post.Id)
		fmt.Println("Name:", post.Name)
		fmt.Println("SubReddit:", post.SubReddit)
		fmt.Println("Title:", post.Title)
		fmt.Println("Type:", post.Type)
		fmt.Println("Content:", post.Content)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

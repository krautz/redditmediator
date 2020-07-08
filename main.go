package main

import (
	"bufio"
	"fmt"
	"os"
	"redditmediator/requester"
)

func main() {
	// get user, password, app id and app token from io
	fmt.Println("Insert user, password, appID and appSecret")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	scanner.Scan()
	password := scanner.Text()
	scanner.Scan()
	appID := scanner.Text()
	scanner.Scan()
	appToken := scanner.Text()

	// print progression
	fmt.Println("Requesting session token...")

	// authentice user
	token := requester.GetToken(username, password, appID, appToken)

	// print token
	fmt.Println("Token:", token)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// print progression
	fmt.Println("Requesting user's sub reddits...")

	// get user's sub reddits
	subReddits := requester.GetSubReddits(username, token)

	// print sub reddits
	for _, subReddit := range subReddits {
		fmt.Println("Id:", subReddit.Id)
		fmt.Println("Name:", subReddit.Name)
		fmt.Println("SubReddit:", subReddit.Display_name_prefixed)
		fmt.Println("Title:", subReddit.Title)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	// print progression
	fmt.Println("Requesting user's sub reddits' posts...")

	// get sub reddits' posts
	postControll := make(map[string]requester.PostControll)
	posts := requester.GetPosts(username, subReddits, postControll, 1, "hot", token)

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

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

	// authentice user
	token := requester.GetToken(username, password, appID, appToken)

	// get user's sub reddits
	subReddits := requester.GetSubReddits(username, token)

	// print token
	fmt.Println(subReddits, len(subReddits))
}

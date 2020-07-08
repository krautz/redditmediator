/*
 * requests user sub reddits
 */

package requester

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
 * STRUCTS
 */
type SubReddit struct {
	Title, Name, Display_name_prefixed, Id string
}

type SubRedditResponseDataChildren struct {
	Data SubReddit
}

type SubRedditResponseData struct {
	Children      []SubRedditResponseDataChildren
	After, Before string
}

type SubRedditResponse struct {
	Data SubRedditResponseData
}

/*
 * I get user's sub reddits.
 *
 * username -> logged user username
 * token    -> user session token
 *
 * returns -> user sub reddits
 */
func GetSubReddits(
	username string,
	token string,
) []SubReddit {

	// create function return value
	var subReddits []SubReddit

	// create http client (so it is abble to set requests timeout to 30s)
	var client = &http.Client{Timeout: 30 * time.Second}

	// create client base url and first request url (25 items per request)
	baseURL := "https://oauth.reddit.com/subreddits/mine/subscriber?limit=25"
	baseURL += "&raw_json=1"
	url := baseURL

	// initialize loop variables
	after := "start"
	count := 0

	// request all sub reddits
	for after != "" {

		// create request object
		request, err := http.NewRequest("GET", url, nil)

		// error while creating the request -> panic
		if err != nil {
			panic(err)
		}

		// set request authorization and user agent headers
		request.Header.Set("Authorization", "bearer "+token)
		request.Header.Set("User-Agent", "RedditMediator/0.1 by "+username)

		// do the request
		response, err := client.Do(request)

		// error while making the request -> panic
		if err != nil {
			panic(err)
		}

		// defer close of request after function finish
		defer response.Body.Close()

		// parse request response (and treat errors)
		subRedditResponse := SubRedditResponse{}
		responseBody, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			panic(readErr)
		}
		readErr = json.Unmarshal(responseBody, &subRedditResponse)
		if readErr != nil {
			panic(readErr)
		}

		// increment count and retrieve after
		count += len(subRedditResponse.Data.Children)
		after = subRedditResponse.Data.After

		// update next request url
		url = baseURL + "&count=" + strconv.Itoa(count) + "&after=" + after

		// add each sub reddit to the response (ignore users)
		for _, subReddit := range subRedditResponse.Data.Children {
			if subReddit.Data.Display_name_prefixed[0] == 'r' {
				subReddits = append(subReddits, subReddit.Data)
			}
		}
	}

	// return user's sub reddits
	return subReddits
}

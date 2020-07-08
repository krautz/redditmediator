/*
 * requests user sub reddits
 */

package requester

import (
	"encoding/json"
	"strconv"
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

	// create client base url and first request url (25 items per request)
	baseURL := "https://oauth.reddit.com/subreddits/mine/subscriber?limit=25"
	baseURL += "&raw_json=1"
	url := baseURL

	// initialize loop variables
	after := "start"
	count := 0

	// request all sub reddits
	for after != "" {

		// request sub reddits
		response := request("GET", url, nil, token, username)

		// load request response into json
		JSONResponse := SubRedditResponse{}
		readErr := json.Unmarshal(response, &JSONResponse)
		if readErr != nil {
			panic(readErr)
		}

		// increment count and retrieve after
		count += len(JSONResponse.Data.Children)
		after = JSONResponse.Data.After

		// update next request url
		url = baseURL + "&count=" + strconv.Itoa(count) + "&after=" + after

		// add each sub reddit to the response (ignore users)
		for _, subReddit := range JSONResponse.Data.Children {
			if subReddit.Data.Display_name_prefixed[0] == 'r' {
				subReddits = append(subReddits, subReddit.Data)
			}
		}
	}

	// return user's sub reddits
	return subReddits
}

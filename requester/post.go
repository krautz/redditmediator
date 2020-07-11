/*
 * requests sub reddits posts
 */

package requester

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/*
 * STRUCTS
 */
type PostControll struct {
	After string
	Count int
}

type Post struct {
	Title, Name, Id, Type, Content, SubReddit string
}

type PostMediaVideo struct {
	Fallback_url string
	Is_gif       bool
}

type PostMedia struct {
	Reddit_video PostMediaVideo
}

type PostResponse struct {
	Title, Name, Id, Post_hint, Url_overridden_by_dest string
	Url, Selftext, Subreddit_name_prefixed             string
	Is_self                                            bool
	Media                                              PostMedia
	Crosspost_parent_list                              []PostResponse
}

type PostsResponseDataChildren struct {
	Data PostResponse
}

type PostsResponseData struct {
	Children      []PostsResponseDataChildren
	After, Before string
}

type PostsResponse struct {
	Data PostsResponseData
}

/*
 * I get sub reddits' posts.
 *
 * username   -> logged user username
 * subReddits -> sub reddits to get the posts from
 * limit      -> ammount of posts to get from each sub reddit
 * sortType   -> how to get posts: by "hot" or "new"
 * token      -> user session token
 *
 * returns -> sub reddits' posts
 */
func GetPosts(
	username string,
	subReddits []SubReddit,
	postControll map[string]PostControll,
	limit int,
	sortType string,
	token string,
) []Post {

	// analyze sort type
	if sortType != "hot" && sortType != "new" {
		fmt.Println("Invalid sort type. It must be \"hot\" or \"new\"")
		return []Post{}
	}

	// create function return value
	var posts []Post

	// create client base url
	baseURL := "https://oauth.reddit.com/"

	// request posts from all sub reddits
	for _, subReddit := range subReddits {

		// set initial url with limit query param
		url := baseURL + subReddit.Display_name_prefixed + "/" + sortType
		url += "?limit=" + strconv.Itoa(limit) + "&raw_json=1"

		// sub reddit already queried before -> add progression to the url
		count := 0
		if _, ok := postControll[subReddit.Id]; ok {
			after := postControll[subReddit.Id].After
			count = postControll[subReddit.Id].Count
			url = url + "&after=" + after + "&count=" + strconv.Itoa(count)
		}

		// request posts. Treat request errors
		response, err := request("GET", url, nil, token, username)
		if err != nil {
			fmt.Println("Error while requesting", url, ":", err)
			continue
		}

		// load request response into json. Treat parse errors
		JSONResponse := PostsResponse{}
		readErr := json.Unmarshal(response, &JSONResponse)
		if readErr != nil {
			fmt.Println("Error while parsing request response:", readErr)
			continue
		}

		// increment count and retrieve after on the post controll
		increseCount := len(JSONResponse.Data.Children)
		postControll[subReddit.Id] = PostControll{
			Count: count + increseCount,
			After: JSONResponse.Data.After,
		}

		// add each post to the response
		for _, postChildren := range JSONResponse.Data.Children {

			// process post
			newPost := processPost(postChildren.Data)

			// post is valid -> add it to return posts
			if newPost.Title != "" {
				posts = append(posts, newPost)
			}
		}
	}

	// return sub reddits' posts
	return posts
}

/*
 * requests sub reddits posts
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
		panic("Invalid sort type")
	}

	// create function return value
	var posts []Post

	// create http client (so it is abble to set requests timeout to 30s)
	var client = &http.Client{Timeout: 30 * time.Second}

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
		postsResponse := PostsResponse{}
		responseBody, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			panic(readErr)
		}
		readErr = json.Unmarshal(responseBody, &postsResponse)
		if readErr != nil {
			panic(readErr)
		}

		// increment count and retrieve after on the post controll
		increseCount := len(postsResponse.Data.Children)
		postControll[subReddit.Id] = PostControll{
			Count: count + increseCount,
			After: postsResponse.Data.After,
		}

		// add each post to the response
		for _, postChildren := range postsResponse.Data.Children {

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

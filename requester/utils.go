/*
 * requests utils
 */

package requester

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
 * GLOBALS
 */
// create http client (so it is abble to set requests timeout to 30s)
var client = &http.Client{Timeout: 30 * time.Second}

/*
 * I process a post.
 *
 * post -> post to be processed
 *
 * returns -> formatted post (if post is supported)
 */
func processPost(post PostResponse) Post {

	// set new post
	newPost := Post{
		Title:     post.Title,
		Name:      post.Name,
		Id:        post.Id,
		SubReddit: post.Subreddit_name_prefixed,
	}

	// add content and type to the post
	// decisions made here can be found at redditapi.txt
	if post.Subreddit_name_prefixed == "r/announcements" {
		return Post{}
	} else if post.Is_self == true {
		// simple text post -> set content with post text
		newPost.Type = "text"
		newPost.Content = post.Selftext
	} else if post.Post_hint == "image" {
		// image post -> set content with image url
		newPost.Type = "image"
		newPost.Content = post.Url_overridden_by_dest
	} else if post.Post_hint == "hosted:video" {
		if post.Media.Reddit_video.Is_gif == true {
			// post with gif uploaded to reddit server -> set type as gif
			newPost.Type = "gif"
		} else {
			// post with video uploaded to reddit server -> set type as video
			newPost.Type = "video"
		}
		// post with gif/video uploaded to reddit server -> set content with url
		newPost.Content = post.Media.Reddit_video.Fallback_url
	} else if post.Post_hint == "rich:video" {
		// post with video uploaded off reddit server -> set type as link
		// and content with link to external website
		newPost.Type = "link"
		newPost.Content = post.Url
	} else if post.Post_hint == "link" {
		if len(post.Crosspost_parent_list) == 0 {
			// post has a link to another website -> set type as link and
			// add url to the outside website as content
			newPost.Type = "link"
			newPost.Content = post.Url_overridden_by_dest
		} else {
			// post has a link to another reddit post -> set type as the
			// reposted content and add info from repost to title
			crossPost := processPost(post.Crosspost_parent_list[0])
			newPost.Type = crossPost.Type
			newPost.Content = crossPost.Content
			titlePrefix := "[REPOST FROM " + crossPost.SubReddit + ": "
			titlePrefix += crossPost.Title + "] "
			newPost.Title = titlePrefix + newPost.Title
		}
	} else if post.Url_overridden_by_dest != "" {
		// post has no info on its type -> guess it is link (empiric check)
		newPost.Type = "link"
		newPost.Content = post.Url_overridden_by_dest
	} else {
		// unrecognized post -> log for further check
		fmt.Println("Unsupported post received:", post)
		return Post{}
	}

	// return formatted post
	return newPost
}

/*
 * I do a request.
 *
 * method   -> request rest API method
 * url      -> where to send the request to
 * body     -> request body (nil if no body needed)
 * token    -> user session token
 * username -> logged user username
 *
 * returns -> request response bytes
 */
func request(
	method string,
	url string,
	body *strings.Reader,
	token string,
	username string,
) []byte {

	// create request object
	var request *http.Request
	var err error
	if body == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, body)
	}

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
	responseBody, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		panic(readErr)
	}

	// return bytes of request response
	return responseBody
}

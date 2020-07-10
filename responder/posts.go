package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redditmediator/globals"
	"redditmediator/requester"
	"strconv"
)

/*
 * GLOBALS
 */
var HotPostsControll = make(map[string]requester.PostControll)

/*
 * STRUCTS
 */
type PostsResponse struct {
	Data []requester.Post
}

/*
 * I answer a request to get user sub reddit's posts.
 *
 * w -> request response writter
 * r -> request
 *
 * returns -> nothing
 */
func GET_Posts_Hot(w http.ResponseWriter, r *http.Request) {
	// get numbers of posts of each sub reddit to retrieve
	query := r.URL.Query()
	limit := query.Get("limit")
	if limit == "" {
		limit = "3"
	}
	// TODO: threat errors
	limitInt, err := strconv.Atoi(limit)

	// print progression
	fmt.Println("Requesting " + limit + " hot posts of each user's sub reddits")

	// get sub reddits' hot posts
	posts := requester.GetPosts(
		globals.USERNAME,
		globals.SUB_REDDITS,
		HotPostsControll,
		limitInt,
		"hot",
		globals.TOKEN,
	)

	// respond request
	w.Header().Set("Content-Type", "application/json")
	response := PostsResponse{posts}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	w.Write(responseJSON)
	fmt.Println("Answered user's sub reddits' hot posts")
}

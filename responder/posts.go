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
var NewPostsControll = make(map[string]requester.PostControll)

/*
 * STRUCTS
 */
type PostsResponse struct {
	Data []requester.Post
}

type FailureResponse struct {
	Data Error
}

type Error struct {
	Error string
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
	// print received  request
	fmt.Println("Received request to get hot posts of each user's sub reddits")

	// get numbers of posts of each sub reddit to retrieve
	query := r.URL.Query()
	limit := query.Get("limit")
	if limit == "" {
		limit = "3"
	}

	// convert limit to int. In case of error, log and fail request
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println("Failing request:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			FailureResponse{Error{"Invalid limit. It must be an integer"}},
		)
		return
	}

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

	// respond request. If fails to encode response fail request
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(PostsResponse{posts})
	if err != nil {
		fmt.Println("Failing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			FailureResponse{Error{"Failed to parse reddit response."}},
		)
		return
	}
	fmt.Println("Answered user's sub reddits' new posts")
}

/*
 * I answer a request to get user sub reddit's new posts.
 *
 * w -> request response writter
 * r -> request
 *
 * returns -> nothing
 */
func GET_Posts_New(w http.ResponseWriter, r *http.Request) {
	// print received  request
	fmt.Println("Received request to get new posts of each user's sub reddits")

	// get numbers of posts of each sub reddit to retrieve
	query := r.URL.Query()
	limit := query.Get("limit")
	if limit == "" {
		limit = "3"
	}

	// convert limit to int. In case of error, log and fail request
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println("Failing request:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			FailureResponse{Error{"Invalid limit. It must be an integer"}},
		)
		return
	}

	// print progression
	fmt.Println("Requesting " + limit + " new posts of each user's sub reddits")

	// get sub reddits' new posts
	posts := requester.GetPosts(
		globals.USERNAME,
		globals.SUB_REDDITS,
		NewPostsControll,
		limitInt,
		"new",
		globals.TOKEN,
	)

	// respond request. If fails to encode response fail request
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(PostsResponse{posts})
	if err != nil {
		fmt.Println("Failing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			FailureResponse{Error{"Failed to parse reddit response."}},
		)
		return
	}
	fmt.Println("Answered user's sub reddits' new posts")
}

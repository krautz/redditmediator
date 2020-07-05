/*
 * requests user session token and return it
 */

package requester

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
 * STRUCTS
 */
type TokenResponse struct {
	Access_token, Error string
}

/*
 * I get an user token.
 *
 * username  -> username of the user to get the token
 * password  -> password of the user to get the token
 * appID     -> Reddit app id which the user is subscribed to
 * appSecret -> Reddit app secret which the user is subscribed to
 *
 * returns -> user session token (lasts 1 hour)
 */
func GetToken(
	username string,
	password string,
	appID string,
	appSecret string,
) string {
	// set POST body
	body := strings.NewReader(
		"grant_type=password&username=" + username + "&password=" + password,
	)

	// create request objet
	request, err := http.NewRequest(
		"POST",
		"https://www.reddit.com/api/v1/access_token",
		body,
	)

	// error while creating the request -> log error and panic
	if err != nil {
		panic(err)
	}

	// set request basic auth and user agent headers
	request.SetBasicAuth(appID, appSecret)
	request.Header.Set("User-Agent", "RedditMediator/0.1 by "+username)

	// create http client (so it is abble to set requests timeout to 30s)
	var client = &http.Client{Timeout: 30 * time.Second}

	// do the request
	response, err := client.Do(request)

	// error while making the request -> panic
	if err != nil {
		panic(err)
	}

	// defer close of request after function finish
	defer response.Body.Close()

	// parse request response (and treat errors)
	tokenResponse := TokenResponse{}
	responseBody, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		panic(readErr)
	}
	// TODO: if this fails due to error being int, parse it into new structure
	readErr = json.Unmarshal(responseBody, &tokenResponse)
	if readErr != nil {
		panic(readErr)
	}

	// request failed -> panic
	if tokenResponse.Error != "" {
		panic(tokenResponse.Error)
	}

	// return acess token
	return tokenResponse.Access_token
}

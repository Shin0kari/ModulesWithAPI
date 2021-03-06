package jokes

import (
	"fmt"
	"net/http"

	"encoding/json"

	"workshop/internal/api"
)

const getJokePath = "https://geek-jokes.sameerkumar.website/api?format=json"

// JokeClient is a joke Api client
type JokeClient struct {
	url string
}

// NewJokeClient create a new joke client
func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

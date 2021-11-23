package api

// JokeResponse is joke API response type
type JokeResponse struct {
	// ответ, который приходит
	Joke string `json:"joke"`
}

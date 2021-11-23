package api

// сделан интерфейс, окотрый описывает ApiClient
// в нём есть метод GetJoke, который возвращает
// либо ошибку либо ответ. ответ описан в models

// Client interacts with 3-rd party joke API
type Client interface {
	// GetJoke return one joke
	GetJoke() (*JokeResponse, error)
}

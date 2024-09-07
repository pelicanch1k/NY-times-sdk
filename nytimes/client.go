package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"NY-times-sdk/nytimes/book"
	"NY-times-sdk/nytimes/view"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

const (
	BOOK = "https://api.nytimes.com/svc/books/v3/lists.json?list=%s&api-key=%s"
	VIEW = "https://api.nytimes.com/svc/mostpopular/v2/viewed/%s.json?api-key=%s"
)

type Client struct {
	client  *http.Client
	api_key string
}

func NewClient(timeout time.Duration, api_key string) *Client {
	api_key, exists := os.LookupEnv(api_key)
	if exists {
		return &Client{
			client: &http.Client{
				Timeout: timeout,
			},
			api_key: api_key,
		}
	}

	panic(fmt.Sprintf("No variable %s", api_key))
}

func (c *Client) doHTTP(url string) ([]byte, error) {
	resp, err := c.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getAssets[T Answer](c *Client, url string) (*T, error) {
	body, err := c.doHTTP(url)
	if err != nil {
		panic(err)
	}

	answer, err := unmarshalBody[T](body)
	return answer, err
}

func unmarshalBody[T Answer](body []byte) (*T, error) {
	var answer T

	if err := json.Unmarshal(body, &answer); err != nil {
		panic(err)
	}

	return &answer, nil
}

func (c *Client) GetAssetsBook(list string) (*book.AnswerBook, error) {
	url := fmt.Sprintf(BOOK, list, c.api_key)

	answer, err := getAssets[book.AnswerBook](c, url)

	return answer, err
}

func (c *Client) GetAssetsMostViewed(period string) (*view.AnswerView, error) {
	url := fmt.Sprintf(VIEW, period, c.api_key)

	answer, err := getAssets[view.AnswerView](c, url)

	return answer, err
}

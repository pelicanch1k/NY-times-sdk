package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (b *Book) Info() string {
	return fmt.Sprintf("%s | %s | %s\n", b.Author, b.Title, b.Description)
}

type Results struct {
	Book_details     []Book
	Bestsellers_date string `json:"bestsellers_date"`
	Published_date   string `json:"published_date"`
}

type Post struct {
	Results []Results `json:"results"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	// client := &http.Client{
	// 	Timeout: 30 * time.Second,
	// }

	api_key, _ := os.LookupEnv("API_KEY_NY_TIMES")
	println(api_key)

	// url := fmt.Sprintf("https://api.nytimes.com/svc/books/v3/lists.json?list=hardcover-fiction&api-key=%s", api_key)

	// resp, err := client.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// var r Post
	// if err = json.Unmarshal(body, &r); err != nil {
	// 	panic(err)
	// }

	// for _, a := range r.Results {
	// 	for _, b := range a.Book_details {
	// 		println(b.Info())
	// 	}
	// }
}

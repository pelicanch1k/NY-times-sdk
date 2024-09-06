package client

import "fmt"

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

type Answer struct {
	Results []Results `json:"results"`
}

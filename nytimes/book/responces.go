package book

import "fmt"

type book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (b *book) Info() string {
	return fmt.Sprintf("%s | %s | %s\n", b.Author, b.Title, b.Description)
}

type results struct {
	Book_details     []book
	Bestsellers_date string `json:"bestsellers_date"`
	Published_date   string `json:"published_date"`
}

type AnswerBook struct {
	Results []results `json:"results"`
}

func (a *AnswerBook) Info() {
	for _, results := range a.Results {
		for _, Book_details := range results.Book_details {
			fmt.Println(Book_details.Info())
		}
	}
}

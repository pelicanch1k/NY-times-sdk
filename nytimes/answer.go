package client

import (
	"NY-times-sdk/nytimes/book"
	"NY-times-sdk/nytimes/view"
)

type Answer interface {
	book.AnswerBook | view.AnswerView
}

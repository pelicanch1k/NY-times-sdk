<h1 align="center">NY-times-sdk</h1>

<p align="center">
   <img alt="Static Badge" src="https://img.shields.io/badge/Version-v1.0-blue?style=flat&color=blueviolet">
   <img alt="Static Badge" src="https://img.shields.io/badge/License-MIT-green?style=flat">
</p>

## About

Test api for https://developer.nytimes.com/

## Documentation
``` go
...
  GetAssetsBook(list string) (*book.AnswerBook, error)
  GetAssetsMostViewed(period string) (*view.AnswerView, error)
...
```

## Example usage:

```golang
package main

import (
	client "NY-times-sdk/nytimes"
	"time"
)

const (
	API_KEY = "API_KEY_NY_TIMES"
)

func main() {
	client := client.NewClient(time.Second*6, API_KEY)

	answer, _ := client.GetAssetsBook("hardcover-fiction")
	answer.Info()
}
```
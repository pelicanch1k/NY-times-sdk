package view

type view struct {
	Url            string `json:"url"`
	Asset_id       string `json:"asset_id"`
	Source         string `json:"source"`
	Published_date string `json:"published_date"`
	Updated        string `json:"updated"`
	Byline         string `json:"byline"`
	Title          string `json:"title"`
}

type results struct {
	View_details []view
}

type AnswerView struct {
	Results []results `json:"results"`
}

func (a *AnswerView) Info() {
	// Проверяем, есть ли данные для вывода
	if len(a.Results) == 0 {
		println("Нет доступных данных для отображения.")
		return
	}

	// Итерируемся по результатам
	for _, result := range a.Results {
		for _, viewDetail := range result.View_details {
			// Выводим необходимые данные из структуры view
			println("URL:", viewDetail.Url)
			println("Asset ID:", viewDetail.Asset_id)
			println("Source:", viewDetail.Source)
			println("Published Date:", viewDetail.Published_date)
			println("Updated:", viewDetail.Updated)
			println("Byline:", viewDetail.Byline)
			println("Title:", viewDetail.Title)
			println() // Пустая строка для удобства
		}
	}
}

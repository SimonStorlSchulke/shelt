package animal

import "time"

type StrapiDataAnimal struct {
	Data struct {
		Attributes struct {
			Description       string    `json:"Description"`
			Gender            string    `json:"Gender"`
			Name              string    `json:"Name"`
			Race              string    `json:"Race"`
			Species           string    `json:"Species"`
			Birthdate         string    `json:"birthdate"`
			Castrated         bool      `json:"castrated"`
			CreatedAt         time.Time `json:"createdAt"`
			CustomArticleLink any       `json:"custom_article_link"`
			PublishedAt       time.Time `json:"publishedAt"`
			Shoulderheight    int       `json:"shoulderheight"`
			UpdatedAt         time.Time `json:"updatedAt"`
		} `json:"attributes"`
		ID int `json:"id"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}

type Animal struct {
	Data struct {
		Attributes struct {
			Description       string    `json:"Description"`
			Gender            string    `json:"Gender"`
			Name              string    `json:"Name"`
			Race              string    `json:"Race"`
			Species           string    `json:"Species"`
			Birthdate         string    `json:"birthdate"`
			Castrated         bool      `json:"castrated"`
			CreatedAt         time.Time `json:"createdAt"`
			CustomArticleLink any       `json:"custom_article_link"`
			PublishedAt       time.Time `json:"publishedAt"`
			Shoulderheight    int       `json:"shoulderheight"`
			UpdatedAt         time.Time `json:"updatedAt"`
		} `json:"attributes"`
		ID int `json:"id"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}

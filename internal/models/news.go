package models

import "time"

type News struct {
	TotalArticles int       `json:"totalArticles"`
	Articles      []Article `json:"articles"`
}

type Article struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	PublishedAt time.Time `json:"publishedAt"`
	Lang        string    `json:"lang"`
	Source      Source    `json:"source"`
}

type Source struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Country string `json:"country"`
}

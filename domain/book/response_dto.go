package book

import "time"

type BooksDto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}

type BooksItemDto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Headline string `json:"headline"`
	Tag      string `json:"tag"`
}

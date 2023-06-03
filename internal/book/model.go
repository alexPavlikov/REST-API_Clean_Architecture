package book

import "github.com/alexPavlikov/REST-API_Clean_Architecture/internal/author"

type Book struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Author []author.Author `json:"author"`
}

type Author struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

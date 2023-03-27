package presentation

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Descr  string `json:"description"`
}

var Books = []Book{}

package endpoint

type BookRequest struct {
	ID int64 `json:"id"`
}

type AddBookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type BookResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type BooksRequest struct {
	IDs []int64 `json:"id"`
}

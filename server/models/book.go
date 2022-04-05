package models

// Book
type Book struct {
	ID        int64   `db:"book_id"`
	Price     float32 `db:"price"`
	Pages     int32   `db:"pages"`
	Title     string  `db:"title"`
	Author    string  `db:"author"`
	Publisher string  `db:"publisher"`
	ISBN      string  `db:"book_isbn"`
}

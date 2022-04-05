package mysql

import (
	"database/sql"
	"fmt"
	"server/models"

	"go.uber.org/zap"
)

// 根据ID获取对应的Book
func QueryByID(ID int64) (book *models.Book, err error) {
	sqlStr := "select * from `book` where `book_id` = ?"
	book = new(models.Book)

	err = db.Get(book, sqlStr, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community record")
		} else {
			zap.L().Error("Select syntax error")
		}
		return nil, err
	}

	return book, nil
}

// 根据ISBN获取对应的Book
func QueryByISBN(ISBN string) (book *models.Book, err error) {
	sqlStr := "select * from `book` where `book_id` = ?"
	book = new(models.Book)

	err = db.Get(book, sqlStr, ISBN)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community record")
		} else {
			zap.L().Error("Select syntax error")
		}
		return nil, err
	}

	return book, nil
}

// 根据书名模糊匹配
func QueryByName(name string) (books []models.Book, err error) {
	sqlStr := fmt.Sprintf("select * from `book` where `title` like '%%%s%%'", name)
	err = db.Select(&books, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("There is no community record")
		} else {
			zap.L().Error("Select syntax error")
		}
		return nil, err
	}

	return books, nil
}

// 加入新书籍
func InsertBook(book *models.Book) (code int, err error) {
	sqlStr := "insert into `book` values(?,?,?,?,?,?,?)"

	_, err = db.Exec(sqlStr, book.ID, book.Price, book.Pages, book.Title, book.Author, book.Publisher, book.ISBN)
	if err != nil {
		zap.L().Error("Syntax error")
		return -1, err
	}

	return 1, nil
}

// 删除书籍
func DeleteBookByID(ID int64) (code int, err error) {
	sqlStr := "delete from `book` where `book_id`=?"

	_, err = db.Exec(sqlStr, ID)
	if err != nil {
		zap.L().Error("Syntax error")
		return -1, err
	}

	return 1, nil
}

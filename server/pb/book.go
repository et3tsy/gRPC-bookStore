package service

import (
	context "context"
	"server/dao/mysql"
	"server/models"

	"go.uber.org/zap"
)

// 数据服务类,实现接口方法
type BookService struct {
}

// 插入新的书籍
func (b *BookService) Add(_ context.Context, book *ParamBook) (*ParamCode, error) {
	code, err := mysql.InsertBook(ConvertParamToBook(book))
	if err != nil {
		zap.L().Error("Insert Fails")
		return &ParamCode{Code: int32(code)}, err
	}

	return &ParamCode{Code: int32(code)}, nil
}

// 根据 Id 查询书籍
func (b *BookService) QueryByID(_ context.Context, p *ParamInt) (*ParamBook, error) {
	book, err := mysql.QueryByID(p.Param)
	if err != nil {
		zap.L().Error("Query By ID Fails")
		return nil, err
	}

	return ConvertBookToParam(book), nil
}

// 根据书籍名字模糊查询书籍
func (b *BookService) QueryByName(_ context.Context, p *ParamString) (*ParamBookList, error) {
	bookList, err := mysql.QueryByName(p.Param)
	if err != nil {
		zap.L().Error("Query By Name Fails")
		return nil, err
	}

	return ConvertBookListToParam(bookList), nil
}

// 根据 Id 删除书籍
func (b *BookService) DeleteById(_ context.Context, p *ParamInt) (*ParamCode, error) {
	code, err := mysql.DeleteBookByID(p.Param)
	if err != nil {
		zap.L().Error("Delete Fails")
		return &ParamCode{Code: int32(code)}, err
	}

	return &ParamCode{Code: int32(code)}, nil
}

// 将 ParamBook 转化成 Book
func ConvertParamToBook(book *ParamBook) *models.Book {
	return &models.Book{
		ID:        book.Id,
		Price:     book.Price,
		Pages:     book.Pages,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		ISBN:      book.ISBN,
	}
}

// 将 Book 转化成 ParamBook
func ConvertBookToParam(book *models.Book) *ParamBook {
	return &ParamBook{
		Id:        book.ID,
		Price:     book.Price,
		Pages:     book.Pages,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		ISBN:      book.ISBN,
	}
}

// 将 BookList 转化成 ParamBookList
func ConvertBookListToParam(bookList []models.Book) (p *ParamBookList) {
	p = new(ParamBookList)
	p.BookList = make([]*ParamBook, 0)
	for _, j := range bookList {
		p.BookList = append(p.BookList, ConvertBookToParam(&j))
	}
	return p
}

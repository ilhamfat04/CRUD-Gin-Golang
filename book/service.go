package book

type Service interface {
	GetBooks() ([]Book, error)
	GetBook(ID int) (Book, error)
	AddBook(book Book) (Book, error)
	UpdateBook(ID int) (Book, error)
}

type service struct {
	repository Repository
}

package book

type Service interface {
	GetBooks() ([]Book, error)
	GetBook(ID int) (Book, error)
	AddBook(bookInput BookInput) (Book, error)
	// UpdateBook(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks() ([]Book, error) {
	books, err := s.repository.GetBooks()
	return books, err
}

func (s *service) GetBook(ID int) (Book, error) {
	book, err := s.repository.GetBook(ID)
	return book, err
}

func (s *service) AddBook(bookInput BookInput) (Book, error) { // BookInput dari struct book.go
	book := Book{
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Price:       bookInput.Price,
	}
	newBook, err := s.repository.AddBook(book)

	return newBook, err
}

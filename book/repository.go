package book

import "gorm.io/gorm"

type Repository interface {
	GetBooks() ([]Book, error)
	GetBook(ID int) (Book, error)
	AddBook(book Book) (Book, error)
	UpdateBook(ID int) (Book, error)
}

type repository struct {
	db *gorm.DB
}

// func nama-func(paramas) nilai-yg-dikembalikan
func NewRepository(db *gorm.DB) *repository {
	return &repository{db} // agar bisa diakses di main
}

// func (struct) nama-func(paramas) nilai-yg-dikembalikan
func (r *repository) GetBooks() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) GetBook(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) AddBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) UpdateBook(book Book) error {
	err := r.db.Save(&book).Error
	return err
}

func (r *repository) DeleteBook(book Book) error {
	err := r.db.Delete(&book).Error
	return err
}

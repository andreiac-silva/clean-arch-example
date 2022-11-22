package book

import "clean-arch/domain"

type Repository interface {
	Save(book domain.Book) (string, error)
	Update(book domain.Book) error
	Delete(id string) error
	FindOne(id string) (domain.Book, error)
	FindAll() ([]domain.Book, error)
}

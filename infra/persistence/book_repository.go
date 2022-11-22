package persistence

import "clean-arch/domain"

type BookRepository struct {
	// Omitted, just a fake repository to demonstrate Clean Arch model.
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (b BookRepository) Save(book domain.Book) (string, error) {
	// TODO
	return "", nil
}

func (b BookRepository) Update(book domain.Book) error {
	// TODO
	return nil
}

func (b BookRepository) Delete(id string) error {
	// TODO
	return nil
}

func (b BookRepository) FindOne(id string) (domain.Book, error) {
	// TODO
	return domain.Book{}, nil
}

func (b BookRepository) FindAll() ([]domain.Book, error) {
	// TODO
	return []domain.Book{}, nil
}

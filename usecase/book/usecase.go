package book

import (
	"clean-arch/domain"
)

type UseCase interface {
	Create(e domain.Book) (string, error)
	Update(e domain.Book) error
	Delete(id string) error
	GetByID(id string) (domain.Book, error)
	GetAll() ([]domain.Book, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) *useCase {
	return &useCase{
		repo: r,
	}
}

func (s useCase) Create(book domain.Book) (string, error) {
	return s.repo.Save(book)
}

func (s useCase) Update(book domain.Book) error {
	return s.repo.Update(book)
}

func (s useCase) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s useCase) GetByID(id string) (domain.Book, error) {
	return s.repo.FindOne(id)
}

func (s useCase) GetAll() ([]domain.Book, error) {
	return s.repo.FindAll()
}

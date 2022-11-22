package user

import "clean-arch/domain"

type UseCase interface {
	Create(e domain.User) (string, error)
	Update(e domain.User) error
	Delete(id string) error
	GetByID(id string) (domain.User, error)
	GetAll() ([]domain.User, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) *useCase {
	return &useCase{
		repo: r,
	}
}

func (s useCase) Create(user domain.User) (string, error) {
	return s.repo.Save(user)
}

func (s useCase) Update(user domain.User) error {
	return s.repo.Update(user)
}

func (s useCase) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s useCase) GetByID(id string) (domain.User, error) {
	return s.repo.FindOne(id)
}

func (s useCase) GetAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

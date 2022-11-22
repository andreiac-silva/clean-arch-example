package user

import "clean-arch/domain"

type Repository interface {
	Save(user domain.User) (string, error)
	Update(user domain.User) error
	Delete(id string) error
	FindOne(id string) (domain.User, error)
	FindAll() ([]domain.User, error)
}

package persistence

import "clean-arch/domain"

type UserRepository struct {
	// Omitted, just a fake repository to demonstrate Clean Arch model.
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (b UserRepository) Save(domain.User) (string, error) {
	// TODO
	return "", nil
}

func (b UserRepository) Update(user domain.User) error {
	// TODO
	return nil
}

func (b UserRepository) Delete(id string) error {
	// TODO
	return nil
}

func (b UserRepository) FindOne(id string) (domain.User, error) {
	// TODO
	return domain.User{}, nil
}

func (b UserRepository) FindAll() ([]domain.User, error) {
	// TODO
	return []domain.User{}, nil
}

package user

import (
	"github.com/ramonrune/assina-backend/internal/app/domain/user"
)

type UserRepository struct {
}

func (r *UserRepository) Create(c *user.User) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) FindById(id string) (*user.User, error) {

	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(c *user.User) (*user.User, error) {

	return nil, nil
}

func (r *UserRepository) Delete(id string) error {
	return nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

package user

import "time"

type UserService struct {
	repository Repository
}

func (s *UserService) Create(name string, email string, password string) (*User, error) {

	user := NewUser(name, email, password)

	res, err := s.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetByID(id string) (*User, error) {

	user, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(email string) (*User, error) {

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id string, name string) (*User, error) {

	user, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	user.Name = name
	user.UpdatedAt = time.Now()

	user, err = s.repository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}

func NewUserService(repository Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

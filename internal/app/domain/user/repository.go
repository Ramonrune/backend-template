package user

type Repository interface {
	Create(user *User) (*User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
}

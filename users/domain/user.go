package domain

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Save(user *User) error
}

package service

type User struct {
}

func NewUserService() User {
	service := User{}

	return service
}

func (s User) Create() error {
	return nil
}

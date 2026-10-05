package users

import (
	"backend-golang/domain"
	userHandler "backend-golang/rest/handler/users"
)


type UserService interface {
	userHandler.Services
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Get(id int) (*domain.User, error)
	List() []*domain.User
	Update(u domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
}

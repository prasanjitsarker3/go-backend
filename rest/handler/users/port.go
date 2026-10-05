package users

import "backend-golang/domain"



type Services interface {
	Create(u domain.User) (*domain.User, error)
	Get(id int) (*domain.User, error)
	List() []*domain.User
	Update(u domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
}		
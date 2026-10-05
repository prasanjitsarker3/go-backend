package users

import "backend-golang/domain"

type Service struct {
	repo UserRepo
}

func NewService(repo UserRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(u domain.User) (*domain.User, error) {
	return s.repo.Create(u)
}

func (s *Service) Get(id int) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *Service) List() []*domain.User {
	return s.repo.List()
}

func (s *Service) Update(u domain.User) (*domain.User, error) {
	return s.repo.Update(u)
}

func (s *Service) Delete(id int) (bool, error) {
	return s.repo.Delete(id)
}

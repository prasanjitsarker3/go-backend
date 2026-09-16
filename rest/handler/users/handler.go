package users

import "backend-golang/repo"

type Handler struct {
	repo repo.UserRepo
}

func NewHandler(repo repo.UserRepo) *Handler {
	return &Handler{repo: repo}
}

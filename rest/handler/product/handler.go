package product

import "backend-golang/repo"

type Handler struct {
	repo repo.ProductRepo
}

func NewHandler(repo repo.ProductRepo) *Handler {
	return &Handler{repo: repo}
}

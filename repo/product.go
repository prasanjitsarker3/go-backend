package repo

import "fmt"

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() []*Product
	Update(p Product) (*Product, error)
	Delete(id int) (bool, error)
}

type productRepo struct {
	productList [] * Product
}

//Constructor or constructor functions
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = r.nextID()
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].ID == id {
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product %d not found", id)
}

func (r *productRepo) List() []*Product {
	return r.productList
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].ID == p.ID {
			r.productList[i] = &p
			return r.productList[i], nil
		}
	}

	return nil, fmt.Errorf("product %d not found", p.ID)
}

func (r *productRepo) Delete(id int) (bool, error) {
	for i := range r.productList {
		if r.productList[i].ID == id {
			r.productList = append(r.productList[:i], r.productList[i+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("product %d not found", id)
}

func (r *productRepo) nextID() int {
	maxID := 0
	for _, p := range r.productList {
		if p.ID > maxID {
			maxID = p.ID
		}
	}

	return maxID + 1
}

func generateProducts(r *productRepo) {
	prd1 := Product{
		ID:          1,
		Title:       "MacBook Air",
		Description: "Apple MacBook Air laptop",
		ImageURL:    "https://example.com/macbook.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "iPhone 17",
		Description: "Apple iPhone",
		ImageURL:    "https://example.com/iphone.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "AirPods Pro",
		Description: "Apple wireless earbuds",
		ImageURL:    "https://example.com/airpods.jpg",
	}

	r.productList = append(r.productList, &prd1, &prd2, &prd3)
}

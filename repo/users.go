package repo

import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(id int) (*User, error)
	List() []*User
	Update(u User) (*User, error)
	Delete(id int) (bool, error)
}

type userRepo struct {
	userList []*User
}

//Constructor or constructor functions
func NewUserRepo() UserRepo {
	return &userRepo{userList: []*User{}}
}

func (r *userRepo) Create(u User) (*User, error) {
	u.ID = r.nextID()
	r.userList = append(r.userList, &u)
	return &u, nil
}

func (r *userRepo) Get(id int) (*User, error) {
	for i := range r.userList {
		if r.userList[i].ID == id {
			return r.userList[i], nil
		}
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func (r *userRepo) List() []*User {
	return r.userList
}

func (r *userRepo) Update(u User) (*User, error) {
	for i := range r.userList {
		if r.userList[i].ID == u.ID {
			r.userList[i] = &u
			return r.userList[i], nil
		}
	}

	return nil, fmt.Errorf("user %d not found", u.ID)
}

func (r *userRepo) Delete(id int) (bool, error) {
	for i := range r.userList {
		if r.userList[i].ID == id {
			r.userList = append(r.userList[:i], r.userList[i+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("user %d not found", id)
}

func (r *userRepo) nextID() int {
	maxID := 0
	for _, u := range r.userList {
		if u.ID > maxID {
			maxID = u.ID
		}
	}

	return maxID + 1
}

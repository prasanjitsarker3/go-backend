package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ErrUserNotFound is returned when no user matches the given id.
var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(id int) (*User, error)
	List() []*User
	Update(u User) (*User, error)
	Delete(id int) (bool, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}


func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{dbCon: dbCon}
}

func (r *userRepo) Create(u User) (*User, error) {

    var exists bool

    err := r.dbCon.Get(
        &exists,
        `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`,
        u.Email,
    )

    if err != nil {
        return nil, fmt.Errorf("check email: %w", err)
    }

    if exists {
        return nil, fmt.Errorf("email already exists")
    }

    query := `
        INSERT INTO users (name, email)
        VALUES ($1, $2)
        RETURNING id, name, email
    `

    var created User

    err = r.dbCon.Get(&created, query, u.Name, u.Email)

    if err != nil {
        return nil, fmt.Errorf("create user: %w", err)
    }

    return &created, nil
}

func (r *userRepo) Get(id int) (*User, error) {
	var u User
	err := r.dbCon.Get(&u, `SELECT id, name, email FROM users WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: id %d", ErrUserNotFound, id)
	}
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return &u, nil
}

func (r *userRepo) List() []*User {
	users := []*User{}
	if err := r.dbCon.Select(&users, `SELECT id, name, email FROM users ORDER BY id`); err != nil {
		return []*User{}
	}
	return users
}

func (r *userRepo) Update(u User) (*User, error) {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email`
	var updated User
	err := r.dbCon.Get(&updated, query, u.Name, u.Email, u.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: id %d", ErrUserNotFound, u.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("update user %d: %w", u.ID, err)
	}
	return &updated, nil
}

func (r *userRepo) Delete(id int) (bool, error) {
	res, err := r.dbCon.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("delete user %d: %w", id, err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if n == 0 {
		return false, fmt.Errorf("%w: id %d", ErrUserNotFound, id)
	}
	return true, nil
}

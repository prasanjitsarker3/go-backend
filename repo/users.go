package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"backend-golang/domain"
	"backend-golang/users"
)

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) users.UserRepo {
	return &userRepo{dbCon: dbCon}
}

func (r *userRepo) Create(u domain.User) (*domain.User, error) {

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

	var created domain.User

	err = r.dbCon.Get(&created, query, u.Name, u.Email)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &created, nil
}

func (r *userRepo) Get(id int) (*domain.User, error) {
	var u domain.User
	err := r.dbCon.Get(&u, `SELECT id, name, email FROM users WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: id %d", domain.ErrUserNotFound, id)
	}
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return &u, nil
}

func (r *userRepo) List() []*domain.User {
	users := []*domain.User{}
	if err := r.dbCon.Select(&users, `SELECT id, name, email FROM users ORDER BY id`); err != nil {
		return []*domain.User{}
	}
	return users
}

func (r *userRepo) Update(u domain.User) (*domain.User, error) {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email`
	var updated domain.User
	err := r.dbCon.Get(&updated, query, u.Name, u.Email, u.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: id %d", domain.ErrUserNotFound, u.ID)
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
		return false, fmt.Errorf("%w: id %d", domain.ErrUserNotFound, id)
	}
	return true, nil
}

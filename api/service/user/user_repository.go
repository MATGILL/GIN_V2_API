package user

import (
	"database/sql"
	"fmt"

	"github.com/MATGILL/GIN_V2/api/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByEmail(email string) (*types.User, error) {
	query := `SELECT id, firstName, lastName, email, password, created_at FROM users WHERE email = $1`

	rows, err := r.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := new(types.User)
	if rows.Next() {
		user, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user types.User) error {
	query := `INSERT INTO users (firstName, lastName, email, password) VALUES($1, $2, $3, $4)`
	_, err := r.db.Query(query, user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(id int) (*types.User, error) {
	query := `SELECT id, firstName, lastName, email, password FROM users WHERE id = $1`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("err1")
	}
	defer rows.Close()

	user := new(types.User)
	if rows.Next() {
		user, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("err2")
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// Utils
func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

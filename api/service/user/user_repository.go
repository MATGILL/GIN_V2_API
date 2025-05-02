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
	rows, err := r.db.Query("SELECT user FROM users WHERE user.email = ?", email)
	if err != nil {
		return nil, err
	}

	//create an empty pointer to user
	user := new(types.User)
	for rows.Next() {
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

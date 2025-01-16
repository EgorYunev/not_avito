package data

import (
	"database/sql"

	"github.com/EgorYunev/not_avito/internal/models"
)

type UserQuery struct {
	DB *sql.DB
}

func (u *UserQuery) Insert(model *models.User) error {
	stmt := `INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)`

	_, err := u.DB.Exec(stmt, model.Name, model.Email, model.Password)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserQuery) GetById(id int) (*models.User, error) {
	stmt := `SELECT * FROM users
			WHERE id = $1`

	row := u.DB.QueryRow(stmt, id)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	if row.Err() != nil {
		return nil, row.Err()
	}

	return user, nil

}

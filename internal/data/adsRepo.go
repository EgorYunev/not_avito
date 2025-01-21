package data

import (
	"database/sql"

	"github.com/EgorYunev/not_avito/internal/models"
)

type AdRepository struct {
	DB *sql.DB
}

func (r *AdRepository) Insert(ad *models.Ad) error {
	stmt := `INSERT INTO ads (title, description, price, user_id)
			VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(stmt, ad.Title, ad.Description, ad.Price, ad.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdRepository) Delete(adId int, email string) error {

	row := r.DB.QueryRow("SELECT id FROM users WHERE email = $1", email)

	var id int
	row.Scan(&id)

	if row.Err() != nil {
		return row.Err()
	}

	stmt := `DELETE FROM ads
			WHERE id = $1 AND user_id = $2`

	_, err := r.DB.Exec(stmt, adId, id)

	if err != nil {
		return err
	}

	return nil
}

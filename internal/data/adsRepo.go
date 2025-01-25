package data

import (
	"database/sql"

	"github.com/EgorYunev/not_avito/internal/models"
)

type AdRepository struct {
	DB *sql.DB
}

func (r *AdRepository) Insert(ad *models.Ad) error {
	stmt := `INSERT INTO ads (title, description, price, user_id, city)
			VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(stmt, ad.Title, ad.Description, ad.Price, ad.UserId, ad.City)
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

func (r *AdRepository) ChangeAd(ad *models.Ad, email string) error {
	row := r.DB.QueryRow("SELECT id FROM users WHERE email = $1", email)

	var id int
	row.Scan(&id)

	if row.Err() != nil {
		return row.Err()
	}

	ad.UserId = id

	stmt := `UPDATE ads SET title = $1, description = $2, price = $3
			WHERE id = $4`
	_, err := r.DB.Exec(stmt, ad.Title, ad.Description, ad.Price, ad.Id)

	if err != nil {
		return err
	}

	return err

}

func (r *AdRepository) GetAdsByCityAndTitle(req *models.Ad) ([]*models.Ad, error) {
	stmt := `SELECT * FROM ads
			WHERE title = $1 AND city = $2`

	rows, err := r.DB.Query(stmt, req.Title, req.City)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.Ad{}
	for rows.Next() {
		ad := &models.Ad{}

		rows.Scan(&ad.Id, &ad.Title, &ad.Description, &ad.Price, &ad.UserId, &ad.City)
		result = append(result, ad)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}

package persistence

import (
	"errors"
	"time"

	"github.com/challenge/pkg/models"
)

func CreateUser(u models.User) (int64, error) {
	db := GetConnection()

	q := `INSERT INTO USER (USERNAME, PASSWORD, LAST_UPDATED)
            VALUES(?, ?, ?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	r, err := stmt.Exec(u.Username, u.Password, time.Now())
	if err != nil {
		return 0, err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return 0, errors.New("ERROR: Se esperaba una fila afectada")
	}

	lastID, _ := r.LastInsertId()

	return lastID, nil
}

func GetByUsername(username string) (*models.User, error) {
	db := GetConnection()
	row := db.QueryRow("SELECT ID, USERNAME FROM USER WHERE USERNAME=?", username)

	us := new(models.User)
	if err := row.Scan(&us.ID, &us.Username); err != nil {
		if _, errGetAll := GetAll(); errGetAll != nil {
			return nil, errGetAll
		}
		return nil, err
	}

	return us, nil

}

func GetAll() ([]models.User, error) {
	db := GetConnection()
	q := `SELECT ID, USERNAME FROM USER`
	
	rows, err := db.Query(q)
	if err != nil {
		return []models.User{}, err
	}
	
	defer rows.Close()
	
	users := []models.User{}
	
	for rows.Next() {
		u := models.User{}
		rows.Scan(
			&u.ID,
			&u.Username,
		)

		users = append(users, u)
	}
	return users, nil
}

package persistence

import (
	"errors"
	"time"

	"github.com/nahuelmarianolosada/go-boilerplate/pkg/models"
)

func GetAllMessages(recipient, start, limit int) ([]models.Message, error) {
	db := GetConnection()
	q := `SELECT ID, SENDER, RECIPIENT, TYPE, TEXT, LAST_UPDATED FROM MESSAGE WHERE RECIPIENT = ? AND ID >= ? `

	rows, err := db.Query(q, recipient, start)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	messages := []models.Message{}

	for rows.Next() {
		u := models.Message{}
		rows.Scan(
			&u.ID,
			&u.Sender,
			&u.Recipient,
			&u.Content.Type,
			&u.Content.Text,
			&u.LastUpdated,
		)

		messages = append(messages, u)
	}
	return messages, nil
}

func CreateMessage(m models.Message) (*models.Message, error) {
	db := GetConnection()

	q := `INSERT INTO MESSAGE (SENDER, RECIPIENT, TYPE, TEXT, LAST_UPDATED)
            VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	contentText, contentType := m.Content.Text, m.Content.Type
	m.LastUpdated = time.Now()
	r, err := stmt.Exec(m.Sender, m.Recipient, &contentType, &contentText, &m.LastUpdated)
	if err != nil {
		return nil, err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return nil, errors.New("ERROR: Expected at least, one row affected")
	}

	lastID, _ := r.LastInsertId()
	m.ID = lastID
	return &m, nil
}

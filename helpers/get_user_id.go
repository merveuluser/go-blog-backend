package helpers

import "database/sql"

func GetUserID(db *sql.DB, username string) (int, error) {
	var userID int
	query := `SELECT id FROM authors WHERE username = $1;`

	err := db.QueryRow(query, username).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

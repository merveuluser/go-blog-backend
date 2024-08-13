package helpers

import (
	"database/sql"
)

func CheckUserByID(db *sql.DB, id int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM authors WHERE id = $1);`

	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckUserByName(db *sql.DB, username string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM authors WHERE username = $1);`

	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckEmail(db *sql.DB, email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM authors WHERE email = $1);`

	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckPostByID(db *sql.DB, id int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1);`

	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckCommentByID(db *sql.DB, id int) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM comments WHERE id = $1);`

	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckCategoryByName(db *sql.DB, name string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM categories WHERE category_name = $1);`

	err := db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckTable(db *sql.DB, table string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = $1);`

	row := db.QueryRow(query, table)

	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

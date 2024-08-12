package helpers

import (
	"database/sql"
	"fmt"
)

func CheckExistsByID(db *sql.DB, table string, id int) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = %d)", table, id)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckExistsByName(db *sql.DB, table string, name string) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE category_name = '%s')", table, name)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CheckTable(db *sql.DB, table string) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = '%s')", table)

	row := db.QueryRow(query)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CheckUser(db *sql.DB, id int) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM authors WHERE id = %d)", id)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

package models

import "database/sql"

type User struct {
	ID   int
	Name string
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) CreateTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`
	_, err := m.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) Insert(name string) error {
	stmt := `INSERT INTO users (name)
	SELECT ? 
	FROM dual
	WHERE NOT EXISTS (
		SELECT 1
		FROM users
		WHERE name = ?
	);`
	_, err := m.DB.Exec(stmt, name, name)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) Latest() ([]string, error) {
	stmt := `SELECT name FROM users LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var userNames []string

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Name)
		if err != nil {
			return nil, err
		}
		userNames = append(userNames, u.Name)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return userNames, nil
}

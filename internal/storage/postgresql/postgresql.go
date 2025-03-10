package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(dbPath string) (*Storage, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("db cannot open, %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlForSave string, name string) error {
	stmt, err := s.db.Prepare("INSERT INTO urlshortener(url, name) values($1, $2)")
	if err != nil {
		return fmt.Errorf("error:%s", err.Error())
	}
	_, err = stmt.Exec(urlForSave, name)
	if err != nil {
		return fmt.Errorf("error:%s", err.Error())
	}
	return nil
}
func (s *Storage) GetURL(name string) (string, error) {
	stmt, err := s.db.Prepare("SELECT name, url FROM urlshortener WHERE name = $1")
	if err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}
	var resUrl string
	err = stmt.QueryRow(name).Scan(&resUrl)
	if err != nil {
		return "", fmt.Errorf("error:%s", err.Error())
	}
	return resUrl, nil
}
func (s *Storage) DeleteURL(name string) error {
	stmt, err := s.db.Prepare("DELETE FROM urlshortener WHERE name = $1")
	if err != nil {
		return fmt.Errorf("error:%s", err.Error())
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("error:%s", err.Error())
	}
	return nil
}

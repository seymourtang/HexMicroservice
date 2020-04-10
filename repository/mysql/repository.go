package mysql

import (
	"HexMicroservice/shortener"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

type mysqlRepository struct {
	db     *sql.DB
	dbName string
}

func NewMysqlRepository(mysqlURL, dbName string) (shortener.RedirectService, error) {
	log.Println("Using MySQL...")
	repo := &mysqlRepository{}
	db, err := newMysqlDB(mysqlURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMysqlRepository")
	}
	repo.db = db
	repo.dbName = dbName
	return repo, nil
}
func newMysqlDB(mysqlURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (m *mysqlRepository) Find(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	stmt, err := m.db.Prepare(fmt.Sprintf("SELECT Id,Code,URL,CreatedAt FROM %s  WHERE CODE=?", m.dbName))
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	defer stmt.Close()
	var Id int
	var Code string
	var URL string
	var CreatedAt int64
	err = stmt.QueryRow(code).Scan(&Id, &Code, &URL, &CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.Wrap(err, fmt.Sprintf("No shot url with that code=%s", code))
	case err != nil:
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = Code
	redirect.URL = URL
	redirect.CreatedAt = CreatedAt
	return redirect, nil
}
func (m *mysqlRepository) Store(redirect *shortener.Redirect) error {
	stmt, err := m.db.Prepare(fmt.Sprintf(`INSERT %s (Code,URL,CreatedAt) VALUES (?,?,?)`, m.dbName))
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	defer stmt.Close()
	_, err = stmt.Exec(redirect.Code, redirect.URL, redirect.CreatedAt)
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}

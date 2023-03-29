package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"

	"github.com/dchest/uniuri"
)

const dbPath = "data.db"
const adminUser = "admin"

var database *sql.DB

func getDb() (*sql.DB, error) {
	if database == nil {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return db, err
		}
		database = db
	}
	return database, nil
}

func createTable() error {
	db, err := getDb()
	if err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS data (key TEXT NOT NULL PRIMARY KEY, val TEXT NOT NULL);"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (username TEXT NOT NULL PRIMARY KEY, password TEXT NOT NULL);"); err != nil {
		return err
	}

	return nil
}

func get(key string) (string, error) {
	db, err := getDb()
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT val FROM data WHERE key=?;", key)
	val := ""

	if err = row.Scan(&val); err == sql.ErrNoRows {
		return "", errors.New("not-found")
	}
	return val, err
}

func set(key string, val string) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO data VALUES(?, ?);", key, val)
	if err != nil {
		return err
	}

	return nil
}

func del(key string) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM data WHERE key=?;", key)
	if err != nil {
		return err
	}

	return nil
}

func checkUser(username string) bool {
	db, err := getDb()
	if err != nil {
		return false
	}

	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username=?;", username)
	val := 0

	if err = row.Scan(&val); err == sql.ErrNoRows {
		return false
	}

	return val > 0
}

func authUser(username string, password string) bool {
	db, err := getDb()
	if err != nil {
		return false
	}

	row := db.QueryRow("SELECT password FROM users WHERE username=?;", username)
	val := ""

	if err = row.Scan(&val); err == sql.ErrNoRows {
		return false
	}

	hashed := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashed[:]) == val
}

func addUser(username string, password string) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	hashed := sha256.Sum256([]byte(password))

	_, err = db.Exec("INSERT INTO users VALUES(?, ?);", username, hex.EncodeToString(hashed[:]))
	if err != nil {
		return err
	}

	return nil
}

func delUser(username string) error {
	if username == adminUser {
		return errors.New("forbidden")
	}

	db, err := getDb()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM users WHERE username=?;", username)
	if err != nil {
		return err
	}

	return nil
}

func initAdminUser() string {
	exists := checkUser(adminUser)

	if !exists {
		password := uniuri.New()
		addUser(adminUser, password)
		return password
	}

	return ""
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type postgresql struct {
	Ip           string `toml:"ip"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Database     string `toml:"database"`
	MaxConns     int    `toml:"maxconns"`
	MaxIdleConns int    `toml:"maxidleconns"`
}

func (db postgresql) Connect() *sql.DB {
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", db.Ip, db.Port, db.User, db.Password, db.Database)
	var err error
	conn, err = sql.Open("postgres", params)
	if err != nil {
		log.Fatal(err)
		log.Fatalln("Cannot connect to database.")
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
		log.Fatalln("Cannot connect to database")
	}
	conn.SetMaxOpenConns(db.MaxConns)
	conn.SetMaxIdleConns(db.MaxIdleConns)
	return conn
}

func (db postgresql) Disconnect() {
	conn.Close()
}

func (db postgresql) Query(attributes []string, tables []string,
	condition string, tx *sql.Tx) (*sql.Rows, error) {
	query := `select ` + strings.Join(attributes, ",") +
		` from ` + strings.Join(tables, ",")
	if condition != "" {
		query += ` where ` + condition
	}
	if tx != nil {
		return tx.Query(query)
	}
	return conn.Query(query)
}

func (db postgresql) Insert(table string, attributes []string,
	values []string, tx *sql.Tx) error {
	query := "insert into " + table
	if attributes != nil {
		query += "(" + strings.Join(attributes, ",") + ")"
	}
	query += " values(" + strings.Join(values, ",") + ")"
	return db.ExecSQL(query, tx)
}

func (db postgresql) Update(table string, attribute string, value string,
	condition string, tx *sql.Tx) error {
	query := "update " + table + " set " + attribute + "=" + value
	if condition != "" {
		query += " where " + condition
	}
	return db.ExecSQL(query, tx)
}

func (db postgresql) Delete(table string, condition string,
	cascadeOrStrict bool, tx *sql.Tx) error {
	query := "delete from " + table
	if condition != "" {
		query += " where " + condition
	}
	if cascadeOrStrict {
		query += " cascade"
	}
	return db.ExecSQL(query, tx)
}

func (db postgresql) QuerySQL(query string, tx *sql.Tx) (*sql.Rows, error) {
	if tx != nil {
		return tx.Query(query)
	}
	return conn.Query(query)
}

func (db postgresql) ExecSQL(query string, tx *sql.Tx) error {
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = conn.Prepare(query)
	}
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (db postgresql) BeginTransaction() (*sql.Tx, error) {
	return conn.Begin()
}

func (db postgresql) Commit(tx *sql.Tx) {
	tx.Commit()
}

func (db postgresql) Rollback(tx *sql.Tx) {
	tx.Rollback()
}

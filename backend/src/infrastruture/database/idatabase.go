package database

import (
	"database/sql"
)

type relationalDatabaseManipulation interface {
	//establish connection with database
	Connect() *sql.DB
	//close connection with database
	Disconnect()
	//formal query `select [attributes] from [tables] where [condition]`
	Query(attributes []string, tables []string, condition string, tx *sql.Tx) (*sql.Rows, error)
	//formal insertion `insert into [table]([attributes]) values([values])`
	Insert(table string, attributes []string, values []string, tx *sql.Tx) error
	//formal update `update [table] set [attribute]=[value] where [condition]`
	Update(table string, attribute string, value string, condition string, tx *sql.Tx) error
	//formal delete `delete from [table] where [condition] [cascade|strict]`
	Delete(table string, condition string, cascadeOrStrict bool, tx *sql.Tx) error
	//special query
	QuerySQL(query string, tx *sql.Tx) (*sql.Rows, error)
	//special SQL
	ExecSQL(query string, tx *sql.Tx) error
	//begin a transaction
	BeginTransaction() (*sql.Tx, error)
	//Commit a transaction
	Commit(*sql.Tx)
	//Rollback a transaction
	Rollback(*sql.Tx)
}

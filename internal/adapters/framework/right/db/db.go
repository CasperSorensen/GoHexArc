package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("db conn failed")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failed: %v", err)
	}
	return &Adapter{db: db}
}

func (databaseAdapter Adapter) CloseDbConnection() {
	err := databaseAdapter.db.Close()
	if err != nil {
		log.Fatal("Could not close connection: %v", err)
	}
}

func (databaseAdapter Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()

	if err != nil {
		return err
	}
	_, err = databaseAdapter.db.Exec(queryString, args...)

	if err != nil {
		return err
	}

	return nil
}

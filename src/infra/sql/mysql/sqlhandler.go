package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nemusou/note-api/src/config"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(dbconfig *config.DBConfig) *SqlHandler {
	dburl := dbconfig.User + ":" + dbconfig.Password + "@tcp([database]:3306)/" + dbconfig.Database
	conn, err := sql.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (sqlHandler *SqlHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sqlHandler.Conn.Query(query, args...)
}

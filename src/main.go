package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "popo:popo@tcp([database]:3306)/note")
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

func main() {
	sqlHandler := NewSqlHandler()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user_test", func(c echo.Context) error {
		rows, err := sqlHandler.Query("SELECT * FROM users")
		if err != nil {
			panic(err)
		}
		users := []User{}
		for rows.Next() {
			var user User
			err := rows.Scan(&user.UserId, &user.UserName)
			if err != nil {
				panic(err)
			}

			users = append(users, user)
		}
		return c.JSON(http.StatusOK, users)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

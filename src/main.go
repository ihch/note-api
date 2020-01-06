package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"github.com/nemusou/note-api/src/config"
)

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type SqlHandler struct {
	Conn *sql.DB
}

type UserRepository struct {
	sqlHandler *SqlHandler
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

func NewUserRepository(sqlHandler *SqlHandler) *UserRepository {
	userRepository := new(UserRepository)
	userRepository.sqlHandler = sqlHandler
	return userRepository
}

func (userRepository *UserRepository) Users() []User {
	rows, err := userRepository.sqlHandler.Query("SELECT * FROM users")
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
	return users
}

func (sqlHandler *SqlHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sqlHandler.Conn.Query(query, args...)
}

func main() {
	e := echo.New()
	setRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func setRoute(e *echo.Echo) {
	dbconfig := config.NewDBConfig()
	sqlHandler := NewSqlHandler(dbconfig)
	userRepository := NewUserRepository(sqlHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user_test", func(c echo.Context) error {
		users := userRepository.Users()
		return c.JSON(http.StatusOK, users)
	})
}

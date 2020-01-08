package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/nemusou/note-api/src/config"
	. "github.com/nemusou/note-api/src/infra/sql/mysql"
)

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type UserRepository struct {
	sqlHandler *SqlHandler
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

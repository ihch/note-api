package config

import "os"

type DBConfig struct {
	Database string
	User     string
	Password string
}

func NewDBConfig() *DBConfig {
	dbconfig := new(DBConfig)
	dbconfig.Database = os.Getenv("MYSQL_DATABASE")
	dbconfig.User = os.Getenv("MYSQL_USER")
	dbconfig.Password = os.Getenv("MYSQL_PASSWORD")
	return dbconfig
}

package db

import (
	"Estudos/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

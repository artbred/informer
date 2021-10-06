package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var database *sqlx.DB
var once sync.Once
var databaseUrl string

func Connection () *sqlx.DB {
	if database == nil {
		once.Do(Init)
	}

	return database
}

func Init () {
	var err error

	database, err = sqlx.Connect("pgx", databaseUrl); if err != nil {
		go panic(err)
	}

	if err = database.Ping(); err != nil {
		go panic(err)
	}

	database.SetMaxOpenConns(32)
	database.SetMaxIdleConns(32)

	logrus.Printf("Connected to %s", databaseUrl)
}

func init () {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	databaseUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	Init()
}

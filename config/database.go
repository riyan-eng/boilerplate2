package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var PostgreSQLDB *sql.DB

func ConnDatabase() {
	var err error
	dsn := fmt.Sprintf(
		`host='%v' user='%v' password='%v' dbname='%v' port='%v' sslmode=disable TimeZone='%v'`,
		GetEnv("DB_HOST"), GetEnv("DB_USERNAME"), GetEnv("DB_PASSWORD"), GetEnv("DB_NAME"), GetEnv("DB_PORT"), GetEnv("TIME_ZONE"))
	PostgreSQLDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("database: can't connect to database")
		os.Exit(1)
	}
	maxIdleConns, _ := strconv.Atoi(GetEnv("DB_MAX_IDDLE_CONN"))
	maxOpenConns, _ := strconv.Atoi(GetEnv("DB_MAX_OPEN_CONN"))
	connMaxLifetime, _ := strconv.Atoi(GetEnv("DB_CONN_MAX_LIFETIME_HOUR"))
	PostgreSQLDB.SetMaxIdleConns(maxIdleConns)
	PostgreSQLDB.SetMaxOpenConns(maxOpenConns)
	PostgreSQLDB.SetConnMaxLifetime(time.Hour * time.Duration(connMaxLifetime))
	if err := PostgreSQLDB.Ping(); err != nil {
		log.Fatal("database: can't ping to database")
	}

	fmt.Println("database: connection opened to database")
}

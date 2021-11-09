package configs

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type database struct {
	sqlDB *sql.DB
}

var db *database

func InitMySql() {
	db = new(database)
	connectionString := viper.GetString("MySqlDatabase.ConnectionString")
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(fmt.Errorf("error while connecting mysql database : %w", err))
	}
	err = sqlDb.Ping()
	if err != nil {
		panic(fmt.Errorf("unable to ping database : %w", err))
	}
	sqlDb.SetConnMaxIdleTime(3 * time.Minute)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)
	db.sqlDB = sqlDb
}

func GetMySqlDB() *sql.DB {
	return db.sqlDB
}

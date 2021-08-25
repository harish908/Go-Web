package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

type database struct {
	sqlDB *sql.DB
}

var db *database

func InitDB() {
	db = new(database)
	connectionString := viper.GetString("MySqlDatabase.ConnectionString")
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(fmt.Errorf("Error while connecting mysql database : %w\n", err))
	}
	sqlDb.SetConnMaxIdleTime(3 * time.Minute)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)
	db.sqlDB = sqlDb
}

func GetMySqlDB() *sql.DB {
	return db.sqlDB
}

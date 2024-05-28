package mysql

import (
	"database/sql"
	"fmt"

	"github.com/adeisbright/go-db/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConnection() (*gorm.DB, *sql.DB, error) {
	userName := config.GetConfig().SqlUserName
	password := config.GetConfig().SqlPassword
	database := config.GetConfig().SqlDatabase
	host := config.GetConfig().SqlHost
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Connection to Mysql DB was successful")
	return db, sqlDb, nil
}

func Init() {
	_, conn, err := MySQLConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

}

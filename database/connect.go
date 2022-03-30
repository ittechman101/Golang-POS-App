// database/connect.go
package database

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ittechman101/go-pos/config"
	"github.com/jinzhu/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	// configData := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	config.Config("DB_HOST"),
	// 	port,
	// 	config.Config("DB_USER"),
	// 	config.Config("DB_PASSWORD"),
	// 	config.Config("DB_NAME"),
	// )

	// DB, err = gorm.Open(
	// 	"postgres",
	// 	configData,
	// )

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"))

	DB, err = gorm.Open("mysql", dsn)
	DB.DB().SetConnMaxLifetime(100)
	DB.DB().SetMaxIdleConns(10)
	// set max connection
	//	DB.SetConnMaxLifetime(100)
	// set max idle connections
	//	DB.SetMaxIdleConns(10)

	// err = DB.DB().Ping()
	// if err != nil {
	// 	log.WithError(err).Fatal("error while pinging DB")
	// }

	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}

package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	viper.SetConfigFile("app.env")
	viper.ReadInConfig()
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")

	postgres_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)

	fmt.Println("Connecting to postgres connection\t\t", postgres_conname)
	//"host=localhost port=5432 user=postgres dbname=goTest password=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(postgres_conname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Book{})

	// m := Book{Author: "TestAuthor", Title: "TestTitle"}
	// db.Create(&m)
	return db
}

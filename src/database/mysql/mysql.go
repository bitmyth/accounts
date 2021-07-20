package mysql

import (
	"fmt"
	"github.com/bitmyth/accounts/src/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Dsn() string {

	username := config.Secret.GetString("database.username")
	password := config.Secret.GetString("database.password")
	host := config.Secret.GetString("database.host")
	port := config.Secret.GetString("database.port")
	database := config.Secret.GetString("database.schema")

	if h := viper.GetString("DATABASE_USERNAME"); h != "" {
		username = h
	}
	if h := viper.GetString("DATABASE_PASSWORD"); h != "" {
		password = h
	}
	if h := viper.GetString("DATABASE_HOST"); h != "" {
		host = h
	}
	if h := viper.GetString("DATABASE_PORT"); h != "" {
		port = h
	}
	if h := viper.GetString("DATABASE_DB"); h != "" {
		database = h
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	println(dsn)

	return dsn
}

func Connect() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	if err := Bootstrap(); err != nil {
		return nil, err
	}

	return DB, nil
}

func Bootstrap() error {

	if DB != nil {
		return nil
	}

	dsn := Dsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

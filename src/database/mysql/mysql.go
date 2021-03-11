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

	username := viper.GetString("database.username")
	password := config.Secret.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.schema")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

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
	println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

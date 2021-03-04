package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var S *gorm.DB

func Dsn() string {

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.schema")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	return dsn
}

func DB() (*gorm.DB, error) {
	if S != nil {
		return S, nil
	}

	if err := Bootstrap(); err != nil {
		return nil, err
	}

	return S, nil
}

func Bootstrap() error {

	if S != nil {
		return nil
	}

	dsn := Dsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	S = db

	return nil
}

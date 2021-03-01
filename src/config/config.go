package config

import (
    "github.com/spf13/viper"
    "os"
)

var (
    RootPath string
)

func init() {
    RootPath, _ = os.Getwd()
}

func Bootstrap() error {
    viper.SetConfigName("dev")
    viper.AddConfigPath(RootPath + "/config") // path to look for the config file in
    err := viper.ReadInConfig()               // Find and read the config file
    if err != nil { // Handle errors reading the config file
        return err
    }
    return nil
}

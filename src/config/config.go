package config

import (
    "fmt"
    "github.com/spf13/viper"
)

func Read() {

    viper.SetConfigName("dev")
    viper.AddConfigPath("/Users/gsh/go/src/github.com/bitmyth/accounts/src/config") // path to look for the config file in
    err := viper.ReadInConfig()          // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}

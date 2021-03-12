package config

import (
	"github.com/spf13/viper"
	"os"
)

var (
	RootPath string
	Secret   *viper.Viper
)

func init() {
	RootPath, _ = os.Getwd()
}

func Bootstrap() error {
	viper.AutomaticEnv()
	viper.SetConfigName("plain")
	viper.AddConfigPath(RootPath + "/config") // path to look for the config file in
	viper.AddConfigPath("/config")            // path to look for the config file in
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		return err
	}

	Secret = viper.New()

	Secret.SetConfigName("secret")
	Secret.AddConfigPath(RootPath + "/config")
	Secret.AddConfigPath("/config")
	err = Secret.ReadInConfig()
	if err != nil {
		return err
	}

	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//    fmt.Println("Config file changed:", e.Name)
	//})
	return nil
}

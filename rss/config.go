package rss

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

//
func newConfig() (*viper.Viper, error) {

	configName := fmt.Sprint("config/config.yml")

	c := viper.New()
	c.SetConfigFile(configName)
	c.SetConfigType("yaml")
	c.WatchConfig()

	c.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed:", e.Name)
	})

	err := c.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		log.Panicf("fatal error config file: %s", err)

	}
	return c, nil
}

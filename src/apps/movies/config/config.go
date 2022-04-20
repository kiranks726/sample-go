// Config and settings here
// TODO: Load yaml file
// TODO: Load secrets here?
// REF: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
package config

import "os"

type Config struct {
	Movies struct {
		TableName    string ""
		HelloSetting string ""
	}
}

// GetConfig will setup valies and return
// an instance of the configuration
func (c Config) GetConfig() Config {
	c.Movies.HelloSetting = "Hello Monkey"
	c.Movies.TableName = os.Getenv("movieTableName")

	return c
}

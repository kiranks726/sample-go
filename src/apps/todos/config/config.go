// Config and settings here
// TODO: Load yaml file
// TODO: Load from env variables
// REF: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
package config

import "os"

type Config struct {
	Todos struct {
		TableName    string ""
		HelloSetting string ""
	}
}

// GetConfig will setup valies and return
// an instance of the configuration
func (c Config) GetConfig() Config {
	c.Todos.HelloSetting = "Hello Monkey"
	c.Todos.TableName = string(os.Getenv("todoTableName"))
	return c
}

// TODO: Load secrets here?

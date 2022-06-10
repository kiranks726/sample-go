// Config and settings here
// TODO: Load yaml file
// TODO: Load secrets here?
// REF: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Version string
	Stacks  struct {
		Auth struct {
			Name             string
			Region           string
			AccountId        string
			UserPoolId       string
			UserPoolClientId string
		}
		Movies struct {
			Name        string
			Region      string
			AccountId   string
			ApiEndpoint string `json:"MoviesApiEndpoint"`
			Tablename   string `json:"MoviesTablename"`
		} `json:"movie-stack"`
	}
	Assets struct {
		Welcome struct {
			Key      string
			Src      string
			Location string
		}
	}
	Secrets struct {
		AcsGitUser struct {
			Name          string
			Description   string
			ParameterName string
		} `json:"acs-git-user"`
	}
}

// GetConfig will setup valies and return
// an instance of the configuration
func (c Config) GetConfig() Config {
	start := time.Now()
	// TODO: How to pass in config URL, it should not really be hardcoded.

	jsonStr := "{}"
	stage := os.Getenv("STAGE")
	if stage == "local" {
		jsonStr = c.LocalConfig()
	} else {
		appConfigUrl := "http://localhost:2772/applications/" +
			stage + "-ctx-kitchensink-go-application/environments/" +
			stage + "/configurations/stacks-config"
		jsonStr = c.AppConfig(appConfigUrl)
	}
	duration := time.Since(start)

	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		log.Error().Msgf("Decode JSON config failed: %v", err)
	}
	log.Info().Msgf("Get app config duration: %v", duration)

	return c
}

func (c Config) LocalConfig() string {
	const configPath = "config/local-local.json"
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error().Msgf("Unable to load config %v", err)
	}
	log.Info().Msgf("Loaded config from %s\n", configPath)
	return string(file)
}

func (c Config) AppConfig(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal().Msgf("Get config failed: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Msgf("Parse config failed: %v", err)
	}
	return string(body)
}

func (c Config) GetExports() {
	sess := session.Must(session.NewSession())

	svc := cloudformation.New(sess)
	list, err := svc.ListExports(&cloudformation.ListExportsInput{})
	if err != nil {
		log.Error().Err(err).Msg("List exports failed")
	}
	log.Info().Msgf("list.Exports: %v\n", list.Exports)
}

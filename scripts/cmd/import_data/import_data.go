package main

// @NAME: data_import
// @DESCRIPTION: Seed database with csv file
// @REF: https://gosamples.dev/read-csv
// @REF: https://blog.logrocket.com/making-http-requests-in-go
// @REF: https://gobyexample.com/command-line-flags
// @REF: https://levelup.gitconnected.com/easy-reading-and-writing-of-csv-files-in-go-7e5b15a73c79 - Map CSV Header
// REF: https://github.com/rs/zerolog

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"net/http"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/kiranks726/sample-go/src/apps/movies/services/movies"
)

// CONSTANTS
var RESOURCE_URL string = "https://j7grpr203l.execute-api.us-east-1.amazonaws.com/movies"
var DEFAULT_CSV_PATH string = "extra/data/movies.csv"

// FUNCTIONS
func storeItem(url string, item *movies.Movie) {
	body, _ := json.Marshal(&item)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("http ERROR: %v \n", err)
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("response ERROR: %v \n", err)
			// Failed to read response.
			panic(err)
		}

		// Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

	} else {
		// The status is not Created. print the error.
		fmt.Println("Get failed with error: ", resp.Status)
	}
}

func storeItems(url string, items []movies.Movie) {
	for i, item := range items {
		storeItem(url, &item)
		fmt.Printf("storedItem: %v \n", i)
		fmt.Println("-------------------------------------------------")
	}
}

// TODO: Make this use go routines
func storeItemsFast(url string, items []movies.Movie) {
	log.Info().Msg("Fast Option: using storeItemsFast()")
	for i, item := range items {
		storeItem(url, &item)
		fmt.Println("")
		fmt.Printf("storedItem: %v \n", i)
		fmt.Println("-------------------------------------------------")
	}
}

func loadCsvToItems(csvPathFlag *string, items []movies.Movie) []movies.Movie {
	f, err := os.Open(*csvPathFlag)
	if err != nil {
		log.Error().Err(err).Msg("FILE ERROR")
	}

	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &items); err != nil {
		panic(err)
	}

	return items
}

// MAIN
func main() {

	// Import commandline flags
	debug := flag.Bool("debug", false, "sets log level to debug")
	csvPathFlag := flag.String("p", DEFAULT_CSV_PATH, "Path to the csv file to be imported.")
	urlFlag := flag.String("u", RESOURCE_URL, "Target url for importing data.")
	fastFlag := flag.Bool("f", false, "Run import with parallel execution for faster results.")
	flag.Parse()

	log.Printf("Debug value: %v", *debug)
	log.Printf("Path value: %v", *urlFlag)
	log.Printf("Path value: %v", *csvPathFlag)
	log.Printf("Fast option: %t", *fastFlag)

	// Setup logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix // UNIX Time is faster and smaller than most timestamps
	// default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Main CSV code

	items := loadCsvToItems(csvPathFlag, []movies.Movie{})
	// Store Items
	if *fastFlag {
		storeItemsFast(*urlFlag, items)
	} else {
		storeItems(*urlFlag, items)
	}

}

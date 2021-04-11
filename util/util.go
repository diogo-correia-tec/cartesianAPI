package util

import (
	"cartesianAPI/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CoordinateList shares the global Coordinates imported
var (
	CoordinateList []domain.Coordinate
)

// LoadData loads data from the points file
func LoadData() {
	// Filename is the path to the json data file
	file, err := os.Open("data/points.json")
	if err != nil {
		FailOnError(err, "No file or invalid file in data folder")
	}

	byteValue, errReader := ioutil.ReadAll(file)
	if errReader != nil {
		FailOnError(errReader, "Unable to read from file")
	}

	errUnmarshal := json.Unmarshal([]byte(byteValue), &CoordinateList)
	if errUnmarshal != nil {
		FailOnError(err, "Error decoding data file ")
	}
}

//FailOnError logs fatal errors
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

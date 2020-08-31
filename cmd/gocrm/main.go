package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/demsasha4yt/gocrm.git/internal/app/gocrm"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/gocrm.json", "path to config file")
}

func getConfigData(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	flag.Parse()

	config := gocrm.NewConfig()
	configData, err := getConfigData(configPath)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
		return
	}

	if err := gocrm.Start(config); err != nil {
		log.Fatal(err)
	}
}

package fynelibs

import (
	"fmt"
	"os"
	
	"gopkg.in/yaml.v3"
)

type Config struct {
	Game struct {
		Width int `yaml:"width"`
		Height int `yaml:"height"`
		SavePath string `yaml:"save_path"`
		Fullscreen bool `yaml:"fullscreen"`
	} `yaml:"game"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(filename string, cfg *Config) {
	f, err := os.Open(filename)
	if err != nil {
		processError(err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		processError(err)
	}
}

func Setup() Config {
	var cfg Config
	readFile("./config.yaml", &cfg)

	return cfg
}
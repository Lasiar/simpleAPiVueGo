package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	_once   sync.Once
	_config *config
)

type config struct {
	Port string `json:"port"`
}

func (c *config) load() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		return err
	}

	return nil
}

// singleton
func GetConfig() *config {
	_once.Do(func() {
		_config = new(config)

		if err := _config.load(); err != nil {
			log.Fatalf("Error read config: %v", err)
		}
	})
	return _config
}

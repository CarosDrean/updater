package utils

import (
	"encoding/json"
	"github.com/CarosDrean/updater/models"
	"os"
)

func GetConfigs() (models.Configs, error) {
	var configs models.Configs
	file, err := os.Open("./configuration.json")

	if err != nil {
		return configs, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configs)

	if err != nil {
		return configs, err
	}

	return configs, nil
}

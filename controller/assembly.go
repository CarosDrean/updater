package controller

import (
	"fmt"
	"github.com/CarosDrean/updater/models"
	"strconv"
)

func assemblyIDs(configs []models.Config)[]string {
	options := make([]string, 0)
	for _, e := range configs {
		options = append(options, e.ID)
	}
	options = append(options, strconv.Itoa(len(configs) + 1))
	return options
}

func assemblyOptions(configs []models.Config)[]string {
	options := make([]string, 0)
	for _, e := range configs {
		options = append(options, fmt.Sprintf("%s.- Actualizar %s", e.ID, e.NameApp))
	}
	options = append(options, fmt.Sprintf("%d.- Salir", len(configs) + 1))
	return options
}

package controller

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"github.com/CarosDrean/updater/models"
	"github.com/CarosDrean/updater/utils"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	ErrMain        error
	configSelected models.Config
	red            = color.FgRed.Render
	green          = color.FgGreen.Render
)

func Options() (string, error) {
	configs, err := utils.GetConfigs()
	if err != nil {
		return "", err
	}
	option, err := prompts(configs.Configs)
	if err != nil {
		return "", err
	}

	if option == strconv.Itoa(len(configs.Configs) + 1) {
		os.Exit(3)
	}

	config := findConfig(option, configs.Configs)
	err = work(config)
	if err != nil {
		return "", err
	}
	return config.NameApp, nil
}

func findConfig(option string, configs []models.Config) models.Config {
	for _, e := range configs {
		if e.ID == option {
			configSelected = e
			return e
		}
	}
	return models.Config{}
}

func work(config models.Config) error {
	fmt.Println(green(fmt.Sprintf(constants.MessageInit, config.NameApp)))
	err := deleteDir()
	if err != nil {
		return err
	}
	updater(config.RouteFrom)

	if ErrMain != nil {
		return ErrMain
	}
	return nil
}

func updater(routeFrom string) {
	files, err := ioutil.ReadDir(routeFrom)
	checkErr(err, "Read Files")
	if err != nil {
		return
	}

	err = createFolder(configSelected.RouteTo)
	checkErr(err, "Creating Folder")

	if len(getSubRoute(configSelected.RouteFrom, routeFrom)) > 2 {
		err = createFolder(configSelected.RouteTo + getSubRoute(configSelected.RouteFrom, routeFrom))
		checkErr(err, "Creating Folder Sub")
	}

	for _, file := range files {
		if file.IsDir() {
			updater(routeFrom + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				configSelected.RouteTo+getSubRoute(configSelected.RouteFrom, routeFrom),
				configSelected.RouteFrom+getSubRoute(configSelected.RouteFrom, routeFrom),
			)
		}
	}
}

// obtiene el nombre de las subcarpetas
func getSubRoute(route string, subRoute string) string {
	sub := []rune(subRoute)
	return string(sub[len(route):])
}

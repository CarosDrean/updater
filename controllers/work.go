package controllers

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
	ErrMain error
	Routes  models.Config
	red     = color.FgRed.Render
	green   = color.FgGreen.Render
)

func Options() error {
	config, err := utils.GetConfigs()
	if err != nil {
		return err
	}
	option, err := prompts(config.Configs)
	if err != nil {
		return err
	}

	if option == strconv.Itoa(len(config.Configs) + 1) {
		os.Exit(3)
	}

	err = work(findConfig(option, config.Configs))
	if err != nil {
		return err
	}
	return nil
}

func findConfig(option string, configs []models.Config) models.Config {
	for _, e := range configs {
		if e.ID == option {
			Routes = e
			return e
		}
	}
	return models.Config{}
}

func work(routes models.Config) error {
	fmt.Println(green(fmt.Sprintf(constants.MessageInit, routes.NameApp)))
	err := deleteDir()
	checkErr(err, "Delete Dir")

	if err == nil {
		updater(routes.RouteFrom)
	}

	if ErrMain != nil {
		return ErrMain
	}
	return nil
}

func updater(routeFrom string) {
	files, err := ioutil.ReadDir(routeFrom)
	checkErr(err, "Red Files")
	if err != nil {
		return
	}

	err = createFolder(Routes.RouteTo)
	checkErr(err, "Creating Folder")

	if len(getSubRoute(Routes.RouteFrom, routeFrom)) > 2 {
		err = createFolder(Routes.RouteTo + getSubRoute(Routes.RouteFrom, routeFrom))
		checkErr(err, "Creating Folder Sub")
	}

	for _, file := range files {
		if file.IsDir() {
			updater(routeFrom + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				Routes.RouteTo+getSubRoute(Routes.RouteFrom, routeFrom),
				Routes.RouteFrom+getSubRoute(Routes.RouteFrom, routeFrom),
			)
		}
	}
}

// obtiene el nombre de las subcarpetas
func getSubRoute(route string, subRoute string) string {
	sub := []rune(subRoute)
	return string(sub[len(route):])
}

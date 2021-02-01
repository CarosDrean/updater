package controllers

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"github.com/CarosDrean/updater/models"
	"github.com/CarosDrean/updater/utils"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
)
var (
	ErrMain error
	Routes  models.Config
	red     = color.FgRed.Render
	green   = color.FgGreen.Render
)

func Options() error {
	option, err := prompts()
	if err != nil {
		return err
	}
	if option == "3" {
		os.Exit(3)
	}
	if option == "1" || option == "2" {
		err := work(getRoutesConfig(option))
		if err != nil {
			return err
		}
	} else {
		err = Options()
		if err != nil {
			return err
		}
	}
	return nil
}

func getRoutesConfig(option string) models.Config {
	config, err := utils.GetConfiguration()
	checkErr(err, "Get Configuration")

	Routes = models.Config{
		RouteFrom: config.RouteFrom,
		RouteTo:   config.RouteTo,
		NameApp:   config.NameApp,
	}
	if option == "2" {
		Routes = models.Config{
			RouteFrom: config.RouteFrom2,
			RouteTo:   config.RouteTo2,
			NameApp:   config.NameApp2,
		}
	}
	return Routes
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

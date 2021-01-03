package main

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"github.com/CarosDrean/updater/models"
	"github.com/CarosDrean/updater/utils"
	"github.com/gookit/color"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	errMain     error
	config      models.Configuration
	red         = color.FgRed.Render
	green       = color.FgGreen.Render
	blue        = color.FgLightBlue.Render
)

func main() {
	err := work()
	if err != nil {
		log.Println(red(constants.FinishError))
		log.Println(red(errMain))
	} else {
		fmt.Println(blue(constants.FinishSuccess))
	}

	fmt.Println()
	fmt.Printf("Presione %s para salir...", blue("ENTER"))
	_, _ = fmt.Scanln()
}

func work() error {
	fmt.Println(green(constants.MessageInit))
	err := deleteDir()
	checkErr(err, "Delete Dir")
	config, err = utils.GetConfiguration()
	if err == nil {
		updater(config.RouteFrom)
	}
	checkErr(err, "Get Configuration")

	if errMain != nil {
		return errMain
	}
	return nil
}

func deleteDir() error {
	fmt.Println(constants.DeleteDirOld)
	err := os.RemoveAll(config.RouteTo)
	return err
}

func createFolder(route string) error {
	err := os.MkdirAll(route, 0777)
	return err
}

func updater(routeFrom string) {
	files, err := ioutil.ReadDir(routeFrom)
	checkErr(err, "Red Files")
	if err != nil {
		return
	}

	err = createFolder(config.RouteTo)
	checkErr(err, "Creating Folder")

	if len(getSubRoute(config.RouteFrom, routeFrom)) > 2 {
		err = createFolder(config.RouteTo + getSubRoute(config.RouteFrom, routeFrom))
		checkErr(err, "Creating Folder Sub")
	}

	for _, file := range files {
		if file.IsDir() {
			updater(routeFrom + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				config.RouteTo+getSubRoute(config.RouteFrom, routeFrom),
				config.RouteFrom+getSubRoute(config.RouteFrom, routeFrom),
			)
		}
	}
}

// obtiene el nombre de las subcarpetas
func getSubRoute(route string, subRoute string) string {
	sub := []rune(subRoute)
	return string(sub[len(route):])
}

func copyFile(name string, routeTo string, routeFrom string) {
	fmt.Println("Writing: " + routeTo + "\\" + name)
	srcFile, err := os.Open(routeFrom + "\\" + name)
	checkErr(err, "Open File")
	defer srcFile.Close()

	destFile, err := os.Create(routeTo + "\\" + name) // creates if file doesn't exist
	checkErr(err, "Creating File")
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	checkErr(err, "Copying File")

	err = destFile.Sync()
	checkErr(err, "Sync File")
}

func checkErr(err error, ctx string) {
	if err != nil {
		errMain = err
		log.Println(red(ctx))
	}
}

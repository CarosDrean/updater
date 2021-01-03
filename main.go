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
	config      models.Configuration
	errVerifier = false
	red         = color.FgRed.Render
	green       = color.FgGreen.Render
)

func main() {
	err := work()
	if err != nil {
		log.Println(red(constants.FinishError))
	}
	fmt.Println(green(constants.FinishSuccess))
	fmt.Println()
	fmt.Printf("Presione %s para salir...", green("ENTER"))
	_, _ = fmt.Scanln()
}

func work() error {
	fmt.Println(green(constants.MessageInit))

	config, err := utils.GetConfiguration()
	checkErr(err)
	err = deleteDir()
	checkErr(err)
	updater(config.RouteFrom)

	if errVerifier {
		return err
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
	err := createFolder(config.RouteTo)
	checkErr(err)
	files, err := ioutil.ReadDir(routeFrom)
	checkErr(err)
	fmt.Println(getSubRoute(config.RouteFrom, routeFrom))

	if len(getSubRoute(config.RouteFrom, routeFrom)) > 2 { // si devuelve subcarpeta
		// esto se puede poner directamente en la primera linea de esta funcion
		err = createFolder(config.RouteTo +getSubRoute(config.RouteFrom, routeFrom))
		checkErr(err)
	}

	for _, file := range files {
		if file.IsDir() {
			updater(routeFrom + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				config.RouteFrom+getSubRoute(config.RouteFrom, routeFrom),
				config.RouteTo+getSubRoute(config.RouteFrom, routeFrom),
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
	fmt.Println("Writing: " + routeFrom + "\\" + name)
	srcFile, err := os.Open(routeTo + "\\" + name)
	checkErr(err)
	defer srcFile.Close()

	destFile, err := os.Create(routeFrom + "\\" + name) // creates if file doesn't exist
	checkErr(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	checkErr(err)

	err = destFile.Sync()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		errVerifier = true
		log.Println("Error.....................")
		log.Println(err)
	}
}

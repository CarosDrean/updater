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
	errMain error
	routes  models.Routes
	red     = color.FgRed.Render
	green   = color.FgGreen.Render
	blue    = color.FgLightBlue.Render
)

func main() {
	err := options()
	if err != nil {
		log.Println(red(constants.FinishError))
		log.Println(red(errMain))
	} else {
		fmt.Println(blue(fmt.Sprintf(constants.FinishSuccess, routes.NameApp)))
	}

	fmt.Println()
	fmt.Printf("Presione %s para salir...", blue("ENTER"))
	_, _ = fmt.Scanln()
}

func work(option string) error {
	config, err := utils.GetConfiguration()
	checkErr(err, "Get Configuration")

	routes = models.Routes{
		RouteFrom: config.RouteFrom,
		RouteTo:   config.RouteTo,
		NameApp:   config.NameApp,
	}
	if option == "2" {
		routes = models.Routes{
			RouteFrom: config.RouteFrom2,
			RouteTo:   config.RouteTo2,
			NameApp:   config.NameApp2,
		}
	}
	fmt.Println(green(fmt.Sprintf(constants.MessageInit, routes.NameApp)))
	err = deleteDir()
	checkErr(err, "Delete Dir")

	if err == nil {
		updater(routes.RouteFrom)
	}

	if errMain != nil {
		return errMain
	}
	return nil
}

func options() error {
	fmt.Println(fmt.Sprintf("===============%s===============", blue("OPCIONES")))
	fmt.Println("1.- Actualizar Sigesoft")
	fmt.Println("2.- Actualizar Sigesoft Particular")
	fmt.Println("3.- Salir")
	fmt.Println()
	fmt.Println(fmt.Sprintf("* Ingrese la %s deseada y presione %s", blue("OPCION"), blue("ENTER")))
	var option string
	_, _ = fmt.Scanln(&option)
	if option == "3" || option == "exit" {
		os.Exit(3)
	}
	if option == "1" || option == "2" {
		err := work(option)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("------------------------------")
		fmt.Println(red("¡Opcion no valida!"))
		err := options()
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteDir() error {
	fmt.Println(constants.DeleteDirOld)
	err := os.RemoveAll(routes.RouteTo)
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

	err = createFolder(routes.RouteTo)
	checkErr(err, "Creating Folder")

	if len(getSubRoute(routes.RouteFrom, routeFrom)) > 2 {
		err = createFolder(routes.RouteTo + getSubRoute(routes.RouteFrom, routeFrom))
		checkErr(err, "Creating Folder Sub")
	}

	for _, file := range files {
		if file.IsDir() {
			updater(routeFrom + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				routes.RouteTo+getSubRoute(routes.RouteFrom, routeFrom),
				routes.RouteFrom+getSubRoute(routes.RouteFrom, routeFrom),
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
	fmt.Println(fmt.Sprintf("%s: %s\\%s", green("Writing"), routeTo, name))
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

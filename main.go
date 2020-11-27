package main

import (
	"fmt"
	"github.com/CarosDrean/updater/models"
	"github.com/CarosDrean/updater/utils"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var config models.Configuration
var errVerifier = false

func main(){
	fmt.Println("Actualizando Sigesoft...")
	config = getConfig()
	err := deleteDir()
	if err != nil {
		checkErr(err)
	} else {
		updater(config.RouteTo)
	}
	if errVerifier {
		log.Println("¡Hubo un error!")
	} else {
		fmt.Println("¡Actualizado con exito!")
	}
	fmt.Println()
	fmt.Println("Presione ENTER para salir...")
	_, _ = fmt.Scanln()
}

func getConfig() models.Configuration{
	config, err := utils.GetConfiguration()
	checkErr(err)
	return config
}

func deleteDir() error {
	fmt.Println("Eliminando archivos antiguos...")
	err := os.RemoveAll(config.RouteFrom)
	return err
}

func createFolder() {
	err := os.MkdirAll(config.RouteFrom, 0777)
	if err != nil {
		log.Println(err)
	}
}

func updater(route string) {
	createFolder()
	files, err := ioutil.ReadDir(route)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files{
		if len(getSubRoute(config.RouteTo, route)) > 2 {
			err = os.MkdirAll(config.RouteFrom+getSubRoute(config.RouteTo, route), 0777)
			checkErr(err)
		}

		// fmt.Println("Nombre:", file.Name())
		// fmt.Println("Tamaño:", file.Size())
		// fmt.Println("Modo:", file.Mode())
		// fmt.Println("Ultima modificación:", file.ModTime())
		//  fmt.Println("Es directorio?:", file.IsDir())
		if file.IsDir() {
			updater(route + "\\" + file.Name())
		} else {
			copyFile(
				file.Name(),
				config.RouteTo + getSubRoute(config.RouteTo, route),
				config.RouteFrom + getSubRoute(config.RouteTo, route),
			)
		}
	}
}

func getSubRoute(route string, subRoute string) string {
	runer := []rune(subRoute)
	return string(runer[len(route):])
}

func copyFile(name string, routeTo string, routeFrom string) {
	fmt.Println("Copiando: " + routeTo + "\\" + name + " en " + routeFrom + "\\" + name)
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

func checkErr(err error){
	if err != nil {
		errVerifier = true
		log.Println("Error.....................")
		log.Println(err)
	}
}

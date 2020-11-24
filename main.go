package main

import (
	"github.com/CarosDrean/updater/constants"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main(){
	log.Println("Actualizando Sigesoft...")
	err := deleteDir()
	if err != nil {
		log.Println(err)
	} else {
		updater(constants.RouteFolderUpdate)
	}
	log.Println("¡Actualizado con exito!")
}

func deleteDir() error {
	log.Println("Eliminando archivos antiguos...")
	err := os.RemoveAll(constants.RouteFolderDelete)
	return err
}

func createFolder() {
	err := os.MkdirAll(constants.RouteFolderDelete, 0777)
	if err != nil {
		log.Println(err)
	}
}

func updater(route string) {
	log.Println("Copiando los nuevos archivos...")
	createFolder()
	files, err := ioutil.ReadDir(route)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files{
		if len(getSubRoute(constants.RouteFolderUpdate, route)) > 2 {
			err = os.MkdirAll(constants.RouteFolderDelete+getSubRoute(constants.RouteFolderUpdate, route), 0777)
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
				constants.RouteFolderUpdate + getSubRoute(constants.RouteFolderUpdate, route),
				constants.RouteFolderDelete + getSubRoute(constants.RouteFolderUpdate, route),
			)
		}
	}
}

func getSubRoute(route string, subRoute string) string {
	return "\\" + strings.TrimLeft(subRoute, route)
}

func copyFile(name string, routeTo string, routeFrom string) {
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
		log.Println(err)
	}
}

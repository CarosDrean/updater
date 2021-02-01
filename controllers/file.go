package controllers

import (
	"fmt"
	"io"
	"os"
)

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

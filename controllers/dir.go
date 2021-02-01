package controllers

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"os"
)

func deleteDir() error {
	fmt.Println(constants.DeleteDirOld)
	err := os.RemoveAll(Routes.RouteTo)
	return err
}

func createFolder(route string) error {
	err := os.MkdirAll(route, 0777)
	return err
}

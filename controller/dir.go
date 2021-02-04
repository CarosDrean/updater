package controller

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"os"
)

func deleteDir() error {
	fmt.Println(constants.DeleteDirOld)
	err := os.RemoveAll(configSelected.RouteTo)
	return err
}

func createFolder(route string) error {
	err := os.MkdirAll(route, 0777)
	return err
}

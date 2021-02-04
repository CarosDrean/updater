package main

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"github.com/CarosDrean/updater/controller"
	"github.com/gookit/color"
	"log"
)

var (
	red     = color.FgRed.Render
	blue    = color.FgLightBlue.Render
)

func main() {
	execute()
}

func execute() {
	nameApp, err := controller.Options()
	if err != nil {
		log.Println(red(constants.FinishError))
		log.Println(red(controller.ErrMain))
		log.Println(red(err))
	} else {
		fmt.Println(blue(fmt.Sprintf(constants.FinishSuccess, nameApp)))
	}
	execute()
}


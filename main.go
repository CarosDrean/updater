package main

import (
	"fmt"
	"github.com/CarosDrean/updater/constants"
	"github.com/CarosDrean/updater/controllers"
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
	err := controllers.Options()
	if err != nil {
		log.Println(red(constants.FinishError))
		log.Println(red(controllers.ErrMain))
		log.Println(red(err))
	} else {
		fmt.Println(blue(fmt.Sprintf(constants.FinishSuccess, controllers.Routes.NameApp)))
	}
	execute()
}


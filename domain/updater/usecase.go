package updater

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gookit/color"

	"github.com/CarosDrean/updater/model"
)

var green = color.FgGreen.Render

type Updater struct {
	// the routes in app is root
	app  model.App
	from string
}

func New() *Updater {
	return &Updater{}
}

func (u *Updater) SetApp(app model.App) {
	u.app = app
}

func (u *Updater) SetFrom(from string) {
	u.from = from
}

func (u *Updater) DeleteFolder() error {
	if err := os.RemoveAll(u.app.RouteTo); err != nil {
		return err
	}

	return nil
}

func (u *Updater) Updater() error {
	nameFolder := u.app.RouteTo

	subRoute, ok := getSubRoute(u.app.RouteFrom, u.from)
	if ok {
		nameFolder = nameFolder + subRoute
	}

	if err := os.MkdirAll(nameFolder, os.ModePerm); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(u.from)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			// update route from for the next iteration
			u.from = fmt.Sprintf("%s\\%s", u.from, file.Name())

			if err := u.Updater(); err != nil {
				return err
			}

			continue
		}

		if err := copyFile(file.Name(), u.app.RouteTo+subRoute, u.app.RouteFrom+subRoute); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(name string, routeTo string, routeFrom string) error {
	fmt.Println(fmt.Sprintf("%s: %s\\%s", green("Writing"), routeTo, name))

	srcFile, err := os.Open(routeFrom + "\\" + name)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(routeTo + "\\" + name) // creates if file doesn't exist
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return err
	}

	if err = destFile.Sync(); err != nil {
		return err
	}

	return nil
}

func getSubRoute(routeRoot string, newRoute string) (string, bool) {
	route := []rune(newRoute)
	subRoute := string(route[len(routeRoot):])

	if len(subRoute) > 2 {
		return subRoute, true
	}

	return "", false
}

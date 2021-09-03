package updater

import "github.com/CarosDrean/updater/model"

type UseCase interface {
	DeleteFolder() error
	Updater() error
	SetApp(app model.App)
	SetFrom(from string)
}

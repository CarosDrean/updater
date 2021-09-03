package model

type App struct {
	ID        int    `json:"id"`
	RouteFrom string `json:"route_from"`
	RouteTo   string `json:"route_to"`
	NameApp   string `json:"name_app"`
}

type Apps []App

func (as Apps) GetIDs() (IDs []int) {
	for _, app := range as {
		IDs = append(IDs, app.ID)
	}

	return
}

func (as Apps) GetByID(ID int) (App, bool) {
	for _, app := range as {
		if app.ID == ID {
			return app, true
		}
	}

	return App{}, false
}

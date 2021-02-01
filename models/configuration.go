package models

type Configuration struct {
	RouteFrom string `json:"routeFrom"`
	RouteTo   string `json:"routeTo"`
	NameApp   string `json:"nameApp"`

	RouteFrom2 string `json:"routeFrom2"`
	RouteTo2   string `json:"routeTo2"`
	NameApp2   string `json:"nameApp2"`
}

type Configs struct {
	Configs []Config `json:"configs"`
}

type Config struct {
	ID        string `json:"_id"`
	RouteFrom string `json:"routeFrom"`
	RouteTo   string `json:"routeTo"`
	NameApp   string `json:"nameApp"`
}

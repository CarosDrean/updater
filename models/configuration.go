package models

type Configuration struct {
	RouteFrom string `json:"routeFrom"`
	RouteTo   string `json:"routeTo"`

	RouteFrom2 string `json:"routeFrom2"`
	RouteTo2   string `json:"routeTo2"`
}

type Routes struct {
	RouteFrom string `json:"routeFrom"`
	RouteTo   string `json:"routeTo"`
}

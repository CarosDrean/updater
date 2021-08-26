package model

type Configuration struct {
	Messages Messages `json:"messages"`
	Apps     Apps     `json:"apps"`
}

type Messages struct {
	Init          string `json:"init"`
	DeleteOld     string `json:"delete_old"`
	FinishSuccess string `json:"finish_success"`
	FinishError   string `json:"finish_error"`
	InvalidOption string `json:"invalid_option"`
}

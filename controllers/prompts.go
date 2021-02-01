package controllers

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

var simpleQs = []*survey.Question{
	{
		Name: "Option",
		Prompt: &survey.Select{
			Message: "SELECIONE UNA OPCION:\n",
			Options: []string{
				"1.- Actualizar Sigesost",
				"2.- Actualizar Sigesoft Particular",
				"3.- Salir",
			},
		},
		Validate: survey.Required,
	},
}

func prompts() (string, error) {
	answers := struct {
		Option string
	}{}

	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return getOption(answers.Option), nil
}

func getOption(text string) string {
	options := []string{"1", "2", "3"}
	for _, e := range options {
		i := strings.Index(text, e)
		if i != -1 {
			return e
		}
	}
	return "0"
}

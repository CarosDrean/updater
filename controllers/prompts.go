package controllers

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/CarosDrean/updater/models"
	"strconv"
	"strings"
)

type Option struct {
	ID   string
	Text string
}

func prompts(configs []models.Config) (string, error) {
	simpleQs := []*survey.Question{
		{
			Name: "Option",
			Prompt: &survey.Select{
				Message: "SELECIONE UNA OPCION:\n",
				Options: assemblyOptions(configs),
			},
			Validate: survey.Required,
		},
	}
	answers := struct {
		Option string
	}{}

	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return getOption(answers.Option, configs), nil
}

func assemblyIDs(configs []models.Config)[]string {
	options := make([]string, 0)
	for _, e := range configs {
		options = append(options, e.ID)
	}
	options = append(options, strconv.Itoa(len(configs) + 1))
	return options
}

func assemblyOptions(configs []models.Config)[]string {
	options := make([]string, 0)
	for _, e := range configs {
		options = append(options, fmt.Sprintf("%s.- Actualizar %s", e.ID, e.NameApp))
	}
	options = append(options, fmt.Sprintf("%d.- Salir", len(configs) + 1))
	return options
}

func getOption(text string, configs []models.Config) string {
	options := assemblyIDs(configs)
	for _, e := range options {
		i := strings.Index(text, e)
		if i != -1 {
			return e
		}
	}
	return strconv.Itoa(len(configs))
}

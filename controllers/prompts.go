package controllers

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/CarosDrean/updater/utils"
	"log"
	"strconv"
	"strings"
)

type Option struct {
	ID   string
	Text string
}

var simpleQs = []*survey.Question{
	{
		Name: "Option",
		Prompt: &survey.Select{
			Message: "SELECIONE UNA OPCION:\n",
			Options: assemblyOptions(),
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

func assemblyIDs()[]string {
	config, err := utils.GetConfigs()
	if err != nil {
		log.Println(err)
	}
	options := make([]string, 0)
	for _, e := range config.Configs {
		options = append(options, e.ID)
	}
	options = append(options, strconv.Itoa(len(config.Configs) + 1))
	return options
}

func assemblyOptions()[]string {
	config, err := utils.GetConfigs()
	if err != nil {
		log.Println(err)
	}
	options := make([]string, 0)
	for _, e := range config.Configs {
		options = append(options, fmt.Sprintf("%s.- Actualizar %s", e.ID, e.NameApp))
	}
	options = append(options, fmt.Sprintf("%d.- Salir", len(config.Configs) + 1))
	return options
}

func getOption(text string) string {
	options := assemblyIDs()
	for _, e := range options {
		i := strings.Index(text, e)
		if i != -1 {
			return e
		}
	}
	return "0"
}

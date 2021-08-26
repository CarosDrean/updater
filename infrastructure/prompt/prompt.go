package prompt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"

	updaterUseCase "github.com/CarosDrean/updater/domain/updater"
	"github.com/CarosDrean/updater/model"
)

var (
	red  = color.FgRed.Render
	blue = color.FgLightBlue.Render
)

type Prompt struct {
	apps     model.Apps
	messages model.Messages
	updater  updaterUseCase.UseCase
}

func New(config model.Configuration) Prompt {
	updater := updaterUseCase.New()

	return Prompt{
		apps:     config.Apps,
		messages: config.Messages,
		updater:  updater,
	}
}

func (p Prompt) Execute() error {
	prompt := p.makePromptAndWaitResponse()

	answer, err := p.getAnswer(prompt)
	if err != nil {
		return err
	}

	idApp := p.searchIDInOption(answer.Option)

	if idApp == len(p.apps)+1 {
		os.Exit(3)
	}

	app, ok := p.apps.GetByID(idApp)
	if !ok {
		log.Println(red(p.messages.InvalidOption))

		if err := p.Execute(); err != nil {
			return err
		}
	}

	p.updater.SetApp(app)
	p.updater.SetFrom(app.RouteFrom)

	if err := p.updater.DeleteFolder(); err != nil {
		return err
	}

	if err := p.updater.Updater(); err != nil {
		return err
	}

	fmt.Println(blue(p.messages.FinishSuccess))

	if err := p.Execute(); err != nil {
		return err
	}

	return nil
}

func (p Prompt) searchIDInOption(option string) int {
	for _, id := range p.apps.GetIDs() {
		IDString := strconv.Itoa(id)

		i := strings.Index(option, IDString)
		if i != -1 {
			return id
		}
	}

	return len(p.apps) + 1
}

// makePromptAndWaitResponse make and wait until you get an answer
func (p Prompt) makePromptAndWaitResponse() []*survey.Question {
	prompt := []*survey.Question{
		{
			Name: "Option",
			Prompt: &survey.Select{
				Message: "SELECIONE UNA OPCION:\n",
				Options: p.makeOptions(),
			},
			Validate: survey.Required,
		},
	}

	return prompt
}

func (p Prompt) getAnswer(prompt []*survey.Question) (model.Answer, error) {
	answer := model.Answer{}

	err := survey.Ask(prompt, &answer)
	if err != nil {
		return model.Answer{}, err
	}

	return answer, nil
}

func (p Prompt) makeOptions() []string {
	options := make([]string, 0)

	for _, app := range p.apps {
		option := fmt.Sprintf("%d.- Actualizar %s", app.ID, app.NameApp)

		options = append(options, option)
	}

	options = append(options, fmt.Sprintf("%d.- Salir", len(p.apps)+1))

	return options
}

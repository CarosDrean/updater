package bootstrap

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"

	promptInfrastructure "github.com/CarosDrean/updater/infrastructure/prompt"
)

var (
	red = color.FgRed.Render
)

func Run() error {
	config := newConfiguration("./configuration.json")

	prompt := promptInfrastructure.New(config)

	if err := prompt.Execute(); err != nil {
		log.Println(red(config.Messages.FinishError))
		log.Println(red(err))

		fmt.Print("Press 'Enter' to continue...")
		_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

		return err
	}

	return nil
}

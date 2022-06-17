package main

import (
	"comifer/util"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var featureMap = map[string]string{
	"🤖 Android": ":robot:",
	"🍎 iOS":     ":apple:",
	"👪 共通範囲":    ":family:",
}

var selectMap = map[string]string{
	"🚀 improve performance": ":improve:",
	"🐛 fix bug":             ":bug:",
}

var selectOptionsOfEmoji = util.GetKeysOfMap(selectMap)

var qs = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "which kind of commit?",
			Options: selectOptionsOfEmoji,
			Default: "🚀 improve performance",
		},
	},
	{
		Name:     "message",
		Prompt:   &survey.Input{Message: "write commit message"},
		Validate: survey.Required,
	},
}

func main() {
	app := &cli.App{
		Name:    "comifer",
		Version: "0.0.1",
		Usage:   "make emoji prefixed git commit log",
		Action: func(c *cli.Context) error {
			answers := struct {
				Type    string
				Message string
			}{}

			err := survey.Ask(qs, &answers)
			if err != nil {
				log.Fatal(err)
			}
			commitMessage := fmt.Sprintf("%s %s\n", selectMap[answers.Type], answers.Message)
			f, err := os.Create("./.commitlog-tmp")
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write([]byte(commitMessage))
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

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

var prepareCommitShell = `#! /bin/bash
exec < /dev/tty
./comifer
exec < /dev/null

commit_log=$(cat ./.commitlog-tmp)
rm ./.commitlog-tmp
sed -i '.bak' "1s/^/${commit_log}/" $1
echo $commit_log%
`

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
			if c.NArg() == 1 && c.Args().Get(0) == "init" {
				f, err := os.Create(".git/hooks/prepare-commit-msg")
				if err != nil {
					log.Fatal(err)
				}
				_, err = f.Write([]byte(prepareCommitShell))
				fmt.Println("correctly initialized")
			} else if c.NArg() == 0 {
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
			} else {
				cli.ShowAppHelp(c)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

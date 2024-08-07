package main

import (
	"comifer/question"
	"comifer/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

var execCommand = exec.Command

var selectMap = map[string]string{
	"\U0001f680 improve performance": ":improve:",
	"\U0001f41b fix bug":             ":bug:",
}

var selectOptionsOfEmoji = util.GetKeysOfMap(selectMap)

var prepareWindowsCommitShell = `#! /bin/bash
exec < /dev/tty
comifer
exec < /dev/null

commit_log=$(cat ./.commitlog-tmp)
rm ./.commitlog-tmp
sed -i "1s/^/${commit_log}/" $1
echo $commit_log%
`

var prepareMacCommitShell = `#! /bin/bash
exec < /dev/tty
comifer
exec < /dev/null

commit_log=$(cat ./.commitlog-tmp)
rm ./.commitlog-tmp
sed -i '.bak' "1s/^/${commit_log}/" $1
echo $commit_log%
`

var prepareLinuxCommitShell = `#! /bin/bash
exec < /dev/tty
comifer
exec < /dev/null

commit_log=$(cat ./.commitlog-tmp)
rm ./.commitlog-tmp
sed -i "1s/^/${commit_log}/" $1
echo $commit_log%
`

var prepareShells = map[string]string{
	"linux":   prepareLinuxCommitShell,
	"darwin":  prepareMacCommitShell,
	"windows": prepareWindowsCommitShell,
}

func main() {
	app := &cli.App{
		Name:    "comifer",
		Version: "0.1.0",
		Usage:   "make emoji prefixed git commit log",
		Action: func(c *cli.Context) error {
			if c.NArg() == 1 && c.Args().Get(0) == "init" {
				f, err := os.Create(".git/hooks/prepare-commit-msg")
				if err != nil {
					log.Fatal(err)
				}
				_, err = f.Write([]byte(prepareShells[runtime.GOOS]))
				if err != nil {
					log.Fatal(err)
					return nil
				}
				_, err = execCommand("chmod", "a+x", ".git/hooks/prepare-commit-msg").Output()
				if err != nil {
					log.Fatal(err)
					return nil
				}
				fmt.Println("correctly initialized")
			} else if c.NArg() == 0 {
				config := question.GenerateQuestionConfig()
				commitMessage := question.GenerateCommitLog(config)
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

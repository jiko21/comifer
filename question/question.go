package question

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/text/unicode/norm"
)

type Config struct {
	Format string
	Steps  []Step
}

type Step struct {
	Type    string
	Message string
	Options []Option
}

// +gen slice:"Select[string]"
type Option struct {
	Value       string
	Description string
}

func getDescriptionsOfOptions(options []Option) []string {
	arr := make([]string, len(options))
	for i, item := range options {
		arr[i] = norm.NFD.String(item.Description)
	}
	return arr
}

func getValuesOfOptions(options []Option) []string {
	arr := make([]string, len(options))
	for i, item := range options {
		arr[i] = item.Value
	}
	return arr
}

var defaultConfig = Config{
	Format: "$1",
	Steps: []Step{
		{
			Type:    "text",
			Message: "enter commit log",
			Options: []Option{},
		},
	},
}

func GenerateQuestionConfig() Config {
	bytes, err := ioutil.ReadFile("./comifer.json")
	if err != nil {
		return defaultConfig
	}
	config := Config{}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return defaultConfig
	}
	return config
}

func GetValueFromSelect(options []Option, selectedValue string) string {
	for _, option := range options {
		if option.Description == selectedValue {
			return option.Value
		}
	}
	return ""
}

func GenerateCommitLog(config Config) string {
	steps := config.Steps
	commitMessage := config.Format
	for i, step := range steps {
		target := fmt.Sprintf("$%d", i+1)
		if step.Type == "select" {
			selectedItem := ""
			prompt := &survey.Select{
				Message: step.Message,
				Options: getDescriptionsOfOptions(step.Options),
			}
			survey.AskOne(prompt, &selectedItem)
			commitMessage = strings.Replace(commitMessage, target, GetValueFromSelect(step.Options, selectedItem), -1)
		} else if step.Type == "text" {
			message := ""
			prompt := &survey.Input{
				Message: step.Message,
			}
			survey.AskOne(prompt, &message)
			commitMessage = strings.Replace(commitMessage, target, message, -1)
		} else {
			panic("Config mismatched")
		}
	}
	return commitMessage
}

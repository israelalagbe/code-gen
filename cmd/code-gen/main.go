package main

import (
	"fmt"

	"github.com/burl/inquire"
	"github.com/israelalagbe/code-gen/internals/libs"
	"github.com/israelalagbe/code-gen/internals/models"
	"github.com/israelalagbe/code-gen/internals/utils"
	"github.com/israelalagbe/code-gen/validations"
)

func main() {
	var questions *inquire.Questions = nil
	questionModel := models.NewQuestionModel()

	questions = inquire.Query()
	questions.Input(&questionModel.TableName, "Table name", validations.Required)
	questions.Input(&questionModel.ModelName, "Model name", validations.Required)
	questions.Input(&questionModel.Path, "Base directory", validations.Required)

	fields := ""
	questions.Input(&fields, "Fields", validations.Required)

	questions.Exec()

	questionModel.Properties = utils.ParseFields(fields)

	for _, item := range questionModel.Items {
		questions = inquire.Query()
		questions.YesNo(&item.Included, "Include "+item.Name)
		questions.Exec()

		if item.Included {
			questions = inquire.Query()
			questions.Input(&item.Path, "Path to "+item.Name+" file", validations.Required)
			questions.Exec()
		}
	}

	for _, item := range questionModel.Items {
		if !item.Included {
			continue
		}

		itemPath := questionModel.Path + item.Path

		if libs.FileExists(itemPath) {
			questions = inquire.Query()
			questions.YesNo(&item.Included, "File "+itemPath+" already exists. Overwrite?")
			questions.Exec()
		}
	}

	fmt.Println(questionModel)
}

package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/burl/inquire"
	"github.com/israelalagbe/code-gen/internals/libs"
	"github.com/israelalagbe/code-gen/internals/models"
	"github.com/israelalagbe/code-gen/internals/utils"
	"github.com/israelalagbe/code-gen/validations"
)

func main() {

	defaultAppDir, targetDir := libs.GetAppDirectoryInfo()

	fmt.Println("Default App Dir:", defaultAppDir)
	fmt.Println("Target Dir:", targetDir)

	var questions *inquire.Questions = nil
	questionModel := models.NewQuestionModel(targetDir)

	questions = inquire.Query()
	questions.Input(&questionModel.TableName, "Table name", validations.Required)
	questions.Input(&questionModel.ModelName, "Model name", validations.Required)
	questions.Input(&questionModel.Path, "Base directory", validations.Required)

	fields := ""
	questions.Input(&fields, "Fields", validations.Required)

	questions.Exec()

	questionModel.Properties = utils.ParseFields(fields)

	for index, item := range questionModel.Items {
		questions = inquire.Query()
		questions.YesNo(&item.Included, "Include "+item.Name)
		questions.Exec()

		item.Path = strings.Replace(item.Path, "{{.TableName}}", questionModel.TableName, 1)
		questionModel.Items[index] = item
	}

	for _, item := range questionModel.Items {
		if !item.Included {
			continue
		}

		absFilePath := path.Join(targetDir, item.Path)

		if libs.FileExists(absFilePath) {
			questions = inquire.Query()
			questions.YesNo(&item.Included, "File "+item.Path+" already exists. Overwrite?")
			questions.Exec()
		}

		if !libs.FileExists(path.Dir(absFilePath)) {
			libs.CreateDirectory(path.Dir(absFilePath))
		}

		result := libs.RenderTemplate(path.Join("templates", item.Name+".txt"), questionModel)

		libs.WriteFile(path.Join(targetDir, item.Path), result)
	}

	fmt.Println(questionModel)
}

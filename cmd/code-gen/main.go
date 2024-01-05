package main

import (
	"fmt"

	"github.com/burl/inquire"
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

	// for {
	// 	addAnotherProperty := false
	// 	questions = inquire.Query()

	// 	property := models.Property{}

	// 	questions.Input(&property.Name, "Property name", validations.Required)
	// 	questions.Input(&property.Type, "Property type", validations.Required)

	// 	questions.YesNo(&addAnotherProperty, "Add another property?")
	// 	questions.Exec()

	// 	if !addAnotherProperty {
	// 		break
	// 	}

	// 	questionModel.Properties = append(questionModel.Properties, models.Property{})
	// }

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

	fmt.Println(questionModel)
}

package main

import (
	"github.com/burl/inquire"
	"github.com/israelalagbe/code-gen/internals"
	"github.com/israelalagbe/code-gen/validations"
)

func main() {
	questionModel := internals.NewQuestionModel("users", "User")
	questions := inquire.Query()
	questions.Input(&questionModel.TableName, "Table name", validations.Required)
	questions.Input(&questionModel.ModelName, "Model name?", validations.Required)
	questions.Input(&questionModel.Path, "Base directory", validations.Required)
	questions.Exec()

	println("Your name is: ", questionModel.ModelName)
}

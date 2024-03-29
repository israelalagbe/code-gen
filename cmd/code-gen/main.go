package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/burl/inquire"
	"github.com/iancoleman/strcase"
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

		item.Path = strings.Replace(item.Path, "{{.ModelName}}", strcase.ToKebab(questionModel.ModelName), 1)
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

		if !item.Included {
			continue
		}

		if !libs.FileExists(path.Dir(absFilePath)) {
			libs.CreateDirectory(path.Dir(absFilePath))
		}

		renderProperties := []models.TemplateRenderDataProperty{}

		for _, property := range questionModel.Properties {
			dataType, dbType := utils.MapToDomainType(property.Type)
			comment := "The " + strcase.ToSnake((property.Name)+" of the "+questionModel.ModelName+" table")

			if property.Type == "uuid" {
				foreignTableName := strings.TrimSuffix(property.Name, "_id")
				foreignTableName = strings.TrimSuffix(foreignTableName, "id")
				comment = "Foreign key to the " + foreignTableName + " table"
			}

			renderProperties = append(renderProperties, models.TemplateRenderDataProperty{
				Name:         strcase.ToLowerCamel(property.Name),
				HumanizeName: utils.Humanize(property.Name),
				ColumName:    strcase.ToSnake(property.Name),
				Type:         dataType,
				DBTypeName:   dbType,
				Comment:      comment,
			})
		}

		data := models.TemplateRenderData{
			TableName:             questionModel.TableName,
			ModelName:             strcase.ToCamel(questionModel.ModelName),
			LowerModelName:        strcase.ToLowerCamel(questionModel.ModelName),
			SnakeCaseName:         strcase.ToSnake(questionModel.ModelName),
			HypenCaseName:         strcase.ToKebab(questionModel.ModelName),
			SentenceCaseName:      utils.Humanize(questionModel.ModelName),
			TitleSentenceCaseName: utils.HumanizeSentence(questionModel.ModelName),
			Properties:            renderProperties,
		}

		result := libs.RenderTemplate(path.Join(defaultAppDir, "templates", item.Name+".txt"), data)

		libs.WriteFile(absFilePath, result)
	}

	fmt.Println(questionModel)
}

package models

type QuestionModel struct {
	TableName  string
	ModelName  string
	Path       string
	Properties []Property
	Items      []Item
}

type Property struct {
	Name string
	Type string
}

type Item struct {
	Name     string // name of the item
	Path     string // path to the file
	Included bool   // true if the item should be included in the model
}

const (
	MODE_WRITE int = iota
	MODE_APPEND
)

func NewQuestionModel(path string) *QuestionModel {
	items := []Item{
		{Name: "repositories", Included: true, Path: "repositories/{{.TableName}}.go"},
	}

	return &QuestionModel{Items: items, Path: path, Properties: []Property{}}
}

func hello() {}

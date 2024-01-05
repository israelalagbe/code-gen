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

func NewQuestionModel() *QuestionModel {
	items := []Item{
		{Name: "Create", Path: "create.go", Included: true},
		{Name: "Update", Path: "update.go", Included: true},
	}

	return &QuestionModel{Items: items, Path: "./", Properties: []Property{}}
}

func hello() {}

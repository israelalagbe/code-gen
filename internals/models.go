package internals

type QuestionModel struct {
	TableName string
	ModelName string
	Path      string
	items     []Item
}

type Item struct {
	name     string // name of the item
	path     string // path to the file
	included bool   // true if the item should be included in the model
}

const (
	MODE_WRITE int = iota
	MODE_APPEND
)

func NewQuestionModel(tableName string, modelName string) *QuestionModel {
	items := []Item{
		{name: "Create", path: "create.go", included: true},
		{name: "Update", path: "update.go", included: true},
	}

	return &QuestionModel{items: items, Path: "./"}
}

func hello() {}

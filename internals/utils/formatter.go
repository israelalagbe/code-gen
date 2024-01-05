package utils

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func SentenceCase(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func Humanize(str string) string {
	return strings.Replace(strcase.ToSnake(str), "_", " ", -1)
}

func HumanizeSentence(str string) string {
	return SentenceCase(Humanize(str))
}

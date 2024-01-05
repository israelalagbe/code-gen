package utils

import (
	"strings"

	"github.com/israelalagbe/code-gen/internals/models"
)

func parseField(field string) (string, string) {
	splitField := strings.Split(field, ":")
	if len(splitField) < 2 {
		return splitField[0], "string"
	}
	return splitField[0], splitField[1]
}

func ParseFields(fields string) []models.Property {
	splitFields := strings.Split(fields, ",")
	properties := []models.Property{}
	for _, field := range splitFields {
		name, typ := parseField(field)
		properties = append(properties, models.Property{Name: name, Type: typ})
	}
	return properties
}

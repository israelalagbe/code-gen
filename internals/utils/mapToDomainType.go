package utils

func MapToDomainType(name string) (string, string) {
	switch name {
	case "string":
		return "string", "STRING(255)"
	case "int":
		return "number", "INTEGER"
	case "float":
		return "number", "FLOAT"
	case "bool":
		return "boolean", "BOOLEAN"
	case "decimal":
		return "string", "DECIMAL"
	case "text":
		return "string", "TEXT"
	default:
		return "string", "STRING"
	}
}

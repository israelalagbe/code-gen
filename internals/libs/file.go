package libs

import (
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func WriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func CreateDirectory(path string) {
	err := os.MkdirAll(path, 0755)

	if err != nil {
		panic(err)
	}
}

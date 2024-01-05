package libs

import (
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func WriteFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		return err
	}

	return nil
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

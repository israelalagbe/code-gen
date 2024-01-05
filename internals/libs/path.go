package libs

import (
	"os"
	"path/filepath"
)

func GetAppDirectoryInfo() (string, string) {
	executableFilePath, err := os.Executable()

	if err != nil {
		panic(err)
	}

	defaultAppDir := filepath.Dir(executableFilePath)

	currentWorkingDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	targetDir := currentWorkingDirectory

	if len(os.Args) > 2 {
		defaultAppDir = os.Args[1]
		targetDir = os.Args[2]
	}

	return defaultAppDir, targetDir
}

package utils

import (
	"os"
)

// ファイルを読み込む
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ファイルに書き込む
func WriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

// ファイルが存在するかチェックする
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

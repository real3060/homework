package storage

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func WriteFile(data []byte, fileName string) (bool, error) {
	file, err := os.Create("bills/" + fileName)
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return false, err
	}
	return true, nil
}
func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open("bills/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetLastFileName() (string, error) {
	dir := "./bills" // укажи нужную директорию
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Ошибка чтения директории:", err)
		return "", err
	}

	var (
		lastFile string
		lastDate time.Time
		found    bool
	)

	for _, entry := range files {
		name := entry.Name()
		if strings.HasPrefix(name, "bill-") {
			dateStr := strings.TrimPrefix(name, "bill-")
			dateStr = strings.TrimSuffix(dateStr, ".json")
			t, err := time.Parse("02.01.06", dateStr)
			if err != nil {
				// Если не удалось распарсить — пропускаем
				continue
			}
			if !found || t.After(lastDate) {
				lastDate = t
				lastFile = name
				found = true
			}
		}
	}

	return lastFile, nil
}

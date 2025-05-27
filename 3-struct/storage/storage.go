package storage

import (
	"demo/json/bins"
	"encoding/json"
	"fmt"
	"os"
)

const StorageFilename = "storage.json"

// Чтение списка bin в виде джсон из локального файла
func SaveToFile(bin bins.Bin) error {
	var savedStorageBins bins.Bins
	var bytes []byte
	var err error
	if _, err := os.Stat(StorageFilename); err != nil {
		fmt.Println("файла не существует, создаем")
		savedStorageBins = bins.Bins{
			Bins: []*bins.Bin{},
		}
		savedStorageBins.Bins = append(savedStorageBins.Bins, &bin)
	} else {
		// файл существует, перезаписать
		data, err := os.ReadFile(StorageFilename)
		if err != nil {
			return err
		}
		//  необходимо прочитать из файла и присвоить в переменную
		err = json.Unmarshal(data, &savedStorageBins)
		if err != nil {
			fmt.Println("не удалось распарсить данные из хранилища")
			return err
		}
		// получили из сторы данные, Необходимо обновить их полученными данными
		savedStorageBins.Bins = append(savedStorageBins.Bins, &bin)
	}

	bytes, err = json.MarshalIndent(savedStorageBins, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(StorageFilename, bytes, 0600)
	if err != nil {
		fmt.Println("ошибка записи в файл")
		return err
	}
	fmt.Printf("сохранение в файл %s успешно\n", StorageFilename)
	return nil
}

func ReadFromFile() ([]byte, error) {
	file, err := os.ReadFile(StorageFilename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

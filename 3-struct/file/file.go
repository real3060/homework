package file

import (
	"demo/json/bins"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// чтение любого файла
// проверка что это json расширение файла

func ReadFromFile(name string) (bins.Bins, error) {
	isJson := checkFileExtension(name)
	if !isJson {
		return bins.Bins{}, errors.New("укажите файл с расширением .json и повторите снова")
	}
	bytes, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return bins.Bins{}, err
	}
	var binsParsed bins.Bins
	err = json.Unmarshal(bytes, &binsParsed)
	if err != nil {
		return bins.Bins{}, err
	}
	return binsParsed, nil

}
func checkFileExtension(filename string) bool {
	ext := filepath.Ext(filename)
	if ext != ".json" {
		return false
	}
	return true
}

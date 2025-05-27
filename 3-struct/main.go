package main

import (
	"demo/json/bins"
	"demo/json/file"
	"demo/json/storage"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	for {
		fmt.Println("__________что вы хотите сделать__________")
		fmt.Println("1. Создать корзину")
		fmt.Println("2. Показать содержимое корзин")
		fmt.Println("3. Загрузить корзины из файла в хранилище")
		fmt.Println("4. Выход")
		var input int
		fmt.Scan(&input)
		switch input {
		case 1:
			createBin()
			continue
		case 2:
			showBins()
			continue
		case 3:
			getBinFromFile()
			continue
		case 4:
			os.Exit(0)
		}
	}

}
func createBin() {
	var binName string
	var isPrivate bool
getBin:
	for {
		fmt.Println("введите название корзины")
		fmt.Scan(&binName)
		if binName == "" {
			fmt.Println("название не может быть пустым")
			continue getBin
		}
	isPrivate:
		for {
			var input int
			fmt.Println("корзина будет приватной? ")
			fmt.Println("1. true")
			fmt.Println("2. false")
			fmt.Scan(&input)
			switch input {
			case 1:
				isPrivate = true
			case 2:
				isPrivate = false
			default:
				continue isPrivate
			}
			break
		}
		break
	}
	bin := bins.NewBin("", binName, isPrivate)
	err := storage.SaveToFile(*bin)
	if err != nil {
		fmt.Println("не удалось сохранить в файл storage.json")
	}

}
func showBins() {
	data, err := storage.ReadFromFile()
	if err != nil {
		fmt.Println("не удалось прочитать данные из хранилища")
		return
	}
	var binsStruct bins.Bins
	fmt.Println(data)
	err = json.Unmarshal(data, &binsStruct)
	if err != nil {
		fmt.Println("не удалось разобрать данные из хранилища, бардак какой-то!")
		return
	}
	binsStruct.ShowBins()
}
func getBinFromFile() {
	fmt.Println("напишите название файла с расширением .json")
	var fileName string
	fmt.Scan(&fileName)
	binsFromFile, err := file.ReadFromFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	binsFromFile.ShowBins()
	for _, bin := range binsFromFile.Bins {
		err := storage.SaveToFile(*bin)
		if err != nil {
			fmt.Println(err)
			break
		}
		color.Cyan("успешно добавлено из вашего файла в хранилище")
	}
}

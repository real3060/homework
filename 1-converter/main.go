package main

import (
	"fmt"
	"strings"
)

type currenciesMap = map[string]map[string]float64

var currencies currenciesMap = currenciesMap{
	"usd": {"rub": 80.50, "eur": 0.88},
	"eur": {"rub": (80.50 / 0.88)},
}

func main() {
	getMenu()
}

func getMenu() {
	fmt.Println("_____Currency callulator______")
	fmt.Println("Выберите пункт меню чтобы продолжить:")
	fmt.Println("1. Поменять валюту")
	fmt.Println("2. Выход")
	getMenuInput()
}
func getMenuInput() {
	var menuInput string
	fmt.Scan(&menuInput)
	for {
		var err error
		if menuInput == "1" {
			_, _, _, err = getUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if menuInput == "2" {
			panic("Вы вышли из программы")
		}
		break

	}
}
func getRedeemableCurrency() (string, string, error) {
	var src string
	fmt.Println("Введите валюту которую отдаете usd/eur")
	fmt.Scan(&src)
	src = strings.ToLower(src)
	switch src {
	case "usd":
		{
			dstCurrency, err := getDstCurrency(src, []string{"rub", "eur"})
			return src, dstCurrency, err
		}
	case "eur":
		{
			dstCurrency, err := getDstCurrency(src, []string{"rub"})

			return src, dstCurrency, err
		}
	default:
		return "", "", nil
	}
}
func getCurrencyCount() float64 {
	var count float64
	for {
		fmt.Println("Введите кол-во")
		fmt.Scan(&count)
		if count != 0 {
			break
		}
	}
	return count
}
func calculateCurrency(src, dstCurrency string, count float64) {
	if _, ok := currencies[src][dstCurrency]; !ok {
		fmt.Println("Нет такой валюты для конвертации")
		return
	}
	printResult(src, dstCurrency, count, currencies[src][dstCurrency]*float64(count))
}
func getUserInput() (float64, string, string, error) {
	var count float64
	var src, dstCurrency string
	var err error
	for {
		src, dstCurrency, err = getRedeemableCurrency()
		if src == "" && dstCurrency == "" {
			continue
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		count = getCurrencyCount()
		calculateCurrency(src, dstCurrency, count)
		break
	}
	return count, src, dstCurrency, nil
}
func printResult(src, dstCurrency string, count float64, result float64) {
	fmt.Printf("Вы конвертируете %v %s в %.2f %s \n", count, src, result, dstCurrency)
}
func getDstCurrency(src string, values []string) (string, error) {
	var dst string

	for {
		fmt.Printf("Введите валюту которую хотите получить: %s -> %v?\n", src, values)
		fmt.Scan(&dst)
		dst = strings.ToLower(dst)
		switch dst {
		case "rub", "eur":
			return dst, nil
		default:
			{
				fmt.Println("Валюта выбрана некорректно")
				continue
			}
		}
	}
}

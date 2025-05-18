package main

import (
	"fmt"
	"os"
	"strings"
)

const USDtoEUR, USDtoRUB float64 = 0.88, 80.50

func main() {
	getMenu()
}

func usdToEur(count int) float64 {
	return USDtoEUR * float64(count)
}
func usdToRub(c int) float64 {
	return USDtoRUB * float64(c)
}
func eurToRub(c int) float64 {
	return (usdToRub(c) / usdToEur(c)) * float64(c)
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
		if menuInput == "1" {
			getUserInput()
			continue
		} else if menuInput == "2" {
			fmt.Println("Программа завершена")
			os.Exit(0)
		}
		break

	}
}
func getRedeemableCurrency() (string, string) {
	var src string
	fmt.Println("Введите валюту которую отдаете usd/eur")
	fmt.Scan(&src)
	src = strings.ToLower(src)
	switch src {
	case "usd":
		{
			dstCurrency := getDstCurrency(src, []string{"rub", "eur"})
			return src, dstCurrency
		}
	case "eur":
		{
			dstCurrency := getDstCurrency(src, []string{"rub"})

			return src, dstCurrency
		}
	default:
		return "", ""
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
	if src == "usd" {
		switch dstCurrency {
		case "rub":
			printResult(src, dstCurrency, count, usdToRub(int(count)))
		case "eur":
			printResult(src, dstCurrency, count, usdToEur(int(count)))
		default:
			fmt.Println("Нет такой валюты для конвертации")
		}
	} else if src == "eur" {
		switch dstCurrency {
		case "rub":
			printResult(src, dstCurrency, count, eurToRub(int(count)))
		default:
			fmt.Println("Нет такой валюты для конвертации")
		}
	}
}
func getUserInput() (float64, string, string) {
	var count float64
	var src, dstCurrency string
	for {
		var err error
		src, dstCurrency = getRedeemableCurrency()
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
	return count, src, dstCurrency
}
func printResult(src, dstCurrency string, count float64, result float64) {
	fmt.Printf("Вы конвертируете %s в кол-ве %v в %.2f %s", src, count, result, dstCurrency)
}
func getDstCurrency(src string, values []string) string {
	var dst string

	for {
		fmt.Printf("Введите валюту которую хотите получить: %s -> %v?\n", src, values)
		fmt.Scan(&dst)
		dst = strings.ToLower(dst)
		switch dst {
		case "rub", "eur":
			{
				return dst
			}
		default:
			{
				fmt.Println("Валюта выбрана некорректно")
				continue
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var extMap = map[string]func(){
	"avg":    avg,
	"sum":    sumWrapper,
	"median": med,
}

func main() {
	for {
		fmt.Println("Выберите операцию: ")
		fmt.Println("avg")
		fmt.Println("sum")
		fmt.Println("median")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			fmt.Println("Значение не может быть пустой строкой")
			continue
		}
		if _, ok := extMap[input]; !ok {
			fmt.Println("такого значения не существует, попробуйте еще раз")
			continue
		}
		extMap[input]()
		//switch input {
		//case 1:
		//	avg()
		//case 2:
		//	sum([]int{})
		//case 3:
		//	med()
		//default:
		//	continue
		//}
	}
}
func avg() {
	fmt.Println("Вы выбрали avg")
	slice := []int{}
	var value int
	for {
		fmt.Println("Введите числа для рассчета, для окончания ввода нажмите 0")
		fmt.Scan(&value)
		if value < 0 {
			continue
		}
		if value == 0 {
			fmt.Println(sum(slice) / len(slice))
			break
		}
		slice = append(slice, value)
	}
}
func sum(s []int) int {
	summ := 0
	if len(s) > 0 {
		for _, v := range s {
			summ += v
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Вы выбрали получение суммы")
		for {
			fmt.Println("Введите числа для рассчета(Enter для выхода):")
			scanner.Scan()
			value := scanner.Text()
			if strings.TrimSpace(value) == "" {
				fmt.Println("Сумма значений равна = ", summ)
				break
			}
			num, err := strconv.Atoi(value)
			if err != nil || num < 0 {
				continue
			}
			summ += num
			continue
		}
	}

	return summ
}
func med() {
	fmt.Println("Вы выбрали получение медианы")
	var slice []int
	for {
		fmt.Println("Введите числа для рассчета, для окончания ввода нажмите 0")
		var value int
		fmt.Scan(&value)
		if value < 0 {
			continue
		}
		if value == 0 {
			slices.SortFunc(slice, func(a, b int) int {

				if a > b {
					return -1
				}
				return 0
			})
			fmt.Println("Медианное число = ", slice[len(slice)/2])
			break
		}

		slice = append(slice, value)
	}
}
func sumWrapper() {
	var value []int
	fmt.Println(sum(value))
}

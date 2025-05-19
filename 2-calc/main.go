package main

import (
	"fmt"
	"slices"
)

func main() {
	for {

		fmt.Println("Выберите операцию: ")
		fmt.Println("1.AVG")
		fmt.Println("2.SUM")
		fmt.Println("3.Median")
		var input int
		fmt.Scan(&input)
		if input <= 0 {
			fmt.Println("Введите положительное значение")
			continue
		}
		switch input {
		case 1:
			avg()
		case 2:
			sum([]int{})
		case 3:
			med()
		default:
			continue
		}
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
		fmt.Println("Вы выбрали получение суммы")
		var value int
		for {
			fmt.Println("Введите числа для рассчета, для окончания ввода нажмите 0")
			fmt.Scan(&value)
			if value < 0 {
				continue
			}
			if value == 0 {
				fmt.Println("Сумма значений равна = ", summ)
				break
			}
			summ += value
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

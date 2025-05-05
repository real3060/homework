package main

import "fmt"

const USD, EUR float32 = 0.88, 80.50

func main() {
	var usdTOrub, usdTOeur, eurTOrub = 10, 15, 20

	fmt.Printf("долларов %v в евро: %v \n", usdTOeur, usdToEur(usdTOeur))
	fmt.Printf("долларов %v в рубли: %v \n", usdTOrub, usdToRub(usdTOrub))
	fmt.Printf("евро %v в рубли: %v \n", eurTOrub, eurToRub(eurTOrub))
}

func usdToEur(count int) float32 {
	return USD * float32(count)
}
func usdToRub(c int) float32 {
	return EUR * float32(c)
}
func eurToRub(c int) float32 {
	return usdToRub(c) / usdToEur(c)
}

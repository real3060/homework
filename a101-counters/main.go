package main

import (
	"a101-counters/storage"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Electricity struct {
	T1, T2, T3 float32
}
type Water struct {
	Cold, Hot, WasteWater float32
}

type Counters struct {
	Water       Water
	Electricity Electricity
	Heater      float32
	Internet    float32
	Total       float32
}
type Values struct {
	Water       Water
	Electricity Electricity
	Heater      float32
	Internet    float32
	TotalPrice  float32
	CreatedAt   string
}

var coefficient = map[string]map[string]float32{
	"electricity": map[string]float32{
		"T1": 8.94,
		"T2": 3.02,
		"T3": 6.15,
	},
	"water": map[string]float32{
		"cold":       59.8,
		"hot":        234.56,
		"wasteWater": 45.91,
	},
	"Heater": map[string]float32{
		"apartment": 2803.24,
	},
	"internet": map[string]float32{
		"inet": 610,
	},
}

func main() {
	var previousValues Values
	lastFileName, err := storage.GetLastFileName()
	if err != nil {
		fmt.Println("не удалось найти последний файл", err.Error())
	}
	prev, err := storage.ReadFile(lastFileName)
	if err != nil {
		fmt.Println("не удалось открыть последний файл", err.Error())
	}
	err = json.Unmarshal(prev, &previousValues)
	if err != nil {
		fmt.Println("не удалось анмаршалить последние данные из файла", err.Error())
	}
	fmt.Println("Предыдущие показания были: ")
	fmt.Printf("%+v\n", previousValues)

	var counter Counters = Counters{
		Internet: coefficient["internet"]["inet"],
	}
	currValue := Values{
		Internet: coefficient["internet"]["inet"],
	}

	for {
		fmt.Println("__________Показания счетчиков__________: ")
		getWater(&counter, &currValue, previousValues.Water)
		getElectricity(&counter, &currValue, previousValues.Electricity)
		getHeater(&counter, &currValue, previousValues.Heater)
		break
	}
	fmt.Println("Электроэнергия: ")
	electr, err := json.MarshalIndent(counter.Electricity, "", "  ")
	if err != nil {
		fmt.Println("чет не так разобралось")
	}
	color.Red(string(electr))

	fmt.Println("Водоснабжение: ")
	water, _ := json.MarshalIndent(counter.Water, "", "  ")
	if err != nil {
		fmt.Println("вода не так разобралась")
	}
	color.Blue(string(water))

	fmt.Print("Отопление: ")
	heater, _ := json.MarshalIndent(counter.Heater, "", "  ")
	color.HiRed(string(heater) + "\n")

	fmt.Println("Интернет: ", counter.Internet)

	var electricity = counter.Electricity.T1 + counter.Electricity.T2 + counter.Electricity.T3
	var waterService = counter.Water.WasteWater + counter.Water.Hot + counter.Water.Cold
	var heaterService = counter.Heater
	counter.Total = heaterService + waterService + electricity + counter.Internet
	currValue.TotalPrice = counter.Total
	currValue.CreatedAt = time.Now().Format("02.01.06")

	// save
	currValuesMarshaled, err := json.MarshalIndent(currValue, "", "  ")
	if err != nil {
		fmt.Println("ошибка маршалинга")
	}
	fileName := "bill-" + currValue.CreatedAt + ".json"
	storage.WriteFile(currValuesMarshaled, fileName)
	if err != nil {
		color.Red("ошибка записи в файл", err.Error())
	}
	// save

	fmt.Printf("Итого: %.2f \n", currValue.TotalPrice)

}

func getElectricity(c *Counters, currValue *Values, previousValues Electricity) {
	var electricityT1 float32
	var electricityT2 float32
	var electricityT3 float32
electricity:
	for {
		fmt.Println("введите значение для электричества: ")
		fmt.Println("T1: ")
		fmt.Scan(&electricityT1)
		if electricityT1 < 0.0 {
			fmt.Println("должно быть больше 0")
			continue electricity
		}
		fmt.Println("T2: ")
		fmt.Scan(&electricityT2)
		if electricityT2 < 0.0 {
			fmt.Println("должно быть больше 0")
			continue electricity
		}
		fmt.Println("T3: ")
		fmt.Scan(&electricityT3)
		if electricityT3 < 0.0 {
			fmt.Println("должно быть больше 0")
			continue electricity
		}

		break
	}

	currValue.Electricity = Electricity{
		T1: electricityT1,
		T2: electricityT2,
		T3: electricityT3,
	}
	countForCalc := Electricity{
		T1: electricityT1 - previousValues.T1,
		T2: electricityT2 - previousValues.T2,
		T3: electricityT3 - previousValues.T3,
	}
	fmt.Println("Использовано электроэнергии: ")
	fmt.Printf("%+v\n", countForCalc)

	c.Electricity = Electricity{
		T1: coefficient["electricity"]["T1"] * countForCalc.T1,
		T2: coefficient["electricity"]["T2"] * countForCalc.T2,
		T3: coefficient["electricity"]["T3"] * countForCalc.T3,
	}
}

func getWater(c *Counters, currValue *Values, previousValues Water) {
	var coldWater float32
	var hotWater float32
water:
	for {
		fmt.Println("введите показания счетчика ХВС: ")
		fmt.Scan(&coldWater)
		if coldWater < 0.0 {
			fmt.Println("введите положительное значение")
			continue water
		}
		fmt.Println("введите показания счетчика ГВС: ")
		fmt.Scan(&hotWater)
		if hotWater < 0.0 {
			fmt.Println("должно быть больше нуля")
			continue water
		}
		break
	}
	currValue.Water = Water{
		Cold:       coldWater,
		Hot:        hotWater,
		WasteWater: coldWater + hotWater,
	}
	countForCalc := Water{
		Cold:       coldWater - previousValues.Cold,
		Hot:        hotWater - previousValues.Hot,
		WasteWater: currValue.Water.WasteWater - previousValues.WasteWater,
	}
	fmt.Println("Использовано воды: ")
	fmt.Printf("%+v\n", countForCalc)
	c.Water = Water{
		Cold:       coefficient["water"]["cold"] * countForCalc.Cold,
		Hot:        coefficient["water"]["hot"] * countForCalc.Hot,
		WasteWater: coefficient["water"]["wasteWater"] * countForCalc.WasteWater,
	}
}

func getHeater(c *Counters, currValue *Values, previousValues float32) {
	var heater float32
heater:
	for {
		fmt.Println("показания термодатчика: ")
		fmt.Scan(&heater)
		if int(heater) == 0 {
			currValue.Heater = previousValues
			break
		}
		if heater < 0.0 {
			fmt.Println("должно быть больше 0")
			continue heater
		}
		break
	}
	if currValue.Heater == previousValues {
		color.Green("за указанный период потребления не было")
		return
	}
	currValue.Heater = heater
	countForCalc := currValue.Heater - previousValues
	fmt.Println("Использовано тепла: ", countForCalc)
	c.Heater = coefficient["Heater"]["apartment"] * countForCalc
}

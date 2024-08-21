package main

import (
	"fmt"
	"math"
)

func main() {
	// Радиус трубки и тора для полного тора
	r_total := 3.0
	R_total := 9.0

	// Радиус трубки и тора для части Саши
	r_sasha := 1.5
	R_sasha := 7.5

	// Объем полного тора
	volumeTotal := 2 * math.Pi * math.Pi * R_total * r_total * r_total
	volumeTotal = volumeTotal / 2.0

	// Объем части Саши
	volumeSasha := 2 * math.Pi * math.Pi * R_sasha * r_sasha * r_sasha

	// Разница объемов
	difference := volumeTotal - volumeSasha

	// Разделить разницу на π
	result := difference / math.Pi

	// Округлить до ближайшего целого числа
	roundedResult := math.Round(result)

	// Вывести результат
	fmt.Printf("%.0f\n", roundedResult)
}

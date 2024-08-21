package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные данные
	a := 100.0 // метров
	b := 200.0 // метров
	L := 500.0 // метров
	V := 5.0   // метров в секунду

	// Функция для вычисления времени пути
	time := func(x float64) float64 {
		lawnDistance := math.Sqrt(a*a + x*x)           // расстояние по лугу
		forestDistance := math.Sqrt(b*b + (L-x)*(L-x)) // расстояние через лес
		lawnTime := lawnDistance / V
		forestTime := forestDistance / (V / 3)
		return lawnTime + forestTime
	}

	// Поиск минимального времени
	minTime := math.Inf(1)
	for x := 0.0; x <= L; x += 0.01 {
		currentTime := time(x)
		if currentTime < minTime {
			minTime = currentTime
		}
	}

	fmt.Printf("Самое короткое время пути: %.3f секунд\n", minTime)
}

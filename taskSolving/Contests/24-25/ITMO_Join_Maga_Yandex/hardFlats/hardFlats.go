package main

import (
	"fmt"
	"math"
)

func surfaceZ(x, y float64) float64 {
	// Вычисление z по уравнению
	return math.Pow((math.Pow(x, 8)+math.Pow(y, 8))/(math.Pow(x, 4)+math.Pow(y, 4)), 1/6.0)
}

func main() {
	// Определяем диапазон поиска и шаг
	min := -1000000000.0
	max := 10000000000.0
	step := 0.1

	maxZ := -math.Inf(1)
	points := []struct {
		x, y, z float64
	}{}

	for x := min; x <= max; x += step {
		for y := min; y <= max; y += step {
			z := surfaceZ(x, y)
			if z > maxZ {
				maxZ = z
				points = []struct {
					x, y, z float64
				}{{x, y, z}}
			} else if z == maxZ {
				points = append(points, struct {
					x, y, z float64
				}{x, y, z})
			}
		}
	}

	// Вычисляем сумму координат
	sum := 0.0
	for _, p := range points {
		sum += p.x + p.y + p.z
	}

	// Выводим ответ с точностью 3 знака после запятой
	fmt.Printf("%.3f\n", sum)
}

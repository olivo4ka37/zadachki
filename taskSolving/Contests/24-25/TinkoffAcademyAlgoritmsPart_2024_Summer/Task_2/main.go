package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	raspisanie := map[string]int{}
	raspisanie["MON"] = 0
	raspisanie["TUE"] = 1
	raspisanie["WED"] = 2
	raspisanie["THU"] = 3
	raspisanie["FRI"] = 4
	raspisanie["SAT"] = 5
	raspisanie["SUN"] = 6
	month := make([]bool, 28)

	for i := 0; i < 4; i++ {
		scan.Scan()
		line := scan.Text()
		if line != "" {
			days := strings.Split(line, " ")
			for _, day := range days {
				if day != "" {
					month[raspisanie[day]+7*i] = true
				}
			}
		}
	}

	bestResult := 0
	currentResult := 0
	startDay := 0
	currentStart := 0

	for i := 0; i < 28; i++ {
		if !month[i] {
			if currentResult == 0 {
				currentStart = i + 1
			}
			currentResult++
		} else {
			if currentResult > bestResult {
				bestResult = currentResult
				startDay = currentStart
			}
			currentResult = 0
		}
	}

	if currentResult > bestResult {
		bestResult = currentResult
		startDay = currentStart
	}

	if bestResult == 0 {
		out.WriteString("0 0")
	} else {
		result := fmt.Sprintf("%d %d\n", startDay, startDay+bestResult-1)
		out.WriteString(result)
	}

}

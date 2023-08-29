package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var t, n, startTime, endTime int
	var otrezok string
	var result bool

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)

		result = true

		var startArray [20001]int
		var endArray [20001]int

		for i := range startArray {
			startArray[i] = -1
		}

		for i := range endArray {
			endArray[i] = -1
		}

		schetchik := 0

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &otrezok)

			if result == false {
				continue
			}

			strArray := strings.Split(otrezok, "-")

			startTime = podschetVSeconds(strArray[0])
			if startTime == -1 {
				result = false
				continue
			}
			endTime = podschetVSeconds(strArray[1])
			if endTime == -1 {
				result = false
				continue
			}

			if startTime > endTime {
				result = false
				continue
			}

			for z := 0; z < len(startArray); z++ {
				if startArray[z] == -1 {
					break
				} else if startTime == startArray[z] || endTime == startArray[z] || startTime == endArray[z] || endTime == endArray[z] {
					result = false
					continue
				} else if startTime <= startArray[z] && endTime >= startArray[z] || startTime <= endArray[z] && endTime >= endArray[z] {
					result = false
					continue
				} else if startTime >= startArray[z] && endTime <= startArray[z] || startTime >= endArray[z] && endTime <= endArray[z] {
					result = false
					continue
				} else if startTime >= startArray[z] && endTime >= startArray[z] && startTime <= endArray[z] && endTime <= endArray[z] {
					result = false
					continue
				}
			}

			startArray[schetchik] = startTime
			endArray[schetchik] = endTime
			schetchik++
		}

		if result != false {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
		defer out.Flush()
	}

}

func podschetVSeconds(strArray string) int {
	chasti := strings.Split(strArray, ":")

	hours, _ := strconv.Atoi(chasti[0])
	if hours < 0 || hours > 23 {
		return -1
	}

	minutes, _ := strconv.Atoi(chasti[1])
	if minutes < 0 || minutes > 59 {
		return -1
	}

	seconds, _ := strconv.Atoi(chasti[2])
	if seconds < 0 || seconds > 59 {
		return -1
	}

	totalSeconds := hours*3600 + minutes*60 + seconds
	return totalSeconds
}

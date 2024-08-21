package CodeWars

import "fmt"

// ChooseBestSum is a recursive function to find the kvvalues in ls that
// sum to the greatest number <= t. returns -1 if not possible.
func ChooseBestSum(t, k int, ls []int) int {
	outerbest := -1
	for i, d := range ls {
		fmt.Println("d is euqal to", d, "i is equal to", i)
		// not enough remaining values for this d to work
		if len(ls) < k {
			continue
		}
		// recursively choose best from t-d, until final level k=1
		if k > 1 {
			innerbest := ChooseBestSum(t-d, k-1, ls[i+1:])
			// if no best available at lower level, this d cant work
			if innerbest < 0 {
				continue
			}
			d += innerbest
		}
		if d <= t && d > outerbest {
			outerbest = d
		}
	}
	return outerbest
}

package LeetCode

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: nil,
	}
}

func (this *RecentCounter) ping(t int) int {
	this.requests = append(this.requests, t)
	result := 0
	for _, n := range this.requests {
		if n >= t-3000 {
			result++
		}
	}

	return result
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

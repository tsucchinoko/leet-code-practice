package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	lower := t - 3000
	i := 0
	for i < len(this.queue) && this.queue[i] < lower {
		i++
	}
	if i > 0 {
		this.queue = this.queue[i:]
	}
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

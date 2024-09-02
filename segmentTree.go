package main

import "fmt"

type SegmentTree struct {
	size     int
	capacity int
	a        []int
}

func NewSegmentTree(v []int) *SegmentTree {
	var size = 0
	var n = len(v)
	for (1 << size) <= n {
		size++
	}
	size++
	size = (1 << size)
	var capacity = (size >> 1)
	var a = make([]int, size)
	for i := 0; i < n; i++ {
		a[capacity+i] = v[i]
	}
	for i := capacity - 1; i > 0; i-- {
		a[i] = a[i*2] + a[i*2+1]
	}
	return &SegmentTree{size, capacity, a}
}

func sum(cur int, left int, right int, l int, r int, t *SegmentTree) int64 {
	if right < l || left > r {
		return 0
	}
	if l <= left && r >= right {
		return int64(t.a[cur])
	}
	return sum(cur*2, left, (left+right)/2, l, r, t) + sum(cur*2+1, (left+right)/2+1, right, l, r, t)
}

func getSum(t *SegmentTree, l int, r int) int64 {
	return sum(1, 0, t.capacity-1, l, r, t)
}

func testSegmentTree() {
	var n int
	fmt.Scanf("%d", &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var t = NewSegmentTree(a)
	var q, l, r int
	fmt.Scanf("%d", &q)
	for q > 0 {
		q--
		fmt.Scanf("%d%d", &l, &r)
		l--
		r--
		fmt.Printf("%d ", getSum(t, l, r))
	}
}

func main() {
	testSegmentTree()
}

package main

import "fmt"

// CF 1131F

type DSU struct {
	a   []int
	res [][]int
}

func NewDSU(n int) *DSU {
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, 1)
		res[i][0] = i
	}
	return &DSU{a, res}
}

func get(dsu *DSU, x int) int {
	if dsu.a[x] == x {
		return x
	} else {
		dsu.a[x] = get(dsu, dsu.a[x])
	}
	return dsu.a[x]
}

func join(dsu *DSU, x int, y int) {
	x = get(dsu, x)
	y = get(dsu, y)
	if x != y {
		if len(dsu.res[x]) > len(dsu.res[y]) {
			x, y = y, x
		}
		dsu.a[x] = y
		dsu.res[y] = append(dsu.res[y], dsu.res[x]...)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	dsu := NewDSU(n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		join(dsu, a, b)
	}
	x := get(dsu, 0)
	for _, elem := range dsu.res[x] {
		fmt.Printf("%d ", elem+1)
	}
}

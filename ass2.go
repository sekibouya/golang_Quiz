package main

import "fmt"

type Integer int

func (i *Integer) f() int {
	*i += 2
	return int(*i)
}

func g(y []int, z int) int {
	y[1] = 20
	z = z * 10
	return z
}

func main() {
	var sum int
	const a = 5.0 / 3
	var b float64 = 5 / 3
	if a == b {
		sum++
	}

	var t bool
	const num = 1
	var i Integer
	for i < 13 {

		if !t && 7&(&i).f()<<1 < 1<<2+1 {
			sum += 10
		}
		i++
	}

	x := []int{1, 2, 3}
	y := append(x, 4)
	z := append(y, 5)
	y[0] = g(z, z[3])
	n := 1000000000000
	for _, v := range z {
		n = n / 100
		sum += v * n
	}

	fmt.Println(sum)
}

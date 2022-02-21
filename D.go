package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	t := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&t[i])
	}
	p := make([]int, n)
	tt := 0
	j := 0
	i := 0
	for {
		if p[i] == 0 {
			p[i] = t[j]
			j++

			fmt.Println(i, tt)
		}
		p[i] -= 1
		if len(t) == j {
			break
		}
		i++
		if i == n {
			i = 0
			tt++
		}

	}

}

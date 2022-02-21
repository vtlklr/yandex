package main

import "fmt"

func main() {
	var n, m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		return
	}

	t := make([]int, m)
	for i := 0; i < m; i++ {
		_, err := fmt.Scan(&t[i])
		if err != nil {
			return
		}
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

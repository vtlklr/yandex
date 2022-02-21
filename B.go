package main

import "fmt"

func main() {
	var i, k, n, x, y int
	ok := true
	fmt.Scan(&n)
	for ok {
		if i%2 != 0 && i != 0 {
			for j := 0; j < i && ok; j++ {
				x = x - 1
				k++

				//fmt.Println("-x", "счетчик", k, k == n, "координаты", x, y, "счетчик j", j, i, ok)
				if k == n {
					ok = false
				}
			}

			for j := 0; j < i && ok; j++ {
				y = y - 1
				k++

				//fmt.Println("-y", "счетчик", k, "координаты", x, y, "счетчик j", j, i, j < i, ok)
				if k == n {
					ok = false
				}
			}

		}
		if i%2 == 0 && i != 0 {
			for j := 0; j < i && ok; j++ {

				x = x + 1
				k++

				//fmt.Println("+x", "счетчик", k, "координаты", x, y, "счетчик j", j, i, j < i, ok)
				if k == n {
					ok = false
				}
			}

			for j := 0; j < i && ok; j++ {
				y = y + 1
				k++

				//fmt.Println("+y", "счетчик", k, "координаты", x, y, "счетчик j", j, i, j < i, ok)
				if k == n {
					ok = false
				}
			}
		}
		i++
		if k == n {
			ok = false
		}
		//fmt.Println(x, y, ok, k, n)
	}
	fmt.Println(x, y)

}

package main

import "fmt"

func main() {

	i := 0

	max := 10

	for i < max {
		if i == 5 {
			i++
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
		i++

	}

}

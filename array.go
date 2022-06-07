package main

import "fmt"

func main() {

	// tanpa harus di tentukan panjang array
	var array1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("ini adalah hasil array1 :", array1[1])

	// ditentukan panjang array

	var array2 = [3]string{"herdo", "dimas", "pratirto"}
	fmt.Println("ini adalah hasil array2 :", array2[1])

}

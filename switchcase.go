package main

import "fmt"

func main() {

	fmt.Print("golang berjalan pada ")

	switch osgolang := "linux"; osgolang {
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("tidak ada satu pun")

	}

}

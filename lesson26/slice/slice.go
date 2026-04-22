package main

import "fmt"

// 切片指针，修改切片本身的起始地址和容量等
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)

	fmt.Println(planets)
}

package main

import "fmt"

func main() {
	const name string = "Javascript"
	fmt.Println(name)

	const (
		port=5400
		host="localhost"
	)
	fmt.Println(port,host)
	
}
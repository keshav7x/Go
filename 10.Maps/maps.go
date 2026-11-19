package main

import "fmt"

//Maps : = Associated Data strucutre
func main() {
	//Creating map
	m:=make(map[string]string)
	m["name"]="goLang"
	m["areal"]="Backend"

	fmt.Println(m["name"])
}

package main

import "fmt"

func main() {
	// num := make([]int, 0, 2)
	// num = append(num, 1)
	// num = append(num, 2)
	// num = append(num, 3)
	// num = append(num, 4)
	// for k,v := range(num){
	// 	fmt.Println(k,v)
	// }
	// m:=map[string]int{}
	// m["Keshav"]=11
	// for k,v := range m{
	// 	fmt.Println(k,v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}

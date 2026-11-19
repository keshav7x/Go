package main

// import "fmt"

func add(a, b int)int {
	return a + b
}

func getLanguage()(string,string,string){
	return  "go","js","C"
}


func processIt() func(int)int {
	return func(i int)int {
		return  4
	}
}

func main() {
	// fmt.Println(add(4,5))
	// l1,l2,l3:=(getLanguage())
	// fmt.Println(l1,l2,l3)
	fn:=processIt()
	fn(4)

}
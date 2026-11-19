package main

import "fmt"

func main() {
	// age := 12
	// if age >= 18 {
	// 	fmt.Println("person is and adult")
	// }else{
	// 	fmt.Println("The person is teenager")
	// }

	role:="admin"
	hasPermisson:=true

	if(role=="admin" && hasPermisson){
		 fmt.Println("Access granted") 
	}else if(role=="admin" && !hasPermisson){
		fmt.Println("access rejected")
	}else{
		fmt.Println("You're not an admin")
	}
}
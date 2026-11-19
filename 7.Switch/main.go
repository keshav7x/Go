package main

import (
	"fmt"
)

func main() {
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// default:
	// 	fmt.Println("idk")
	// }
	// switch time.Now().Weekday(){
	// 	case time.Saturday , time.Sunday:
	// 		fmt.Println("It's weekend")
	// 	default:
	// 		fmt.Println("Workday")
	// }

	//Type switch

	whoami := func(i any){
		switch j:=i.(type){
			case int:
				fmt.Println("Integer",j)
			case bool:
				fmt.Println("Boolean",j)
			default:
				fmt.Println("idk")
		}
	}
	whoami(true)
}
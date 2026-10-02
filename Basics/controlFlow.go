package main
import "fmt"

func controlFlow(){
	age:=45
	if age>=18{
		fmt.Println("Adult")
	}else{
		fmt.Println("Not adult")
	}
}

func getAge() int{
	return 18
}

func initializationInsideIf(){
	if age:=getAge();age>=18{
		fmt.Println("Adult")
	}
}

func main(){
	
}

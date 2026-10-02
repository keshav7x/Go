package main

import "fmt"

func Arr1(){
	numbers:=[5]int{1,2,3,4,5}
	for i:=0;i<len(numbers);i++{
		fmt.Println(numbers[i])
	}
}


func Arr2(){
	numbers:=[...]int{10,20,30,40}
	for idx,val:= range numbers{
		fmt.Println(idx,val)
	}
}

func reverseAnArray(){
	numbers:=[...]int{10,20,30,40}
	for i := len(numbers) - 1; i >= 0; i-- {
    	fmt.Println(numbers[i])
	}
}

func main(){
	Arr2()
}

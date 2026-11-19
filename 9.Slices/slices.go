package main

import (
	"fmt"
	"slices"
)

//Slice : Dynamic memory allocation
//Most used construct on go
// + useFul methods
func main(){
	//uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums,len(nums))

	slice:=make([]int,0,1)
	slice=append(slice,33)
	slice=append(slice,34)
	fmt.Println(slice)
	
	
	// vals:=[]int{}
	// vals=append(vals,222)
	// fmt.Println(vals)

	//Slice operator
	fmt.Println(slice[0:2])


	//Slices equality

	nums:=[]int{}
	nums2:=[]int{}
	nums2 = append(nums2, 22)
	fmt.Println(slices.Equal(nums,nums2))

	
	
}
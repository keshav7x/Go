package main

import "fmt"

func BasicLoop(){
	for i:=0;i<6;i++{
		fmt.Println(i)
	}
}

func Loop1(){
	i:=0
	for i<5{
		fmt.Println(i)
		i++
	}
}

func Loop2(){
	i:=8
	for i>=0{
		fmt.Println(i)
		if(i==3){
			break
		}
		i--
	}
}


func Loop3(){
	for i:=100;i>=0;i--{
		if(i%2!=0){
			fmt.Println("Odd ",i)
		}
		if(i%2==0){
			continue	
		}
		if(i==55){
			break
		}
	}
}

func outerLoop(){
	outer:
	for i:=0 ;i<3;i++ {
		for j:=0 ;j<3;j++{
			if i==1 && j==1{
				break outer
			}
			fmt.Println(i,j)
		}
	}
}


func countDown(n int){
	i:=n
	for i>0{
		fmt.Println(i)
		i--
	}
}

func evenNum(){
	i:=100
	for i>=0{
		if(i%2==0){
			fmt.Println("Even",i)
			i--
		}else{
			continue
		}
		i--;
	}
}

func main(){
	evenNum()
}

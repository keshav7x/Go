package main

import "fmt"

func add(a int , b int) (sum int){
	sum=a+b
	return
}

func multi(a,b int)(mul int){
	mul=a*b
	return
}

func subs(a,b int)(sub int){
	sub=a-b
	return
}

func divi(a,b int)(div int){
	div=a/b
	return
}

func calc(a,b int) (sum,mul,sub,div int){
	sum=add(a,b)
	mul=multi(a,b)
	sub=subs(a,b)
	div=divi(a,b)
	return 
}

func main(){
	a:=5
	b:=10
	sum,mul,sub,div:=calc(a,b)
	fmt.Println(sub,mul,sum,div)
}
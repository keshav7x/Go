package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	customer  Customer
	createdAt time.Time
}

type Customer struct {
	Id   string
	name string
}

func newCustomer(Id, name string) *Customer {
	customer := Customer{
		Id:   Id,
		name: name,
	}
	return &customer
}

func newOrder(id string,amount float32,status string,customer Customer)(*Order){
	order:=Order{
		id:id,
		amount: amount,
		status: status,
		customer: customer,
		createdAt: time.Now(),
	}
	return &order
}

func main(){
	c1:=newCustomer("001","Keshav")
	o1:=newOrder("001",555,"paid",*c1)
	fmt.Println("customer ",*c1 ,"ordered order",*o1)
}
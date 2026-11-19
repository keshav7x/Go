package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *Order) changeStatus(status string){
	o.status=status
}

func (o *Order) changeAmount(amount float32){
	o.amount=amount
}


func newOrder(id string,amount float32,status string,createdAt time.Time)(*Order){
	newOrder:=Order{
		id:id,
		amount: amount,
		status: status,
		createdAt: createdAt,
	}
	return &newOrder
}

func main() {
	order1:=newOrder("001",545,"paid",time.Now())
	fmt.Println(*order1)
	language:=struct{
		name string
		isGood bool
	}{"golang",true}
	fmt.Println(language)
}

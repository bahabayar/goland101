package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	customerId      int
	customerName    string
	customerSurname string
}
type product struct {
	productId    int
	productName  string
	productPrice float64
}
type order struct {
	orderId    int
	customerId int
	productId  int
	amount     int
}

func giveInfo(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Variable type = ", reflect.TypeOf(i))
	fmt.Println("Variable kind = ", t.Kind())
	fmt.Println("Variable type = ", reflect.ValueOf(i))
}

func main() {
	c := customer{customerId: 123, customerName: "Baha", customerSurname: "Bayar"}
	p := product{productId: 23, productName: "Dalin", productPrice: 12.5}
	o := order{orderId: 1, customerId: 123, productId: 23, amount: 2}

	giveInfo(c)
	giveInfo(p)
	giveInfo(o)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	order := &Order{
		Id: 1001,
		Customer: Customer{
			FirstName: "Alex",
			LastName:  "John",
			Email:     "alex@email.com",
			Phone:     "732-757-2923",
			Addresses: []Address{
				Address{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: true,
				},
				Address{
					Street: "49 Stevenson Street",
					City:   "San Francisco",
					State:  "CA",
					Zip:    "94105",
				},
			},
		},
		Status:   "Placed",
		PlacedOn: time.Date(2016, time.April, 10, 0, 0, 0, 0, time.UTC),
		OrderItems: []OrderItem{
			OrderItem{
				Product: Product{
					Code:        "knd100",
					Name:        "Kindle Voyage",
					Description: "Kindle Voyage Wifi, 6 High-Resolution Display",
					UnitPrice:   220,
				},
				Quantity: 1,
			},
			OrderItem{
				Product: Product{
					Code:        "fint101",
					Name:        "Kindle Case",
					Description: "Fintie Kindle Voyage SmartShell Case",
					UnitPrice:   10,
				},
				Quantity: 2,
			},
		},
	}
	fmt.Println(order.ToString())
	// Change Order status
	order.ChangeStatus("Processing")
	fmt.Println("\n")
	fmt.Println(order.ToString())
}

type Address struct {
	Street, City, State, Zip string
	IsShippingAddress        bool
}
type Customer struct {
	FirstName, LastName, Email, Phone string
	Addresses                         []Address
}

func (c Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email:%s", c.FirstName, c.LastName, c.Email)
}

func (c Customer) ShippingAddress() string {
	for _, v := range c.Addresses {
		if v.IsShippingAddress == true {
			return fmt.Sprintf("%s, %s, %s, Zip - %s", v.Street, v.City, v.State, v.Zip)
		}
	}
	return ""
}

type Order struct {
	Id int
	Customer
	PlacedOn   time.Time
	Status     string
	OrderItems []OrderItem
}

func (o *Order) GrandTotal() float64 {
	var total float64
	for _, v := range o.OrderItems {
		total += v.Total()
	}
	return total
}
func (o *Order) ToString() string {
	var orderStr string
	orderStr = fmt.Sprintf("Order#:%d, OrderDate:%s, Status:%s, Grand Total:%f\n", o.Id,
		o.PlacedOn, o.Status, o.GrandTotal())
	orderStr += o.Customer.ToString()
	orderStr += fmt.Sprintf("\nOrder Items:")
	for _, v := range o.OrderItems {
		orderStr += fmt.Sprintf("\n")
		orderStr += v.ToString()
	}
	orderStr += fmt.Sprintf("\nShipping Address:")
	orderStr += o.Customer.ShippingAddress()
	return orderStr
}
func (o *Order) ChangeStatus(newStatus string) {
	o.Status = newStatus
}

type OrderItem struct {
	Product
	Quantity int
}

func (item OrderItem) Total() float64 {
	return float64(item.Quantity) * item.Product.UnitPrice
}
func (item OrderItem) ToString() string {
	itemStr := fmt.Sprintf("Code:%s, Product:%s -- %s, UnitPrice:%f, Quantity:%d, Total:%f",
		item.Product.Code, item.Product.Name, item.Product.Description, item.Product.
			UnitPrice, item.Quantity, item.Total())
	return itemStr
}

type Product struct {
	Code, Name, Description string
	UnitPrice               float64
}

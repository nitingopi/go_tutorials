package main

import (
	"assignments/assignment2/domain"
	"assignments/assignment2/mapstore"
	"fmt"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (c CustomerController) Add(customer domain.Customer) {
	err := c.store.Create(customer)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("New Customer has been created")
}
func (c CustomerController) GetAll() {
	customers, err := c.store.GetAll()
	if err != nil {
		fmt.Println("Error : ", err)

	} else {
		fmt.Println(customers)
	}

}
func (c CustomerController) GetById(id string) (domain.Customer, error) {
	customer, err := c.store.GetById(id)
	if err != nil {
		fmt.Println("Error : ", err)

	}
	return customer, err

}
func (c CustomerController) Delete(id string) error {
	err := c.store.Delete(id)
	if err != nil {
		fmt.Println("Error : ", err)

	}
	return err

}
func (c CustomerController) Update(id string, customer domain.Customer) error {
	err := c.store.Update(id, customer)
	if err != nil {
		fmt.Println("Error : ", err)

	}
	return err

}

func main() {
	fmt.Println("Hello world")
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:    "101",
		Name:  "Nitin Gopi",
		Email: "nitin@email.com",
	}
	controller.Add(customer)
	controller.GetAll()
	customer2 := domain.Customer{
		ID:    "101",
		Name:  "Gaurav Sharm",
		Email: "gaurav@email.com",
	}
	controller.Update("101", customer2)
	controller.GetAll()
	controller.GetById("101")
	customer3 := domain.Customer{
		ID:    "102",
		Name:  "Nitin Gopi",
		Email: "nitin@email.com",
	}
	controller.Add(customer3)
	controller.GetAll()
}

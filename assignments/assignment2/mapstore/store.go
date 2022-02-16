package mapstore

import (
	"assignments/assignment2/domain"
	"fmt"
)

type MapStore struct {
	// An in-memory store with a map
	// use Customer.ID as key of the map
	store map[string]domain.Customer
}

//Factory method gives a new instance of MapStore
//This is for caller packages to create MapStore instances
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (store MapStore) Create(c domain.Customer) error {
	if _, ok := store.store[c.ID]; ok {
		err := fmt.Errorf("the customer already exists")
		return err
	}
	store.store[c.ID] = c
	return nil

}

func (store MapStore) Delete(id string) error {
	if _, ok := store.store[id]; ok {
		delete(store.store, id)
		return nil
	}
	err := fmt.Errorf("the customer does not exist")
	return err
}

func (store MapStore) Update(id string, c domain.Customer) error {
	if _, ok := store.store[c.ID]; ok {
		store.store[c.ID] = c
		return nil
	}
	err := fmt.Errorf("the customer does not exist")
	return err

}

func (store MapStore) GetById(id string) (domain.Customer, error) {
	var customer domain.Customer
	if _, ok := store.store[id]; ok {
		return store.store[id], nil

	}
	err := fmt.Errorf("could not find the customer")
	return customer, err

}

func (store MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer

	total := len(store.store)
	fmt.Println("total is ", total)
	if total > 0 {
		customers := make([]domain.Customer, total)
		i := 0
		for k, _ := range store.store {
			customers[i] = store.store[k]
			i++
		}
		return customers, nil
	}

	err := fmt.Errorf("the map store is empty")
	return customers, err

}

package singletonpackage

import (
	"errors"
	"log"
	"sync"
)

// Ensures it is only ran once
var so *sync.Once

// Inits initializes the sync once
func init() {
	so = &sync.Once{}
}

var ErrInstanceExists = errors.New("instance exists")

type instance struct {
	name string
}

var inst *instance

func (i *instance) About() {
	log.Println("Hi im instance:", i.name)
}

func (i *instance) AddressInfo() {
	log.Println("Im in:", &inst)
}

func GetInstance() *instance {
	if inst == nil {
		log.Println("[ERR] Instance not yet initialized!")
		return nil
	}

	return inst
}

func CloseInstance() {
	if inst != nil {
		inst = nil
		log.Println("Instance Deleted")
	}
}

func Initialize(name string) error {
	if inst != nil {
		log.Println("[ERR] Instance already exist")
		return ErrInstanceExists
	}
	// Wrapped it in sync once incase the function is called parallel
	// Eventhough in golang concurrency scheduling that shouldn't happen
	so.Do(func() {

	})
	inst = &instance{
		name: name,
	}
	return nil
}

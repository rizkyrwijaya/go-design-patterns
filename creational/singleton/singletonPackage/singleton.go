package singletonpackage

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotInitialized = errors.New("instance not initialized")
var ErrInstanceExists = errors.New("instance already initialized")

type instance struct {
	name string
}

var inst *instance

func (i *instance) About() string {
	return i.name
}

func (i *instance) AddressInfo() string {
	return fmt.Sprintf("%p", i)
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

	inst = &instance{
		name: name,
	}

	return nil
}

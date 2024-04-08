package methodFactory

import (
	"errors"
	"log"
)

// Global Variable for Register
var factory map[string]Initializer

type Initializer interface {
	Init() Runner
}

type Runner interface {
	Run()
}

func init() {
	log.Println("Making sure its run once")
	factory = make(map[string]Initializer)
}

func RegisterMethod(name string, init Initializer) error {
	if _, ok := factory[name]; !ok {
		factory[name] = init
		return nil
	}

	return errors.New("already registered")
}

func CreateRunnerFromFactory(name string) Runner {
	log.Println(len(factory))
	if _, ok := factory[name]; ok {
		return factory[name].Init()
	}

	return nil
}

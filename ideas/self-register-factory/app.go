package main

import (
	"log"

	"github.com/rizkyrwijaya/go-design-patterns/ideas/self-register-factory/methodFactory"
	_ "github.com/rizkyrwijaya/go-design-patterns/ideas/self-register-factory/methodFactory/methodA"
	_ "github.com/rizkyrwijaya/go-design-patterns/ideas/self-register-factory/methodFactory/methodB"

	"github.com/rizkyrwijaya/go-design-patterns/ideas/self-register-factory/bugtest"
)

func main() {
	// Getting Method Factory
	a := methodFactory.CreateRunnerFromFactory("methoda")
	log.Println(a)
	a.Run()

	b := methodFactory.CreateRunnerFromFactory("methodb")
	log.Println(b)
	b.Run()

	bugtest.Test()
}

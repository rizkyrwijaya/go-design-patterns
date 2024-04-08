package methoda

import (
	"log"

	"github.com/rizkyrwijaya/go-design-patterns/ideas/self-register-factory/methodFactory"
)

type methoda struct {
	// Maybe can add a config later
	name string
}

func (m *methoda) Init() methodFactory.Runner {
	return m
}

func (m *methoda) Run() {
	log.Println("Running method ", m.name)
}

func init() {
	initializer := &methoda{
		name: "Method B",
	}
	methodFactory.RegisterMethod("methodb", initializer)
}

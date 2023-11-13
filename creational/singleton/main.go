package main

import (
	"log"
	"strings"

	singletonpackage "github.com/rizkyrwijaya/go-design-patterns/creational/singleton/singletonPackage"
)

func main() {
	// Lets check the current singleton instance!
	testInstance(1)

	// Lets try initializing it once
	log.Println()
	log.Println("Initializing Instance")
	err := singletonpackage.Initialize("instance A")
	if err != nil {
		log.Println("failed to initialize instance:", err)
		return
	}
	log.Println()
	testInstance(2)
	log.Println()
	// Lets try changing the instance
	// First we delete the instnace
	singletonpackage.CloseInstance()
	log.Println()
	err = singletonpackage.Initialize("instance B")
	if err != nil {
		log.Println("failed to initialize instance:", err)
		return
	}
	testInstance(3)

}

// Dont mind this function, its just for logging
func testInstance(x ...int) {
	if len(x) != 0 {
		log.Println(strings.Repeat("=", 8), x[0], strings.Repeat("=", 9))
	} else {
		log.Println(strings.Repeat("=", 20))
	}
	i := singletonpackage.GetInstance()
	if i == nil {
		log.Println("Instance is NIL, do not proceed!")

		if len(x) != 0 {
			log.Println(strings.Repeat("=", 8), x[0], strings.Repeat("=", 9))
		} else {
			log.Println(strings.Repeat("=", 20))
		}
		return
	}

	log.Println("Instance has been created! Printing Info")
	i.About()
	i.AddressInfo()

	if len(x) != 0 {
		log.Println(strings.Repeat("=", 8), x[0], strings.Repeat("=", 9))
	} else {
		log.Println(strings.Repeat("=", 20))
	}
}

package main

import (
	"fmt"

	"github.com/aechiara/gointerfaces/animals"
	"github.com/aechiara/gointerfaces/circus"
	"github.com/aechiara/gointerfaces/people"
)

// Go interfaces encourages one to be lazy, and this is a good thing.
// Instead of writing types to fulfil interfaces, write interfaces to fulfil usage requirements.

func main() {
	dog := animals.Dog{}
	fmt.Println(circus.Perform(dog))

	man := people.Man{}
	fmt.Println(circus.Perform(man))
}

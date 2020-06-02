package main

import (
	"fmt"
)

type Character interface {
	getActions() (string, string)
}

type Pahom struct {
	tellStory string
	poop string
}

type Bratok struct {
	becameBeast string
	useFork string
}

func (p Pahom) getActions () (string, string) {
	return p.tellStory, p.poop
}

func (b Bratok) getActions () (string, string) {
	return b.becameBeast, b.useFork
}

func doPolymorfStuff (c Character) {
	switch c.(type) {
	case Pahom: 
		tellStory, poop := c.(Pahom).getActions()
		fmt.Printf("%v . 3 Hours later... %v", tellStory, poop)
	case Bratok:
		action1, action2 := c.(Bratok).getActions()
		fmt.Println(action1, action2)
	}
}

func main() {
	mad := Pahom{
		tellStory: "Do you wanna story about poop?",
		poop: "uuf uuf uuuuf uf uf",
	}
	bidlo := Bratok{
		becameBeast: "AAAAAAARRRRR",
		useFork: "How I suppose to clean toilet with fork?",
	}
	doPolymorfStuff(mad)
	doPolymorfStuff(bidlo)
}

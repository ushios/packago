package animals

import "fmt"

type Zoo struct {
	Name string
}

func NewZoo(name string) Zoo {
	zoo := Zoo{
		Name: name,
	}

	return zoo
}

func (z Zoo) Check(a Animal) {
	switch {
	default:
		fmt.Println("...")
	case a.Type() == "human":
		fmt.Println("You are human. Please show me your ticket.")
	}
}

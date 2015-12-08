package animals

type Zoo struct {
	Name string
}

func NewZoo(name string) Zoo {
	zoo := Zoo{
		Name: name,
	}

	return zoo
}

func (z Zoo) Check(animal Animal) string {
	switch {
	default:
		return "..."
	case animal.Type() == "human":
		return "Ticket!"
	}
}

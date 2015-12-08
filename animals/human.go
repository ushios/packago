package animals

type Human struct {
	FirstName string
	LastName  string
}

func (h Human) FullName() (fullname string) {
	fullname = h.FirstName + " " + h.LastName
	return
}

func NewHuman(firstName string, lastName string) Human {
	human := Human{
		FirstName: firstName,
		LastName:  lastName,
	}

	return human
}

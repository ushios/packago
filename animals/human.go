package animals

type Human struct {
	FirstName string
	LastName  string
}

func (h Human) FullName() (fullname string) {
	fullname = h.FirstName + " " + h.LastName
	return
}

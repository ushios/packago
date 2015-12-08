package animals

type Dog struct {
}

func (c Dog) Type() string {
	return "dog"
}

func NewDog() Dog {
	return Dog{}
}

package animals

type Cat struct {
}

func (c Cat) Type() string {
	return "cat"
}

func NewCat() Cat {
	return Cat{}
}

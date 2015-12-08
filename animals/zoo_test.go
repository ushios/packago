package animals

import "testing"

func TestZoo(t *testing.T) {
	uenoZoo := NewZoo("上野動物園")

	human := NewHuman("野比", "のび太")
	cat := NewCat()
	dog := NewDog()

	uenoZoo.Check(human)
	uenoZoo.Check(cat)
	uenoZoo.Check(dog)
}

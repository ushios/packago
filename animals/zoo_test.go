package animals

import (
	"fmt"
	"testing"
)

func TestZoo(t *testing.T) {
	uenoZoo := NewZoo("上野動物園")

	human := NewHuman("野比", "のび太")
	cat := NewCat()
	dog := NewDog()

	fmt.Println(uenoZoo.Check(human))
	fmt.Println(uenoZoo.Check(cat))
	fmt.Println(uenoZoo.Check(dog))
}

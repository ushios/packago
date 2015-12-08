package animals

import "testing"

func TestHuman(t *testing.T) {
	human := Human{
		FirstName: "Nobi",
		LastName:  "Nobita",
	}

	fullName := human.FullName()

	if fullName != "Nobi Nobita" {
		t.Errorf("human fullname want (%s)", "Nobi Nobita")
	}
}

func TestNewHuman(t *testing.T) {
	human := NewHuman("Nobi", "Nobita")

	fullName := human.FullName()

	if fullName != "Nobi Nobita" {
		t.Errorf("human fullname want (%s)", "Nobi Nobita")
	}
}

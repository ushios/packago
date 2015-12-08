package animals

import "testing"

func TestHuman(t *testing.T) {
	human := Human{
		FirstName: "Nobi",
		LastName:  "Nobita",
	}

	if human.FullName() != "Nobi Nobita" {
		t.Errorf("human fullname want (%s)", "Nobi Nobita")
	}
}

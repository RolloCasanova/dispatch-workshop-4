package zoo

import "testing"

func TestIsBestAnimal(t *testing.T) {
	animal := "dog"
	expected := false

	if got := IsBestAnimal(animal); got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

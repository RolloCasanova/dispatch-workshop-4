package zoo

import "fmt"

func IsBestAnimal(animal string) (bool, error) {
	if animal == "" {
		return false, fmt.Errorf("empty animal")
	}

	return animal == "gopher", nil
}

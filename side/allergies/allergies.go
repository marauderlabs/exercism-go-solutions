package allergies

import "errors"

var allergies = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns a list of allergies that the person with score s has.
func Allergies(s uint) []string {
	var b uint
	var r []string
	for i := 0; i < len(allergies); i++ {
		if (s>>b)&1 == 1 {
			r = append(r, allergies[b])
		}
		b++
	}
	return r
}

func bitOfAllergy(a string) (uint, error) {
	for i, e := range allergies {
		if e == a {
			return (1 << uint(i)), nil
		}
	}
	return 0, errors.New("not found")
}

// AllergicTo returns true if the person is allergic to the given substance
func AllergicTo(score uint, substance string) bool {
	i, _ := bitOfAllergy(substance)
	if score&i != 0 {
		return true
	}
	return false
}

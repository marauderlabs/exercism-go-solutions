package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot shall be the name string generated for robots
type Robot struct {
	name string
}

var generated = make(map[string]bool)

var maxPossibleNames = 26 * 26 * 10 * 10 * 10

// Name returns a new name for a robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(generated) == maxPossibleNames {
			return "", errors.New("Reached max allowed possible names")
		}

		exists := true
		for exists {
			r.name = randNameString()
			_, exists = generated[r.name]
		}
		generated[r.name] = true
	}
	return r.name, nil
}

// Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func randNameString() string {
	return fmt.Sprintf("%s%s%d%d%d", string(randChar()), string(randChar()),
		randInt(), randInt(), randInt())
}

func randChar() rune {
	return rune('A' + rand.Intn(26))
}

func randInt() int {
	return rand.Intn(10)
}

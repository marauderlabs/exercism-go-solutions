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
	if r.name != "" {
		return r.name, nil
	}

	if len(generated) == maxPossibleNames {
		return "", errors.New("Reached max allowed possible names")
	}

	r.name = randNameString()
	for generated[r.name] {
		r.name = randNameString()
	}
	generated[r.name] = true
	return r.name, nil
}

// Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func randNameString() string {
	return fmt.Sprintf("%s%s%03d", randString(), randString(), rand.Intn(1000))
}

func randString() string {
	return string('A' + rand.Intn(26))
}

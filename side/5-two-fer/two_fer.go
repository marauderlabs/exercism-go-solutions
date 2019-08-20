// Package twofer solves the two-fer exercise at exercism.io
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns two-fer string for the given name
func ShareWith(name string) string {

	if strings.TrimSpace(name) == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

// Package greeting builds the messages printed by the testproject command.
package greeting

import "fmt"

// Message returns the greeting for name, defaulting to "world" when name is empty.
func Message(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("Hello, %s!", name)
}

// Command testproject is a minimal program used to exercise CI security checks.
package main

import (
	"fmt"
	"os"

	"github.com/lee5i3/testproject/internal/greeting"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeting.Message(name))
}

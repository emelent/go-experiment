package misc

import "fmt"

// Greet returns a string greeting a person by the name provided
func Greet(name string) string {
	return fmt.Sprintf("Hello there, %s.", name)
}

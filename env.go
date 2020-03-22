// To set a key/value pair, use `os.Setenv`.
//
// To get a value for a key, use `os.Getenv`.
// This will return an empty string if the key isn’t present in the environment.
//
// Use `os.Environ` to list all key/value pairs in the environment.
// This returns a slice of strings in the form `KEY=value`. You can `strings.SplitN` them to get the key and value.
//

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], "=", pair[1])
	}
}

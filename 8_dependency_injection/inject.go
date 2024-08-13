package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

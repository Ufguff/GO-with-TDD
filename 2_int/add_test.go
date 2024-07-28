package int

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got %d, expected %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}

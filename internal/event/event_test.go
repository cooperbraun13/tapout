package event

import (
	"fmt"
	"testing"
)

func TestLoadAll(t *testing.T) {
	events, err := LoadAll("../../events") // Path relative to test file
	if err != nil {
		t.Fatalf("LoadAll failed: %v", err)
	}

	fmt.Printf("%+v\n", events) // Prints the struct with field names
}

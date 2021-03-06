package textconv

import (
	"fmt"
	"strings"
	"testing"
)

func TestParallelTextConverter(t *testing.T) {
	// Parallel upper-case converter
	conv := NewParallelTextConverter(StringMapFunc(strings.ToUpper), " ")

	got, err := conv.ConvertText("cześć, my name is dr. greenthumb!")
	if err != nil {
		t.Fatal(err)
	}

	if got != "CZEŚĆ, MY NAME IS DR. GREENTHUMB!" {
		t.Error(got)
	}
}

func ExampleStringMapFunc() {
	c := StringMapFunc(strings.ToUpper)
	txt, _ := c.ConvertText("what's up?")

	fmt.Println(txt)

	// Output:
	// WHAT'S UP?
}

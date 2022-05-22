package controller

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Run("testing 2 + 2", func(t *testing.T) {
		got := Add(2, 2)
		wait := 4

		if got != wait {
			t.Errorf("wait %d but got %d", wait, got)
		}
	})
}

func ExampleAdd() {
	result := Add(5, 6)
	fmt.Printf("result %d", result)

	//Output 11
}

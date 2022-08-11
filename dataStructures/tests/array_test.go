package tests

import (
	"testing"

	"github.com/theghostmac/golang-ds-algo/dataStructures"
)

func TestGetFunction(t *testing.T) {
	arr := dataStructures.Array{}

	tests := []struct{
		expected interface{}
	}{
		{1},
		{"1"},
	}
	arr.Push(1)
	arr.Push("1")

	for i, tst := range tests {
		value, err := arr.Get(i)
			if err != nil {
				t.Fatalf("Index %d is out of range. Array length is %d", i, arr.Length())
			}
		}

		if value != tst.expected {
			t.Fatalf("wrong value. Expected: %v, got %v", tst.expected, value)
		}
	}
}
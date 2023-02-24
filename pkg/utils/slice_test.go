package utils

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	t.Run("it removes an element at given index from int slice", func(t *testing.T) {
		sampleSlice := []int{0, 1, 2, 3, 4, 5}
		index := 2
		want := []int{0, 1, 3, 4, 5}
		got := RemoveAt(sampleSlice, index)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

package array

import "testing"

func TestSlice(t *testing.T) {
	arr := NewArray()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arr.Push(4)
	arr.Push(5)

	tests := []struct {
		start, end int
		expected   []interface{}
	}{
		{0, 2, []interface{}{1, 2}},
	}

	for _, tt := range tests {
		slice := arr.Slice(tt.start, tt.end)
		if slice.size != len(tt.expected) {
			t.Errorf("Slice(%d, %d) got size %d, want %d", tt.start, tt.end, slice.size, len(tt.expected))
		}
		for i, v := range tt.expected {
			if slice.elements[i] != v {
				t.Errorf("Slice(%d, %d) got %v at index %d, want %v", tt.start, tt.end, slice.elements[i], i, v)
			}
		}
	}
}

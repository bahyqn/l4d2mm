package internal

import "testing"

func TestAreSlicesEqualUnordered(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want bool
	}{
		{
			name: "exact same order and elements",
			a:    []string{"apple", "banana", "orange"},
			b:    []string{"apple", "banana", "orange"},
			want: true,
		},
		{
			name: "same elements different order",
			a:    []string{"apple", "banana", "orange"},
			b:    []string{"orange", "apple", "banana"},
			want: true,
		},
		{
			name: "duplicate elements with matching frequencies",
			a:    []string{"apple", "apple", "banana"},
			b:    []string{"banana", "apple", "apple"},
			want: true,
		},
		{
			name: "different lengths",
			a:    []string{"apple", "banana"},
			b:    []string{"apple", "banana", "orange"},
			want: false,
		},
		{
			name: "same length different elements",
			a:    []string{"apple", "banana", "orange"},
			b:    []string{"apple", "banana", "grape"},
			want: false,
		},
		{
			name: "same length different frequencies",
			a:    []string{"apple", "apple", "banana"},
			b:    []string{"apple", "banana", "banana"},
			want: false,
		},
		{
			name: "both empty slices",
			a:    []string{},
			b:    []string{},
			want: true,
		},
		{
			name: "nil slice and empty slice",
			a:    nil,
			b:    []string{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AreSlicesEqualUnordered(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AreSlicesEqualUnordered() = %v, want %v (input a: %v, b: %v)", got, tt.want, tt.a, tt.b)
			}
		})
	}
}

func TestAreSlicesEqualUnordered_Int(t *testing.T) {
	s1 := []int{1, 2, 3, 2}
	s2 := []int{2, 3, 1, 2}

	if !AreSlicesEqualUnordered(s1, s2) {
		t.Errorf("Expected int slices to match, but got false")
	}
}

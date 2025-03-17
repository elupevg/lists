package lists

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSinglyLinked(t *testing.T) {
	testCases := []struct {
		name string
		want []int
	}{
		{
			name: "Empty",
			want: []int{},
		},
		{
			name: "OneElement",
			want: []int{3},
		},
		{
			name: "ManyElements",
			want: []int{5, 0, -1, 100},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := NewSinglyLinked(tc.want)
			got := head.Slice()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

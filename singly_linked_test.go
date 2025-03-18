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

func TestAppendNode(t *testing.T) {
	testCases := []struct {
		name string
		node *NodeSL[int]
		want []int
	}{
		{
			name: "One",
			node: &NodeSL[int]{Value: 17},
			want: []int{1, 2, 3, 17},
		},
		{
			name: "JoinTwoLists",
			node: &NodeSL[int]{Value: 17, Next: &NodeSL[int]{Value: 18}},
			want: []int{1, 2, 3, 17, 18},
		},
		{
			name: "Nil",
			node: nil,
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := NewSinglyLinked([]int{1, 2, 3})
			head.AppendNode(tc.node)
			got := head.Slice()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

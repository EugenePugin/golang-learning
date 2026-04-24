package lc10

import (
	"testing"
)

func Test_closestTarget(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		words      []string
		target     string
		startIndex int
		want       int
	}{
		{
			name:       "test",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "hello",
			startIndex: 0,
			want:       0,
		},
		{
			name:       "test",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "heldddlo",
			startIndex: 0,
			want:       -1,
		},
		{
			name:       "test",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "i",
			startIndex: 0,
			want:       1,
		},

		{
			name:       "test",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "leetcode",
			startIndex: 0,
			want:       2,
		},
		{
			name:       "test",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "hello",
			startIndex: 1,
			want:       1,
		},

		{
			name:       "test",
			words:      []string{"a", "b", "leetcode"},
			target:     "leetcode",
			startIndex: 0,
			want:       1,
		}, {
			name:       "test",
			words:      []string{"i", "eat", "leetcode"},
			target:     "ate",
			startIndex: 0,
			want:       -1,
		},
		{
			name:       "test",
			words:      []string{"apple", "banana", "computer", "elephant", "guitar", "hospital", "island", "jacket", "kangaroo", "lemon", "mountain", "notebook", "ocean", "pencil", "rabbit", "sunshine", "tiger", "umbrella", "violin", "waterfall", "xylophone", "yacht", "zebra", "adventure"},
			target:     "elephant",
			startIndex: 1,
			want:       2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closestTarget(tt.words, tt.target, tt.startIndex)
			// TODO: update the condition below to compare got with tt.want.
			// fmt.Println(got, tt.want)
			if got != tt.want {
				t.Errorf("closestTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

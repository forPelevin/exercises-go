package solution_test

import (
	"strconv"
	"testing"

	solution "exercises-go/modules/10-basics/30-funcs"
)

func TestIntToString(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 0,
			want:  "0",
		},
		{
			input: -42,
			want:  "-42",
		},
		{
			input: 100500,
			want:  "100500",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solution.IntToString(tt.input); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

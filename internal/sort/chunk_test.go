package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk_ToArray(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "EmptyChunk",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "NonEmptyChunk",
			input: []string{"b", "a", "c", "c1", "b1", "a1"},
			want:  []string{"a", "a1", "b", "b1", "c", "c1"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ch := NewChunk(len(tt.input))
			for _, in := range tt.input {
				ch.Add(in)
			}

			got := ch.ToArray()
			assert.Equal(t, tt.want, got)
		})
	}
}

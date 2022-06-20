package goladok3

import (
	"testing"

	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestGenderString(t *testing.T) {
	tts := []struct {
		name string
		have *ladoktypes.Student
		want string
	}{
		{
			name: "female",
			have: &ladoktypes.Student{
				KonID: 1,
			},
			want: "female",
		},
		{
			name: "male",
			have: &ladoktypes.Student{
				KonID: 2,
			},
			want: "male",
		},
		{
			name: "n/a",
			have: &ladoktypes.Student{
				KonID: 10,
			},
			want: "n/a",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.have.GenderString()
			assert.Equal(t, tt.want, got)
		})
	}
}

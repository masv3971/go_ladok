package goladok3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenderString(t *testing.T) {
	tts := []struct {
		name string
		have *GetStudentReply
		want string
	}{
		{
			name: "1",
			have: &GetStudentReply{
				KonID: 1,
			},
			want: "female",
		},
		{
			name: "2",
			have: &GetStudentReply{
				KonID: 2,
			},
			want: "male",
		},
		{
			name: "10",
			have: &GetStudentReply{
				KonID: 10,
			},
			want: "n/a",
		},
	}

	for _, tt := range tts {
		got := tt.have.GenderString()
		assert.Equal(t, tt.want, got)
	}
}

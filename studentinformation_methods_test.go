package ladok3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenderString(t *testing.T) {
	tts := []struct {
		name string
		have *StudentReply
		want string
	}{
		{
			name: "1",
			have: &StudentReply{
				KonID: 1,
			},
			want: "female",
		},
		{
			name: "2",
			have: &StudentReply{
				KonID: 2,
			},
			want: "male",
		},
		{
			name: "10",
			have: &StudentReply{
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

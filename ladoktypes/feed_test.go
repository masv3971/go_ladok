package ladoktypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedIDTrim(t *testing.T) {
	tts := []struct {
		name string
		have FeedID
		want FeedID
	}{
		{
			name: "OK",
			have: "urn:id:4856",
			want: "4856",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.have.trim())
		})
	}
}

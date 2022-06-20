package ladoktypes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorToString(t *testing.T) {
	var (
		mockPermissionErrors = PermissionErrors{
			{
				Msg:                 "testMsg1",
				MissingPermissionID: 7771,
				PermissionLevel:     "testPermissionLevel1",
			},
			{
				Msg:                 "testMsg2",
				MissingPermissionID: 7772,
				PermissionLevel:     "testPermissionLevel2",
			},
		}
		mockPermissionErrorsEmpty = PermissionErrors{}

		mockLadokError = LadokError{
			Detaljkod:       "testDetaljkod",
			DetaljkodText:   "testDetaljkodText",
			FelUID:          "testFelUID",
			Felgrupp:        "testFelgrupp",
			FelgruppText:    "testFelgruppText",
			Felkategori:     "testFelkategori",
			FelkategoriText: "testFelkategoriText",
			Meddelande:      "testMeddelande",
			Link:            []interface{}{},
		}

		mockInternalError = errors.New("testInternalError")
	)
	tts := []struct {
		name string
		have interface{}
		want string
	}{
		{
			name: "Internal error",
			have: mockInternalError,
			want: "testInternalError",
		},
		{
			name: "Permission error",
			have: mockPermissionErrors,
			want: "testMsg1,7771,testPermissionLevel1\ntestMsg2,7772,testPermissionLevel2\n",
		},
		{
			name: "Permission error, empty",
			have: mockPermissionErrorsEmpty,
			want: "",
		},
		{
			name: "Ladok error",
			have: mockLadokError,
			want: "felUID: \"testFelUID\", detaljkod_text: \"testDetaljkodText\"",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.have.(type) {
			case PermissionErrors:
				e := tt.have.(PermissionErrors)
				assert.Equal(t, tt.want, e.Error())
			case LadokError:
				e := tt.have.(LadokError)
				assert.Equal(t, tt.want, e.Error())
			default:
				e := tt.have.(error)
				assert.Equal(t, tt.want, e.Error())
			}
		})
	}
}

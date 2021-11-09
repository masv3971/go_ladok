package goladok3

import (
	"testing"

	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestErrorsError(t *testing.T) {
	var mockInternalError = ladoktypes.InternalError{
		Msg:           "testMsg",
		Type:          "testType",
		Func:          "testFunc",
		PreviousError: "testPreviousError",
	}
	var mockLadokError = &ladoktypes.LadokError{
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
	tts := []struct {
		name string
		have *Errors
		want string
	}{
		{
			name: "OK, ladok & internal",
			have: &Errors{Internal: []ladoktypes.InternalError{mockInternalError, mockInternalError}, Ladok: mockLadokError},
			want: "internal error: [{testMsg testType testFunc testPreviousError} {testMsg testType testFunc testPreviousError}], ladok error: &{testDetaljkod testDetaljkodText testFelUID testFelgrupp testFelgruppText testFelkategori testFelkategoriText testMeddelande []}",
		},
		{
			name: "OK, only internal",
			have: &Errors{Internal: []ladoktypes.InternalError{mockInternalError}},
			want: "internal error: [{testMsg testType testFunc testPreviousError}]",
		},
		{
			name: "Only ladok",
			have: &Errors{Ladok: mockLadokError},
			want: "ladok error: &{testDetaljkod testDetaljkodText testFelUID testFelgrupp testFelgruppText testFelkategori testFelkategoriText testMeddelande []}",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.have.Error())
		})
	}
}

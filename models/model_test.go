package models

import (
	"encoding/json"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	ses := Session{}

	fz := fuzz.New()

	for i := 0; i < 10000; i++ {
		fz.Fuzz(&ses)

		_, err := json.Marshal(&ses)
		assert.NoError(t, err)
	}
}

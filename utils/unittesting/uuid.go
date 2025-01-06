package unittesting

import (
	"bytes"

	"github.com/google/uuid"
)

/*
See:
https://www.reddit.com/r/golang/comments/vdl5xx/testing_uuid_how_to_access_same_uuid_as_created/
*/
func FixNextUuid() {
	reader := bytes.NewReader([]byte("1111111111111111"))
	uuid.SetRand(reader)
	uuid.SetClockSequence(1)
}

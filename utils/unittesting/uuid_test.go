package unittesting

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFixNextUuid(t *testing.T) {
	Convey("When generating multiple UUIDs", t, func() {
		FixNextUuid()
		uuid1 := uuid.New().String()

		FixNextUuid()
		uuid2 := uuid.New().String()

		Convey("They should all be the same value", func() {
			So(uuid1, ShouldEqual, uuid2)
		})
	})
}

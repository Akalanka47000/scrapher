package utils_test

import (
	"scrapher/src/utils"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	t.Parallel()
	Convey("Protect", t, func() {
		Convey("should catch and log panic", func() {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go utils.Protect(func() {
				defer wg.Done()
				panic("test")
			})
			wg.Wait()
			SoMsg("this line should execute", 1, ShouldEqual, 1)
		})
		Convey("should catch panic and pass it the error handler", func() {
			wg := sync.WaitGroup{}
			caught := false
			wg.Add(1)
			utils.Protect(func() {
				defer wg.Done()
				panic("test")
			}, func(err any) {
				caught = true
			})
			wg.Wait()
			So(caught, ShouldBeTrue)
			SoMsg("this line should execute", 1, ShouldEqual, 1)
		})
	})
}

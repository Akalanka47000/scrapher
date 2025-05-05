package system_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"scrapher/src/app"
	"scrapher/tests"
	"testing"
)

func TestSystemMetricsHandler(t *testing.T) {
	t.Parallel()

	tests.Setup()

	app := app.New()

	Convey("should be restricted", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "/system/metrics", nil)
		res, err := app.Test(req, -1)

		So(err, ShouldBeNil)

		defer res.Body.Close()

		So(res.StatusCode, ShouldEqual, http.StatusForbidden)
	})
}

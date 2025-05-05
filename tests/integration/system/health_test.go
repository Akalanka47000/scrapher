package system_test

import (
	"fmt"
	"io"
	"net/http"
	"scrapher/src/app"

	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
	"scrapher/tests"
	"testing"
)

func TestSystemHealthHandler(t *testing.T) {
	t.Parallel()

	tests.Setup()

	app := app.New()

	Convey("returns ok", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "/system/health", nil)
		res, err := app.Test(req, -1)

		So(err, ShouldBeNil)

		defer res.Body.Close()

		So(res.StatusCode, ShouldEqual, http.StatusOK)

		body := string(lo.Ok(io.ReadAll(res.Body)))

		So(body, ShouldEqual, fmt.Sprintf("%s - OK", app.Config().AppName))
	})
}

package analysis_v1_test

import (
	"fmt"
	"net/http"
	"net/url"
	"scrapher/src/app"
	"scrapher/src/global"
	"scrapher/src/modules/analysis/api/v1/dto"
	rodext "scrapher/src/pkg/rod"
	"scrapher/tests"
	test_utils "scrapher/tests/__utils__"
	"testing"

	"github.com/akalanka47000/go-modkit/parallel_convey"
	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAnalyseWebpageHandler(t *testing.T) {
	t.Parallel()

	tests.Setup()

	app := app.New()

	ParallelConvey, WaitL1 := pc.New(t)

	reqPath := "/api/v1/analysis/webpage"

	ParallelConvey("succesful requests", t, func() {
		Convey("html version 5", func() {
			server, address := tests.ServeDirectory("__mocks__/html/v5", 6700)
			defer server.Shutdown(t.Context())

			req, _ := http.NewRequest("GET", fmt.Sprintf("%s?url=%s", reqPath, url.QueryEscape(address)), nil)
			res, err := app.Test(req, -1)

			So(err, ShouldBeNil)

			So(res.StatusCode, ShouldEqual, http.StatusOK)

			body := test_utils.ParseResponseBody[global.Response[dto.AnalyseWebpageResult]](res.Body)

			So(body.Message, ShouldEqual, "Analysis complete")

			So(body.Data, ShouldEqual, &dto.AnalyseWebpageResult{
				HTMLVersion:       "HTML5",
				PageTitle:         "Sample HTML5 Document",
				ContainsLoginForm: false,
				HeadingCounts: dto.HeadingCounts{
					H1: 1,
					H2: 4,
					H3: 1,
					H4: 0,
					H5: 0,
					H6: 0,
				},
				InternalLinkCount:     2,
				ExternalLinkCount:     4,
				InaccessibleLinkCount: 1,
			})
		})
	})

	ParallelConvey("failed requests", t, func() {
		Convey("invalid url", func() {
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s?url=invalid-url", reqPath), nil)
			res, err := app.Test(req, -1)

			So(err, ShouldBeNil)

			So(res.StatusCode, ShouldEqual, http.StatusUnprocessableEntity)

			body := test_utils.ParseResponseBody[global.Response[any]](res.Body)

			So(body.Message, ShouldEqual, "Please provide a valid url to analyse")
		})
		Convey("inaccessible url (invalid domain)", func() {
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s?url=http://invalid-domain", reqPath), nil)
			res, err := app.Test(req, -1)

			So(err, ShouldBeNil)

			So(res.StatusCode, ShouldEqual, http.StatusUnprocessableEntity)

			body := test_utils.ParseResponseBody[global.Response[any]](res.Body)

			So(body.Message, ShouldEqual, rodext.ErrConnectionError.Error())

			So(body.Data, ShouldBeNil)

			So(lo.CastJSON[rodext.RodErrorDetail](body.Error), ShouldEqual, rodext.ErrConnectionError.Detail)
		})
	})

	WaitL1()
}

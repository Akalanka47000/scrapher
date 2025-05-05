package analysis_v1_test

import (
	analysis "scrapher/src/modules/analysis/api/v1"
	"scrapher/tests"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAnalyseWebpage(t *testing.T) {
	t.Parallel()

	tests.Setup()

	Convey("existing webpages", t, func() {
		Convey("html version 5", func() {
			server, address := tests.ServeDirectory("__mocks__/html/v5", 6600)
			defer server.Shutdown(t.Context())

			result := analysis.AnalyseWebPage(address)

			So(result.HTMLVersion, ShouldEqual, "HTML5")
			So(result.PageTitle, ShouldEqual, "Sample HTML5 Document")
			So(result.HeadingCounts.H1, ShouldEqual, 1)
			So(result.HeadingCounts.H2, ShouldEqual, 4)
			So(result.HeadingCounts.H3, ShouldEqual, 1)
			So(result.HeadingCounts.H4, ShouldEqual, 0)
			So(result.HeadingCounts.H5, ShouldEqual, 0)
			So(result.HeadingCounts.H6, ShouldEqual, 0)
			So(result.InternalLinkCount, ShouldEqual, 2)
			So(result.ExternalLinkCount, ShouldEqual, 4)
			So(result.InaccessibleLinkCount, ShouldEqual, 1)
			So(result.ContainsLoginForm, ShouldBeFalse)
		})
	})
}

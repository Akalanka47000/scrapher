//nolint:dupl
package analysis_v1_test

import (
	"github.com/akalanka47000/go-modkit/parallel_convey"
	. "github.com/smartystreets/goconvey/convey"
	analysis "scrapher/src/modules/analysis/api/v1"
	rodext "scrapher/src/pkg/rod"
	"scrapher/tests"
	"testing"
)

func TestAnalyseWebpage(t *testing.T) {
	t.Parallel()

	tests.Setup()

	ParallelConvey, WaitL1 := pc.New(t)

	ParallelConvey("existing webpages", t, func() {
		ParallelConvey, WaitL2 := pc.New(t)

		ParallelConvey("html version 5", func() {
			server, address := tests.ServeDirectory("__mocks__/html/v5", 6600)
			defer server.Shutdown(t.Context())

			result := analysis.AnalyseWebPage(address)

			So(result.HTMLVersion, ShouldEqual, "HTML5")
			So(result.PageTitle, ShouldEqual, "Sample HTML5 Document")

			Convey("does not have login form", func() {
				So(result.ContainsLoginForm, ShouldBeFalse)
			})

			Convey("headings", func() {
				So(result.HeadingCounts.H1, ShouldEqual, 1)
				So(result.HeadingCounts.H2, ShouldEqual, 4)
				So(result.HeadingCounts.H3, ShouldEqual, 1)
				So(result.HeadingCounts.H4, ShouldEqual, 0)
				So(result.HeadingCounts.H5, ShouldEqual, 0)
				So(result.HeadingCounts.H6, ShouldEqual, 0)
			})

			Convey("links", func() {
				So(result.InternalLinkCount, ShouldEqual, 2)
				So(result.ExternalLinkCount, ShouldEqual, 4)
				So(result.InaccessibleLinkCount, ShouldEqual, 1)
			})
		})
		ParallelConvey("html version 4", func() {
			server, address := tests.ServeDirectory("__mocks__/html/v4", 6601)
			defer server.Shutdown(t.Context())

			result := analysis.AnalyseWebPage(address)

			So(result.HTMLVersion, ShouldEqual, "HTML 4.01 Transitional")
			So(result.PageTitle, ShouldEqual, "Sample HTML 4 Transitional Document")

			Convey("has login form", func() {
				So(result.ContainsLoginForm, ShouldBeTrue)
			})

			Convey("headings", func() {
				So(result.HeadingCounts.H1, ShouldEqual, 2)
				So(result.HeadingCounts.H2, ShouldEqual, 0)
				So(result.HeadingCounts.H3, ShouldEqual, 2)
				So(result.HeadingCounts.H4, ShouldEqual, 0)
				So(result.HeadingCounts.H5, ShouldEqual, 0)
				So(result.HeadingCounts.H6, ShouldEqual, 1)
			})

			Convey("links", func() {
				So(result.InternalLinkCount, ShouldEqual, 4)
				So(result.ExternalLinkCount, ShouldEqual, 3)
				So(result.InaccessibleLinkCount, ShouldEqual, 3)
			})
		})

		WaitL2()
	})

	ParallelConvey("non-existing webpages", t, func() {
		Convey("invalid url", func() {
			So(func() {
				analysis.AnalyseWebPage("htpss://invalid-url")
			}, ShouldPanicWith, rodext.ErrConnectionError)
		})
		Convey("invalid domain", func() {
			So(func() {
				analysis.AnalyseWebPage("https://domain-that-does-not-exist.scrapher.com")
			}, ShouldPanicWith, rodext.ErrConnectionError)
		})
		Convey("invalid content type", func() {
			server, address := tests.ServeDirectory("__mocks__/html/not-here", 6603)
			defer server.Shutdown(t.Context())

			So(func() {
				analysis.AnalyseWebPage(address)
			}, ShouldPanicWith, rodext.ErrTargetIsNotValidHTML)
		})
	})

	WaitL1()
}

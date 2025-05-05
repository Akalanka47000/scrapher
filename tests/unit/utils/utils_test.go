package utils_test

import (
	"net/url"
	"scrapher/src/utils"
	"testing"

	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils(t *testing.T) {
	t.Parallel()

	Convey("examine valid link", t, func() {
		Convey("internal link", func() {
			external, err := utils.IsExternalLink("/login", *lo.Ok(url.Parse("http://localhost:3000")))
			So(err, ShouldBeNil)
			So(external, ShouldBeFalse)
		})
		Convey("external link", func() {
			external, err := utils.IsExternalLink("http://example.com/login", *lo.Ok(url.Parse("http://localhost:3000")))
			So(err, ShouldBeNil)
			So(external, ShouldBeTrue)
		})
	})
	Convey("examine invalid link", t, func() {
		external, err := utils.IsExternalLink(":\\\\invalid-++{>?@link", *lo.Ok(url.Parse("http://localhost:3000")))
		So(err, ShouldNotBeNil)
		So(external, ShouldBeFalse)
	})
}

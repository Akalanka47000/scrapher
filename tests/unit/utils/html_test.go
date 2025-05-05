package utils_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"scrapher/src/utils"
	"testing"
)

func TestUtilsHTML(t *testing.T) {
	t.Parallel()

	Convey("html version 5", t, func() {
		html := "<!DOCTYPE html><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML5")
	})
	Convey("xhtml", t, func() {
		html := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "XHTML")
	})
	Convey("html 4.01 transitional", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\" \"http://www.w3.org/TR/html4/loose.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.01 Transitional")
	})
	Convey("html 4.01 strict", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Strict//EN\" \"http://www.w3.org/TR/html4/strict.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.01 Strict")
	})
	Convey("html 4.01 frameset", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Frameset//EN\" \"http://www.w3.org/TR/html4/frameset.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.01 Frameset")
	})
	Convey("html 4.0 transitional", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.0 Transitional")
	})
	Convey("html 4.0 strict", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0 Strict//EN\" \"http://www.w3.org/TR/REC-html40/strict.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.0 Strict")
	})
	Convey("html 4.0 frameset", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0 Frameset//EN\" \"http://www.w3.org/TR/REC-html40/frameset.dtd\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 4.0 Frameset")
	})
	Convey("html 3.2", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 3.2 Final//EN\" \"http://www.w3.org/TR/REC-html32\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 3.2")
	})
	Convey("html 2.0", t, func() {
		html := "<!DOCTYPE HTML PUBLIC \"-//IETF//DTD HTML 2.0//EN\"><html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "HTML 2.0")
	})
	Convey("unknown", t, func() {
		html := "<html><head><title>Sample Doc</title></head></html>"
		So(utils.ExtractHTMLVersion(html), ShouldEqual, "Unknown")
	})
}

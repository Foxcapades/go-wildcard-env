package wenv_test

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/go-wildcard-env/pkg/wenv"
)

func TestNewPrefixMatcher(t *testing.T) {
	Convey("prefix matcher", t, func() {
		matcher := wenv.NewPrefixMatcher("test", "MY_PREFIX_")

		So(matcher.Name(), ShouldEqual, "test")

		So(matcher.Matches("MY_PREFIX_"), ShouldBeFalse)
		So(matcher.Matches("MY_PREFIX_FOO"), ShouldBeTrue)

		res := matcher.Process("MY_PREFIX_FOO")

		So(len(res), ShouldEqual, 1)
		So(res[0], ShouldEqual, "FOO")
	})
}

func TestNewSuffixMatcher(t *testing.T) {
	Convey("suffix matcher", t, func() {
		matcher := wenv.NewSuffixMatcher("test", "_MY_SUFFIX")

		So(matcher.Name(), ShouldEqual, "test")

		So(matcher.Matches("_MY_SUFFIX"), ShouldBeFalse)
		So(matcher.Matches("BAR_MY_SUFFIX"), ShouldBeTrue)

		res := matcher.Process("BAR_MY_SUFFIX")

		So(len(res), ShouldEqual, 1)
		So(res[0], ShouldEqual, "BAR")
	})
}

func TestNewWrappedMatcher(t *testing.T) {
	Convey("wrapped matcher", t, func() {
		matcher := wenv.NewWrappedMatcher("test", "MY_PREFIX_", "_MY_SUFFIX")

		So(matcher.Name(), ShouldEqual, "test")

		So(matcher.Matches("MY_PREFIX__MY_SUFFIX"), ShouldBeFalse)
		So(matcher.Matches("MY_PREFIX_FOO_MY_SUFFIX"), ShouldBeTrue)

		res := matcher.Process("MY_PREFIX_FOO_MY_SUFFIX")

		So(len(res), ShouldEqual, 1)
		So(res[0], ShouldEqual, "FOO")
	})
}

func TestNewRegexMatcher(t *testing.T) {
	Convey("regex matcher", t, func() {
		Convey("with a valid regex pattern", func() {
			matcher := wenv.NewRegexMatcher("test", regexp.MustCompile(`^PREFIX_(\w+)_(\w+)_SUFFIX$`))

			So(matcher.Name(), ShouldEqual, "test")

			So(matcher.Matches("PREFIX___SUFFIX"), ShouldBeFalse)
			So(matcher.Matches("PREFIX_FOO__SUFFIX"), ShouldBeFalse)
			So(matcher.Matches("PREFIX__BAR_SUFFIX"), ShouldBeFalse)
			So(matcher.Matches("PREFIX_FOO_BAR_SUFFIX"), ShouldBeTrue)

			res := matcher.Process("PREFIX_FOO_BAR_SUFFIX")

			So(len(res), ShouldEqual, 2)
			So(res[0], ShouldEqual, "FOO")
			So(res[1], ShouldEqual, "BAR")

			So(func() { matcher.Process("foo") }, ShouldPanic)
		})

		Convey("with a regex containing no matching groups", func() {
			matcher := wenv.NewRegexMatcher("test", regexp.MustCompile(`^PREFIX_\w+_\w+_SUFFIX$`))

			So(func() { matcher.Process("PREFIX_FOO_BAR_SUFFIX") }, ShouldPanic)
		})
	})
}

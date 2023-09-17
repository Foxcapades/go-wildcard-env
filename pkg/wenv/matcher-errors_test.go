package wenv_test

import (
	"errors"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/go-wildcard-env/pkg/wenv"
)

func TestMatcherErrors(t *testing.T) {
	Convey("MatcherErrors", t, func() {
		errs := wenv.MatcherErrors([]error{
			errors.New("hello"),
			errors.New("you"),
			errors.New("smelly"),
			errors.New("little"),
			errors.New("biscuit"),
		})

		So(errs.Size(), ShouldEqual, 5)
		So(errs.Get(0).Error(), ShouldEqual, "hello")
		So(errs.HasErrors(), ShouldBeTrue)
		So(errs.IsEmpty(), ShouldBeFalse)
		So(errs.Error(), ShouldEqual, "encountered 5 environment parsing errors")

		sb := new(strings.Builder)

		i, e := errs.WriteLines(sb)

		So(i, ShouldEqual, 31)
		So(e, ShouldBeNil)

		So(sb.String(), ShouldEqual, "hello\nyou\nsmelly\nlittle\nbiscuit")
	})
}

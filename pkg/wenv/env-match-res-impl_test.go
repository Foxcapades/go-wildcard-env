package wenv_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/go-wildcard-env/pkg/wenv"
)

func TestEnvironmentMatcher(t *testing.T) {
	Convey("environment matcher", t, func() {

		Convey("test 1", func() {
			matcher := wenv.NewEnvironmentMatcher().
				AddGroup(wenv.NewMatchGroup("db").
					AddMatcher(wenv.NewPrefixMatcher("name", "DB_NAME_"), true).
					AddMatcher(wenv.NewPrefixMatcher("address", "DB_ADDRESS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("port", "DB_PORT_"), true).
					AddMatcher(wenv.NewPrefixMatcher("user", "DB_USER_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pass", "DB_PASS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pool", "DB_POOL_SIZE_"), false),
					true,
				)

			environ := map[string]string{
				"DB_NAME_FOO":      "my_database_1",
				"DB_ADDRESS_FOO":   "somehost",
				"DB_PORT_FOO":      "1234",
				"DB_USER_FOO":      "username",
				"DB_PASS_FOO":      "password",
				"DB_POOL_SIZE_FOO": "5",

				"DB_NAME_BAR":    "my_database_2",
				"DB_ADDRESS_BAR": "otherhost",
				"DB_PORT_BAR":    "4321",
				"DB_USER_BAR":    "user",
				"DB_PASS_BAR":    "pass",
			}

			envResult := matcher.ParseEnv(environ)

			So(envResult.Size(), ShouldEqual, 1)
			So(envResult.Has("db"), ShouldBeTrue)
			So(envResult.Errors(), ShouldBeNil)
			So(envResult.Get("foo"), ShouldBeNil)

			dbResults := envResult.Get("db")

			So(dbResults.Size(), ShouldEqual, 2)

			for i := 0; i < dbResults.Size(); i++ {
				res := dbResults.Get(i)

				if res.FirstKey() == "FOO" {
					So(res.Size(), ShouldEqual, 6)

					So(res.Has("name"), ShouldBeTrue)
					So(res.Has("address"), ShouldBeTrue)
					So(res.Has("port"), ShouldBeTrue)
					So(res.Has("user"), ShouldBeTrue)
					So(res.Has("pass"), ShouldBeTrue)
					So(res.Has("pool"), ShouldBeTrue)

					So(res.Get("name").Value(), ShouldEqual, "my_database_1")
					So(res.Get("address").Value(), ShouldEqual, "somehost")
					So(res.Get("port").Value(), ShouldEqual, "1234")
					So(res.Get("user").Value(), ShouldEqual, "username")
					So(res.Get("pass").Value(), ShouldEqual, "password")
					So(res.Get("pool").Value(), ShouldEqual, "5")

					So(res.Get("name").Raw(), ShouldEqual, "DB_NAME_FOO")
					So(res.Get("address").Raw(), ShouldEqual, "DB_ADDRESS_FOO")
					So(res.Get("port").Raw(), ShouldEqual, "DB_PORT_FOO")
					So(res.Get("user").Raw(), ShouldEqual, "DB_USER_FOO")
					So(res.Get("pass").Raw(), ShouldEqual, "DB_PASS_FOO")
					So(res.Get("pool").Raw(), ShouldEqual, "DB_POOL_SIZE_FOO")
				} else if res.FirstKey() == "BAR" {
					So(res.Size(), ShouldEqual, 5)

					So(res.Has("name"), ShouldBeTrue)
					So(res.Has("address"), ShouldBeTrue)
					So(res.Has("port"), ShouldBeTrue)
					So(res.Has("user"), ShouldBeTrue)
					So(res.Has("pass"), ShouldBeTrue)
					So(res.Has("pool"), ShouldBeFalse)

					So(res.Get("name").Value(), ShouldEqual, "my_database_2")
					So(res.Get("address").Value(), ShouldEqual, "otherhost")
					So(res.Get("port").Value(), ShouldEqual, "4321")
					So(res.Get("user").Value(), ShouldEqual, "user")
					So(res.Get("pass").Value(), ShouldEqual, "pass")

					So(res.Get("name").Raw(), ShouldEqual, "DB_NAME_BAR")
					So(res.Get("address").Raw(), ShouldEqual, "DB_ADDRESS_BAR")
					So(res.Get("port").Raw(), ShouldEqual, "DB_PORT_BAR")
					So(res.Get("user").Raw(), ShouldEqual, "DB_USER_BAR")
					So(res.Get("pass").Raw(), ShouldEqual, "DB_PASS_BAR")
				} else {
					panic("wut")
				}
			}
		})

		Convey("test 2", func() {
			matcher := wenv.NewEnvironmentMatcher().
				AddGroup(wenv.NewMatchGroup("db").
					AddMatcher(wenv.NewPrefixMatcher("name", "DB_NAME_"), true).
					AddMatcher(wenv.NewPrefixMatcher("address", "DB_ADDRESS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("port", "DB_PORT_"), true).
					AddMatcher(wenv.NewPrefixMatcher("user", "DB_USER_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pass", "DB_PASS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pool", "DB_POOL_SIZE_"), false),
					true,
				)

			environ := map[string]string{}

			envResult := matcher.ParseEnv(environ)

			So(envResult.Size(), ShouldEqual, 0)
			So(envResult.Has("db"), ShouldBeFalse)
			So(envResult.Errors(), ShouldNotBeNil)
			So(len(envResult.Errors()), ShouldEqual, 1)
			So(envResult.Errors()[0].Error(), ShouldEqual, "no environment matches found for environment group db")
		})

		Convey("test 3", func() {
			matcher := wenv.NewEnvironmentMatcher().
				AddGroup(wenv.NewMatchGroup("db").
					AddMatcher(wenv.NewPrefixMatcher("name", "DB_NAME_"), true).
					AddMatcher(wenv.NewPrefixMatcher("address", "DB_ADDRESS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("port", "DB_PORT_"), true).
					AddMatcher(wenv.NewPrefixMatcher("user", "DB_USER_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pass", "DB_PASS_"), true).
					AddMatcher(wenv.NewPrefixMatcher("pool", "DB_POOL_SIZE_"), false),
					true,
				)

			environ := map[string]string{
				"DB_NAME_FOO":      "my_database_1",
				"DB_ADDRESS_FOO":   "somehost",
				"DB_PORT_FOO":      "1234",
				"DB_PASS_FOO":      "password",
				"DB_POOL_SIZE_FOO": "5",
			}

			envResult := matcher.ParseEnv(environ)

			So(envResult.Size(), ShouldEqual, 1)
			So(envResult.Has("db"), ShouldBeTrue)

			So(envResult.Errors(), ShouldNotBeNil)
			So(envResult.Errors().Size(), ShouldEqual, 1)

			So(envResult.Errors().Get(0).Error(), ShouldEqual, "match group db (keys: FOO) does not have a match for required key user")

			dbResults := envResult.Get("db")

			So(dbResults.Size(), ShouldEqual, 1)

			for i := 0; i < dbResults.Size(); i++ {
				res := dbResults.Get(i)

				if res.FirstKey() == "FOO" {
					So(res.Size(), ShouldEqual, 5)

					So(res.Has("name"), ShouldBeTrue)
					So(res.Has("address"), ShouldBeTrue)
					So(res.Has("port"), ShouldBeTrue)
					So(res.Has("user"), ShouldBeFalse)
					So(res.Has("pass"), ShouldBeTrue)
					So(res.Has("pool"), ShouldBeTrue)

					So(res.Value("name"), ShouldEqual, "my_database_1")
					So(res.Value("address"), ShouldEqual, "somehost")
					So(res.Value("port"), ShouldEqual, "1234")
					So(res.Value("pass"), ShouldEqual, "password")
					So(res.Value("pool"), ShouldEqual, "5")

					So(res.Get("name").Raw(), ShouldEqual, "DB_NAME_FOO")
					So(res.Get("address").Raw(), ShouldEqual, "DB_ADDRESS_FOO")
					So(res.Get("port").Raw(), ShouldEqual, "DB_PORT_FOO")
					So(res.Get("pass").Raw(), ShouldEqual, "DB_PASS_FOO")
					So(res.Get("pool").Raw(), ShouldEqual, "DB_POOL_SIZE_FOO")
				} else {
					panic("wut")
				}
			}
		})
	})
}

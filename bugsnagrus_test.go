package bugsnagrus

import (
	"testing"

	"github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBugsnagHook(t *testing.T) {

	Convey("Given empty args", t, func() {

		args := func() (string, string, []logrus.Level, int) {
			return "", "", []logrus.Level{}, -1
		}

		Convey("When NewBugsnagHook", func() {

			_, err := NewBugsnagHook(args())

			Convey("Then error is occured", func() {

				So(err, ShouldNotBeNil)

			})

		})

		Convey("Given APIKey", func() {

			args = func() (string, string, []logrus.Level, int) {
				return "apiKey", "", []logrus.Level{}, -1
			}

			Convey("When NewBugsnagHook", func() {

				_, err := NewBugsnagHook(args())

				Convey("Then error is occured", func() {

					So(err, ShouldNotBeNil)

				})

			})

		})

		Convey("Given APIKey, releaseStage", func() {

			args = func() (string, string, []logrus.Level, int) {
				return "apiKey", "test", []logrus.Level{}, -1
			}

			Convey("When NewBugsnagHook", func() {

				_, err := NewBugsnagHook(args())

				Convey("Then error is occured", func() {

					So(err, ShouldNotBeNil)

				})

			})

		})

		Convey("Given APIKey, releaseStage, levels", func() {

			args = func() (string, string, []logrus.Level, int) {
				return "apiKey", "test", []logrus.Level{logrus.ErrorLevel}, -1
			}

			Convey("When NewBugsnagHook", func() {

				_, err := NewBugsnagHook(args())

				Convey("Then error is occured", func() {

					So(err, ShouldNotBeNil)

				})

			})

		})

		Convey("Given APIKey, releaseStage, levels, skip", func() {

			args = func() (string, string, []logrus.Level, int) {
				return "apiKey", "test", []logrus.Level{logrus.ErrorLevel}, 0
			}

			Convey("When NewBugsnagHook", func() {

				h, err := NewBugsnagHook(args())

				Convey("Then error should be nil", func() {

					So(err, ShouldBeNil)

					Convey("Then hook should be BugsnagHook", func() {

						So(h, ShouldHaveSameTypeAs, new(BugsnagHook))

					})

				})

			})

		})

	})

}

func TestFire(t *testing.T) {

	Convey("Given BugsnagHook", t, func() {

		h, _ := NewBugsnagHook("apiKey", "test", []logrus.Level{logrus.ErrorLevel}, 0)

		Convey("Given logrus.Entry", func() {

			entry := new(logrus.Entry)
			entry.Data = logrus.Fields{
				"userID": "1234",
				"error":  "test",
				"meta":   &BugsnagHook{Skip: 100},
			}

			Convey("Given INFO level", func() {

				entry.Level = logrus.InfoLevel

				Convey("When Fire is called", func() {

					err := h.Fire(entry)

					Convey("Then error should be nil", func() {

						So(err, ShouldBeNil)

					})

				})

			})

			Convey("Given WARN level", func() {

				entry.Level = logrus.WarnLevel

				Convey("When Fire is called", func() {

					err := h.Fire(entry)

					Convey("Then error should be nil", func() {

						So(err, ShouldBeNil)

					})

				})

			})

			Convey("Given ERRO level", func() {

				entry.Level = logrus.ErrorLevel

				Convey("When Fire is called", func() {

					err := h.Fire(entry)

					Convey("Then error should be nil", func() {

						So(err, ShouldBeNil)

					})

				})

			})

		})

	})

}

func TestLevels(t *testing.T) {

	Convey("Given levels", t, func() {

		h := BugsnagHook{Lvs: []logrus.Level{logrus.InfoLevel}}

		Convey("When Levels() is called", func() {

			l := h.Levels()

			Convey("Then logrus.InfoLevel returned", func() {

				So(len(l), ShouldEqual, 1)
				So(l[0], ShouldEqual, logrus.InfoLevel)

			})

		})

	})

}

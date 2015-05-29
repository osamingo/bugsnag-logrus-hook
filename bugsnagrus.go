package bugsnagrus

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/bugsnag/bugsnag-go"
	"github.com/bugsnag/bugsnag-go/errors"
)

// BugsnagHook implements logrus.Hook interface.
type BugsnagHook struct {
	Lvs  []logrus.Level
	Skip int
}

// NewBugsnagHook setups Bugsnag configuration and BugsnagHook.
func NewBugsnagHook(apiKey, releaseStage string, lvs []logrus.Level, skip uint) (h *BugsnagHook, err error) {

	if apiKey == "" {
		err = fmt.Errorf("apiKey should not be empty")
		return
	}

	if releaseStage == "" {
		err = fmt.Errorf("releaseStage should not be empty")
		return
	}

	if len(lvs) == 0 {
		err = fmt.Errorf("lvs should not be empty")
		return
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       apiKey,
		ReleaseStage: releaseStage,
	})

	h = &BugsnagHook{Lvs: lvs, Skip: int(skip)}

	return
}

// Fire forwards an error to Bugsnag.
func (h *BugsnagHook) Fire(entry *logrus.Entry) error {

	s := []interface{}{}

	if m := entry.Data["meta"]; m != nil {

		md, ok := m.(bugsnag.MetaData)
		if !ok {
			md = bugsnag.MetaData{}
			md.AddStruct("meta", m)
		}

		s = append(s, md)
	}

	if u := entry.Data["userID"]; u != nil {
		s = append(s, bugsnag.User{Id: fmt.Sprintf("%v", u)})
	}

	if e := entry.Data["error"]; e != nil {
		s = append(s, bugsnag.ErrorClass{Name: fmt.Sprintf("%v", e)})
	}

	if entry.Level == logrus.InfoLevel {
		s = append(s, bugsnag.SeverityInfo)
	} else if entry.Level == logrus.WarnLevel {
		s = append(s, bugsnag.SeverityWarning)
	} else {
		s = append(s, bugsnag.SeverityError)
	}

	return bugsnag.Notify(errors.New(entry.Message, h.Skip), s)
}

// Levels enumerates the log levels on which the error should be forwarded to bugsnag.
func (h *BugsnagHook) Levels() []logrus.Level {
	return h.Lvs
}

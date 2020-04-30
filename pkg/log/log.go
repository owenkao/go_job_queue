package log

import (
	"github.com/sirupsen/logrus"
)

var logconsole *logrus.Entry

// Alert log api
func Alert(origin, category, event, user string, format string, a ...interface{}) error {
	logconsole.WithField("origin", origin).WithField("event", event).WithField("category", category).WithField("user", user).Errorf(format, a...)
	//return Log("alert", origin, category, event, user, format, a...)
}

func init() {
	logconsole = logrus.WithFields(logrus.Fields{"origin": "logservice"})
}

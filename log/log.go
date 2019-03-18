package log

import (
	"github.com/Madq92/etcd-http-client/config"
	"github.com/sirupsen/logrus"
	"os"
)

var LOGGER *logrus.Logger

func init() {
	LOGGER = logrus.New()
	LOGGER.Formatter = &logrus.TextFormatter{}
	LOGGER.Out = os.Stdout
	level, err := logrus.ParseLevel(config.Conf.Log.Level)
	if err != nil {
		logrus.WithError(err)
		level = logrus.DebugLevel
	}
	LOGGER.SetLevel(level)

	LOGGER = logrus.New()
	LOGGER.Formatter = &logrus.TextFormatter{}
}

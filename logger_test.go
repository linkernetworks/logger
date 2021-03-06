package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	logger.Info("test")
	logger.Infoln("test", "line2")
	logger.Debug("test", "line2")
}

func TestLoggerConfigure(t *testing.T) {
	logger := logrus.New()
	c := LoggerConfig{
		Dir:           "./logs",
		SuffixPattern: ".%Y%m%d",
		LinkName:      "access_log",
		Level:         "info",
		MaxAge:        "7d",
	}
	configure(logger, c)

	assert.Equal(t, c.Level, logger.Level.String())
	// XXX add tests for the logrotater
}

package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func LoggerConfig() *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.WarnLevel)
	log.SetFormatter(&logrus.JSONFormatter{}) // Log as JSON instead of the default ASCII formatter.
	log.SetOutput(os.Stdout)

	return log
}

package logrus_lumberjack

import (
  "fmt"

  "gopkg.in/natefinch/lumberjack.v2"
  "github.com/sirupsen/logrus"
)

// LumberjackHook stores the configuration of the hook
type LumberjackHook struct {
  logger *lumberjack.Logger
}

// NewLumberjackHook creates a new LumberjackHook
func NewLumberjackHook(logger *lumberjack.Logger) (*LumberjackHook, error) {
  if logger == nil {
    return nil, fmt.Errorf("Logger cannot be nil")
  }

  return &LumberjackHook{
    logger: logger,
  }, nil
}

// Fire is called when a log event is fired.
func (hook *LumberjackHook) Fire(entry *logrus.Entry) error {
  // Convert the line to string
  line, err := entry.String()
  if err != nil {
    return err
  }

  // Write the the logger
  _, err = hook.logger.Write([]byte(line))
  if err != nil {
    return err
  }

  return nil
}

// Levels returns the available logging levels
func (hook *LumberjackHook) Levels() []logrus.Level {
  return []logrus.Level{
    logrus.DebugLevel,
    logrus.InfoLevel,
    logrus.WarnLevel,
    logrus.ErrorLevel,
    logrus.FatalLevel,
    logrus.PanicLevel,
  }
}
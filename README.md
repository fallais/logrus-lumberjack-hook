# Lumberjack Hooks for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

## Usage

```go
// Set the Lumberjack logger
lumberjackLogger := &lumberjack.Logger{
  Filename:   "/var/log/misc.log",
  MaxSize:    10,
  MaxBackups: 3,
  MaxAge:     3,
  LocalTime:  true,
}

// Add Lumberjack hook
hook, err := logrus_lumberjack.NewLumberjackHook(lumberjackLogger)
if err != nil {
  logrus.Fatalln("Unable to add the Lumberjack hook :", err)
} else {
  logrus.AddHook(hook)
}
```
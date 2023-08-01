package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	_logger = logrus.New()
)

type (
	Level = logrus.Level
	KVs   = logrus.Fields
)

func SetLevel(levelName string) {
	var _level Level
	switch levelName {
	case "panic":
		_level = logrus.PanicLevel
	case "fatal":
		_level = logrus.FatalLevel
	case "error":
		_level = logrus.ErrorLevel
	case "warn":
		_level = logrus.WarnLevel
	case "info":
		_level = logrus.InfoLevel
	case "debug":
		_level = logrus.DebugLevel
	case "trace":
		_level = logrus.TraceLevel
	default:
		_level = logrus.DebugLevel
	}
	_logger.SetLevel(_level)
}

func Init(c *Config) {
	SetLevel(c.Level)

	initLogger(c)
}

func initLogger(c *Config) {
	// 输出时间格式化
	timestampFormat := c.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	// 根据环境定义输出格式
	switch c.Env {
	case "dev":
		setTextFormatter(timestampFormat)
		// 输出方式，默认os.stderr
		_logger.SetOutput(os.Stdout)
		// 定位行号
		//_logger.SetReportCaller(true)
	case "prod":
		setJsonFormatter(timestampFormat)

		logPath := c.Logfile
		if logPath == "" {
			logPath = "./app.log"
		}
		// 输出方式，默认os.stderr
		logfile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		_logger.SetOutput(logfile)
	case "uat":
		setJsonFormatter(timestampFormat)

		logPath := c.Logfile
		if logPath == "" {
			logPath = "./app.log"
		}
		logfile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		_logger.SetOutput(logfile)
	default:
		setTextFormatter(timestampFormat)
		_logger.SetOutput(os.Stdout)
		//_logger.SetReportCaller(true)
	}
}

func setTextFormatter(timestampFormat string) {
	_logger.SetFormatter(&logrus.TextFormatter{
		// 非TTY中输出彩色
		ForceColors: true,
		// 非TTY中输出彩色时需要开启
		FullTimestamp: true,
		// 禁用截取
		// DisableLevelTruncation: true,
		// 时间格式
		TimestampFormat: timestampFormat,
	})
}

func setJsonFormatter(timestampFormat string) {
	_logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false, // json输出格式化
		// 时间格式
		TimestampFormat: timestampFormat,
	})
}

func Trace(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Trace(args...)
	} else {
		_logger.Trace(args...)
	}
}

func Debug(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Debug(args...)
	} else {
		_logger.Debug(args...)
	}
}

func Print(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Print(args...)
	} else {
		_logger.Print(args...)
	}
}

func Info(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Info(args...)
	} else {
		_logger.Info(args...)
	}
}

func Warn(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Warn(args...)
	} else {
		_logger.Warn(args...)
	}
}

func Warning(args ...interface{}) {
	Warn(args...)
}

func Error(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Error(args...)
	} else {
		_logger.Error(args...)
	}
}

func Fatal(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Fatal(args...)
	} else {
		_logger.Fatal(args...)
	}
}

func Panic(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		args = args[1:]
		_logger.WithFields(kvs.(KVs)).Panic(args...)
	} else {
		_logger.Panic(args...)
	}
}

func Tracef(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Tracef(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Tracef(format, args...)
	}
}

func Debugf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Debugf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Debugf(format, args...)
	}
}

func Printf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Printf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Printf(format, args...)
	}
}

func Infof(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Infof(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Infof(format, args...)
	}
}

func Warnf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Warnf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Warnf(format, args...)
	}
}

func Warningf(args ...interface{}) {
	Warnf(args...)
}

func Errorf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Errorf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Errorf(format, args...)
	}
}

func Fatalf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Fatalf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Fatalf(format, args...)
	}
}

func Panicf(args ...interface{}) {
	kvs := args[0]
	if _, ok := kvs.(KVs); ok {
		format := args[1].(string)
		args = args[2:]
		_logger.WithFields(kvs.(KVs)).Panicf(format, args...)
	} else {
		format := args[0].(string)
		args = args[1:]
		_logger.Panicf(format, args...)
	}
}

package service

import (
	"context"
	"devops-platform/internal/common/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"os"
	"time"
)

type logConfig interface {
	GetOutput() string
	GetFormatter() string
	GetFilePath() string
	GetLevel() string
	GetTimestampFormat() string
	GetSlowThreshold() time.Duration
}

/**
 * 封装logrus为标准logger，为其他组件使用 ，如：gorm
 */
type Logger struct {
	slowThreshold time.Duration
	logLevel      logger.LogLevel
}

// LogMode log mode
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}

// Info print info
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithContext(ctx).Infof(msg, data...)
}

// Warn print warn messages
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithContext(ctx).Warnf(msg, data...)
}

// Error print error messages
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithContext(ctx).Errorf(msg, data...)
}

// Trace print sql message
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logger.Error:
		sql, rows := fc()
		if rows == -1 {
			logrus.WithField("line", utils.FileWithLineNum()).WithError(err).WithField("const(ms)", float64(elapsed/time.Millisecond)).Error(sql)
		} else {
			logrus.WithField("line", utils.FileWithLineNum()).WithError(err).WithField("const(ms)", float64(elapsed.Nanoseconds())/1e6).WithField("rows", rows).Error(sql)
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLevel >= logger.Warn:
		sql, rows := fc()

		if rows == -1 {
			logrus.WithField("line", utils.FileWithLineNum()).WithField("const(ms)", float64(elapsed.Nanoseconds())/1e6).Warn(sql)
		} else {
			logrus.WithField("line", utils.FileWithLineNum()).WithField("const(ms)", float64(elapsed.Nanoseconds())/1e6).WithField("rows", rows).Warn(sql)
		}
	case l.logLevel >= logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logrus.WithField("line", utils.FileWithLineNum()).WithField("const(ms)", float64(elapsed.Nanoseconds())/1e6).Info(sql)
		} else {
			logrus.WithField("line", utils.FileWithLineNum()).WithField("const(ms)", float64(elapsed.Nanoseconds())/1e6).WithField("rows", rows).Info(sql)
		}
	}
}

func (l *Logger) PreInject(getBean func(string) interface{}) {
	logrus.Info("log pre inject")
	logConfig, ok := getBean(config.BeanLog).(logConfig)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", config.BeanLog)
		return
	}
	//日志位置
	setOutput(logConfig)
	//日志级别
	setLevel(logConfig)
	//日志格式
	setFormatter(logConfig)

	l.slowThreshold = logConfig.GetSlowThreshold()

}




func setOutput(logConfig logConfig) {
	//日志位置
	if logConfig.GetOutput() == "file" {
		logfile, _ := os.OpenFile(logConfig.GetFilePath(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		logrus.SetOutput(logfile)
	} else {
		logrus.SetOutput(os.Stdout)
	}
}

//日志级别
func setLevel(logConfig logConfig) {
	switch logConfig.GetLevel() {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func setLogLevel(log *Logger, logConfig logConfig) {
	switch logConfig.GetLevel() {
	case "panic", "fatal":
		log.logLevel = logger.Silent
	case "error":
		log.logLevel = logger.Error
	case "warn":
		log.logLevel = logger.Warn
	default:
		log.logLevel = logger.Info
	}
}

func setFormatter(logConfig logConfig) {
	//日志格式
	if logConfig.GetFormatter() == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false, // json输出格式化
			// 时间格式
			TimestampFormat: logConfig.GetTimestampFormat(),
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			// 非TTY中输出彩色
			ForceColors: true,
			// 非TTY中输出彩色时需要开启
			FullTimestamp: true,
			// 禁用截取
			// DisableLevelTruncation: true,
			// 时间格式
			TimestampFormat: logConfig.GetTimestampFormat(),
		})
	}
}

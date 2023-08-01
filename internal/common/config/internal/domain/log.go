package domain

import "time"

type log struct {
	Output          string `toml:"output"`
	Formatter       string
	FilePath        string `toml:"file_path"`
	Level           string `toml:"level"`
	TimestampFormat string
	SlowThreshold   time.Duration
}

func (log *log) GetOutput() string {
	return log.Output
}

func (log *log) GetFormatter() string {
	if log.Formatter == "" {
		return "text"
	}
	return log.Formatter
}

func (log *log) GetFilePath() string {
	if log.FilePath == "" {
		return "./app.log"
	}
	return log.FilePath
}

func (log *log) GetLevel() string {
	if log.Level == "" {
		return "debug"
	}
	return log.Level
}

func (log *log) GetTimestampFormat() string {
	if log.TimestampFormat == "" {
		return "2006-01-02 15:04:05"
	}
	return log.TimestampFormat
}

func (log *log) GetSlowThreshold() time.Duration {
	if log.SlowThreshold <= 0 {
		return time.Second
	}
	return log.SlowThreshold
}

package util

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	LogOption func(*logOption)

	logOption struct {
		level     slog.Level
		addSource bool
		// To disable the log writer for STDOut.
		// In default, DisableStdout is false (It appends log lines to stdout.)
		disableStdout bool

		// LogFilename is the file to write logs to.
		// If the LogFilename is empty the logger handler will skip the FS log writer.
		logFilename            string
		logFileRotateMaxSizeMB int
		logFileRotateMaxBackup int
		logFileRtateCompress   bool
	}

	nullWriter struct{}
)

const (
	LOG_LEVEL_INFO  = "info"
	LOG_LEVEL_DEBUG = "debug"
	LOG_LEVEL_WARN  = "warn"
	LOG_LEVEL_ERROR = "error"

	logFileRotateMinSizeMB     = 5
	logFileRotateDefaultSizeMB = 100
	logFileRotateMinBackup     = 1
	logFileRotateDefaultBackup = 5
)

func NewLogger(opts ...LogOption) *slog.Logger {
	logOpts := logOption{
		level:                  slog.LevelInfo,
		addSource:              false,
		disableStdout:          false,
		logFilename:            "",
		logFileRotateMaxSizeMB: logFileRotateDefaultSizeMB,
		logFileRotateMaxBackup: logFileRotateDefaultBackup,
		logFileRtateCompress:   false,
	}
	for _, opt := range opts {
		opt(&logOpts)
	}

	logWriter := makeLogWriter(logOpts)
	handler := slog.NewJSONHandler(logWriter,
		&slog.HandlerOptions{Level: logOpts.level, AddSource: logOpts.addSource})
	return slog.New(handler)
}

func (nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func LoggerWithSource(addSource bool) LogOption {
	return func(opt *logOption) {
		opt.addSource = addSource
	}
}

func LoggerWithLevel(levelStr string) LogOption {
	return func(opt *logOption) {
		switch strings.ToLower(strings.TrimSpace(levelStr)) {
		case LOG_LEVEL_DEBUG:
			opt.level = slog.LevelDebug
		case LOG_LEVEL_WARN:
			opt.level = slog.LevelWarn
		case LOG_LEVEL_ERROR:
			opt.level = slog.LevelError
		case LOG_LEVEL_INFO:
			opt.level = slog.LevelInfo
		default:
			opt.level = slog.LevelInfo
		}
	}
}

func makeLogWriter(opts logOption) io.Writer {
	logFilename := strings.TrimSpace(opts.logFilename)
	if logFilename == "" && opts.disableStdout {
		return &nullWriter{}
	}

	writers := make([]io.Writer, 0, 2)
	if !opts.disableStdout {
		writers = append(writers, os.Stdout)
	}
	if logFilename != "" {
		writers = append(writers, &lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    opts.logFileRotateMaxSizeMB,
			MaxBackups: opts.logFileRotateMaxBackup,
			MaxAge:     0,
			Compress:   opts.logFileRtateCompress,
		})
	}
	return io.MultiWriter(writers...)
}

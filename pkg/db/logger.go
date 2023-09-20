package db

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	gormlogger "gorm.io/gorm/logger"

	"github.com/go-ostrich/pkg/log"
)

func NewLogger(level int, log *log.Logger) gormlogger.Interface {
	return &logger{
		LogLevel:      gormlogger.LogLevel(level),
		log:           log.With("module", "gorm"),
		SlowThreshold: 200 * time.Millisecond,
	}
}

type logger struct {
	log           *log.Logger
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

func (l *logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level

	return &newLogger
}

func (l *logger) Info(ctx context.Context, s string, i ...interface{}) {
	l.log.WithContext(ctx).Info(fmt.Sprintf(s, i...))
}

func (l *logger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.log.WithContext(ctx).Warn(fmt.Sprintf(s, i...))
}

func (l *logger) Error(ctx context.Context, s string, i ...interface{}) {
	l.log.WithContext(ctx).Error(fmt.Sprintf(s, i...))
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, rows := fc()
		l.log.WithContext(ctx).Error(err.Error(), "file", fileWithLineNum(), "err", err, "sql", sql+";",
			"rows", rows, "elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		l.log.WithContext(ctx).Warn(slowLog, "file", fileWithLineNum(), "sql", sql+";", "rows", rows,
			"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()

		l.log.WithContext(ctx).Info("gorm trace", "file", fileWithLineNum(), "sql", sql+";", "rows", rows,
			"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6))
	}
}

func fileWithLineNum() string {
	for i := 4; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		// if ok && (!strings.HasPrefix(file, gormSourceDir) || strings.HasSuffix(file, "_test.go")) {
		if ok && !strings.HasSuffix(file, "_test.go") {
			dir, f := filepath.Split(file)

			return filepath.Join(filepath.Base(dir), f) + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}

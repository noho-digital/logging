//go:generate go run github.com/noho-digital/enumer -trimprefix=Level -transform screaming-snake -type=Level
package logging

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

type Level int
const (
	DefaultLevel = LevelInfo
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelDPanic
	LevelPanic
	LevelFatal
)


func LevelZap(l zapcore.Level) Level {
	switch l {
	case zapcore.DebugLevel:
		return LevelDebug
	case zapcore.InfoLevel:
		return LevelInfo
	case zapcore.WarnLevel:
		return LevelWarn
	case zapcore.ErrorLevel:
		return LevelError
	case zapcore.DPanicLevel:
		return LevelDPanic
	case zapcore.PanicLevel:
		return LevelPanic
	case zapcore.FatalLevel:
		return LevelFatal
	default:
		return DefaultLevel
	}
}

func LevelLogRUs(l logrus.Level) Level {
	switch l {
	case logrus.TraceLevel:
		return LevelDebug
	case logrus.DebugLevel:
		return LevelDebug
	case logrus.InfoLevel:
		return LevelInfo
	case logrus.WarnLevel:
		return LevelWarn
	case logrus.ErrorLevel:
		return LevelError
	case logrus.PanicLevel:
		return LevelPanic
	case logrus.FatalLevel:
		return LevelFatal
	default:
		return DefaultLevel
	}
}

func (i Level) Zap() zapcore.Level {
	switch i {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.FatalLevel
	case LevelDPanic:
		return zapcore.DPanicLevel
	case LevelPanic:
		return zapcore.PanicLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return DefaultLevel.Zap()
	}
}


func (i Level) LogRUs() logrus.Level {
	switch i {
	case LevelDebug:
		return logrus.DebugLevel
	case LevelInfo:
		return logrus.InfoLevel
	case LevelWarn:
		return logrus.WarnLevel
	case LevelError:
		return logrus.ErrorLevel
	case LevelDPanic:
		return logrus.PanicLevel
	case LevelPanic:
		return logrus.PanicLevel
	case LevelFatal:
		return logrus.FatalLevel
	default:
		return DefaultLevel.LogRUs()
	}
}

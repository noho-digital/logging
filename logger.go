package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	StdLogger() *log.Logger
	ZapLogger() *zap.Logger
}

type Option func(*logger)

func WithZap(zapLogger *zap.Logger) Option {
	return func(l *logger) {
		l.zapLogger = zapLogger
	}
}

type logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
	stdLogger *log.Logger
}

func (l logger) ZapLogger() *zap.Logger {
	return l.zapLogger
}

func (l logger) StdLogger() *log.Logger {
	return l.stdLogger
}

func NewLogger(options ...Option) Logger {
	l := &logger{}
	for _, o := range options {
		o(l)
	}
	if l.zapLogger == nil {
		cfg := zap.NewDevelopmentEncoderConfig()
		encoder := zapcore.NewConsoleEncoder(cfg)
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(zapcore.ErrorLevel)
		l.zapLogger = zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), atomicLevel))

	}
	l.SugaredLogger = l.zapLogger.Sugar()
	l.stdLogger = zap.NewStdLog(l.zapLogger)
	return l
}

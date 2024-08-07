package logging

import (
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type WriteSyncer = zapcore.WriteSyncer

var defaultOutput WriteSyncer = os.Stderr

type Logger interface {
	Output() WriteSyncer
	SetOutput(WriteSyncer)
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
	Format() Format
	SetFormat(Format)
	Level() Level
	SetLevel(Level)
	LogRUsLogger() *logrus.Logger
	ZapLogger() *zap.Logger
	ZapSugaredLogger() *zap.SugaredLogger
	With(args ...interface{}) Logger
}

func NewLogger(options ...Option) Logger {
	l := &logger{
		level:       DefaultLevel,
		format:      DefaultFormat,
		preset:      DefaultPreset,
		output:      defaultOutput,
		zapMutex:    &sync.RWMutex{},
		logrusMutex: &sync.RWMutex{},
		stdMutex:    &sync.RWMutex{},
	}

	// default zap pre-optiosn
	if l.zap == nil {
		cfg := zap.NewDevelopmentEncoderConfig()
		encoder := zapcore.NewConsoleEncoder(cfg)
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(zapcore.ErrorLevel)
		l.zapAtomicLevel = &atomicLevel
		l.zap = zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), atomicLevel))
		l.format = FormatText
	}

	l.SugaredLogger = l.zap.Sugar()
	l.std = zap.NewStdLog(l.zap)
	l.logrus = logrus.New()
	l.logrus.SetFormatter(l.Format().LogRUsFormatter())

	// to set logrus and atomic levels /formats  if necessary
	l.SetFormat(l.format)
	l.SetLevel(l.level)
	for _, o := range options {
		o(l)
	}
	return l
}

type logger struct {
	*zap.SugaredLogger
	zap            *zap.Logger
	zapAtomicLevel *zap.AtomicLevel
	std            *log.Logger
	logrus         *logrus.Logger
	format         Format
	level          Level
	output         WriteSyncer
	preset         Preset
	zapMutex       *sync.RWMutex
	logrusMutex    *sync.RWMutex
	stdMutex       *sync.RWMutex
}

func (l *logger) resetZap() {
	l.zapMutex.Lock()
	defer l.zapMutex.Unlock()
	cfg := l.Preset().ZapEncoderConfig()
	encoder := l.Format().ZapEncoder(cfg)
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(l.Level().Zap())
	l.zap = zap.New(zapcore.NewCore(encoder, zapcore.Lock(l.Output()), atomicLevel))
	l.SugaredLogger = l.zap.Sugar()
}

func (l *logger) Output() WriteSyncer {
	return l.output
}

func (l *logger) SetOutput(w WriteSyncer) {
	if w == nil {
		return
	}
	l.output = w
	l.LogRUsLogger().SetOutput(w)
	l.resetZap()
}

func (l *logger) Format() Format {
	return l.format
}

func (l *logger) SetFormat(format Format) {
	if !format.IsAFormat() {
		return
	}
	l.format = format
	l.LogRUsLogger().SetFormatter(format.LogRUsFormatter())
	l.resetZap()
}

func (l *logger) Preset() Preset {
	return l.preset
}

func (l *logger) SetPreset(p Preset) {
	l.preset = p

}
func (l *logger) Level() Level {
	return l.level
}

func (l *logger) SetLevel(level Level) {
	if !level.IsALevel() {
		return
	}
	l.level = level
	if l.zapAtomicLevel != nil {
		atomicLevel := *l.zapAtomicLevel
		atomicLevel.SetLevel(level.Zap())
	} else {
		l.resetZap()
	}
	l.LogRUsLogger().SetLevel(level.LogRUs())
}

func (l *logger) With(args ...interface{}) Logger {
	sl := l.ZapSugaredLogger().With(args...)
	return NewLogger(
		WithLevel(l.Level()),
		WithPreset(l.Preset()),
		WithFormat(l.Format()),
		WithOutput(l.Output()),
		WithZap(sl.Desugar()),
		WithZapSugared(sl))
}

func (l *logger) ZapSugaredLogger() *zap.SugaredLogger {
	l.zapMutex.RLock()
	defer l.zapMutex.RUnlock()
	return l.SugaredLogger
}

func (l *logger) ZapLogger() *zap.Logger {
	l.zapMutex.RLock()
	defer l.zapMutex.RUnlock()
	return l.zap
}

func (l *logger) StdLogger() *log.Logger {
	return l.std
}

func (l *logger) LogRUsLogger() *logrus.Logger {
	l.logrusMutex.RLock()
	defer l.logrusMutex.RUnlock()
	return l.logrus
}

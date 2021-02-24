package logging

import "go.uber.org/zap"

type Option func(*logger)

func WithPreset(p Preset) Option {
	return func(l *logger) {
		if l == nil {
			return
		}
		l.SetPreset(p)
	}
}

func WithLevel(lvl Level) Option {
	return func(l *logger) {
		if l == nil {
			return
		}
		l.SetLevel(lvl)
	}
}

func WithOutput(syncer WriteSyncer) Option {
	return func(l *logger) {
		if l == nil {
			return
		}
		l.SetOutput(syncer)
	}
}

func WithFormat(f Format) Option {
	return func(l *logger) {
		if l == nil {
			return
		}
		l.SetFormat(f)
	}
}

func WithZapAtomicLevel(atomicLevel zap.AtomicLevel)  Option {
	return func(l *logger) {
		if l == nil {
			return
		}
		l.zapAtomicLevel = &atomicLevel
	}
}

func WithZap(zapper *zap.Logger) Option {
	return func(l *logger) {
		if l == nil || zapper == nil {
			return
		}
		l.zap = zapper
		l.SugaredLogger = l.zap.Sugar()
	}
}

func WithZapSugared(sugar *zap.SugaredLogger) Option {
	return func(l *logger) {
		if l == nil || sugar == nil {
			return
		}
		l.SugaredLogger = sugar
		l.zap = sugar.Desugar()
	}
}

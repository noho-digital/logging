//go:generate go run github.com/noho-digital/enumer -trimprefix=Format -transform snake -type=Format
package logging

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

type Format int
const (
	DefaultFormat = FormatText
	FormatText Format = iota
	FormatJSON
)


func (i Format) ZapEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch i {
	case FormatText:
		return zapcore.NewConsoleEncoder(cfg)
	case FormatJSON:
		return zapcore.NewJSONEncoder(cfg)
	default:
		return DefaultFormat.ZapEncoder(cfg)
	}

}


func (i Format) LogRUsFormatter() logrus.Formatter{
	switch i {
	case FormatText:
		return &logrus.TextFormatter{}
	case FormatJSON:
		return &logrus.JSONFormatter{}
	default:
		return DefaultFormat.LogRUsFormatter()
	}
}

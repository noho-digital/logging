//go:generate go run github.com/noho-digital/enumer -trimprefix=Preset -transform screaming-snake -type=Preset
package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Preset int

const (
	DefaultPreset = PresetProduction
	PresetDevelopment Preset = iota
	PresetProduction
)


func (i Preset) ZapEncoderConfig() zapcore.EncoderConfig {
	switch i {
	case PresetDevelopment:
		return zap.NewDevelopmentEncoderConfig()
	case PresetProduction:
		return zap.NewProductionEncoderConfig()
	default:
		return DefaultPreset.ZapEncoderConfig()
	}
}

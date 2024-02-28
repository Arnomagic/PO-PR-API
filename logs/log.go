package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "Timstamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Encoding = "json"
	config.OutputPaths = []string{"hexapi.log"}
	Log, _ = config.Build(zap.AddCallerSkip(0))
}

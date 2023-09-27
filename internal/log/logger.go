package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionConfig()
	encoderConfig.EncoderConfig.TimeKey = "timestamp"
	encoderConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	zapLogger, err := encoderConfig.Build()
	if err != nil {
		log.Fatalf("fail to build log. err: %s", err)
	}

	logger = zapLogger.With(zap.String("app", "gp-prototype"))
}

func Logger() *zap.Logger {
	return logger
}

package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
    config := zap.NewProductionConfig()
    configEncoder := zap.NewProductionEncoderConfig()
    configEncoder.TimeKey = "timestamp"
    configEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
    config.EncoderConfig = configEncoder
    return config.Build()
}

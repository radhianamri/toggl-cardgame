package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

const (
	TIME_LAYOUT = "2006-01-02 15:04:05.000"
	LOG_LEVEL   = zapcore.DebugLevel
)

func Init() {
	writers := getLogWriter()
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(writers...), LOG_LEVEL)
	logger = zap.New(core).Sugar()
}

func getLogWriter() (writers []zapcore.WriteSyncer) {
	writers = append(writers, zapcore.AddSync(os.Stdout))
	return writers
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(TIME_LAYOUT)
	cfg.EncodeDuration = zapcore.MillisDurationEncoder
	cfg.EncodeLevel = zapcore.LevelEncoder(func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	})
	return zapcore.NewConsoleEncoder(cfg)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

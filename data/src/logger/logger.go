package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

// func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(filepath.Base(caller.FullPath()))
// }

func Init() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "../log/server/server.log",
		MaxSize:    1,    // 单一档案最大几M
		MaxBackups: 10,   // 最多保留几份
		MaxAge:     7,    // 最多保留几天
		Compress:   true, // 压缩成gz
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:    "level",
		TimeKey:     "time",
		MessageKey:  "message",
		NameKey:     "name", // 可以放自定义x-api-id
		CallerKey:   "caller",
		FunctionKey: "func",
		// StacktraceKey: "trace",
		// LineEnding:     "\r\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// EncodeCaller:   MyCaller, // 自定义
	})
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	logger = zap.New(core, zap.AddCaller(),
		zap.AddCallerSkip(1),
		// zap.AddStacktrace(zap.DebugLevel),
	)
}

func Close() {
	logger.Sync()
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func InfofName(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Infof(template, args...)
}

func ErrorfName(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Errorf(template, args...)
}

func WarnfName(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Warnf(template, args...)
}

func DebugfName(name, template string, args ...interface{}) {
	logger.Sugar().Named(name).Debugf(template, args...)
}

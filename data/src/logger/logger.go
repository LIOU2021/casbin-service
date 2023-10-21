package logger

import (
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var loggerForAccess *zap.Logger

//	func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
//		enc.AppendString(filepath.Base(caller.FullPath()))
//	}

func GetAccessLogger() *zap.Logger {
	return loggerForAccess
}

func Init() {
	initBase()
	initForAccess()
}

func initBase() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "../log/server/server.log",
		MaxSize:    1,    // 单一档案最大几M
		MaxBackups: 10,   // 最多保留几份
		MaxAge:     7,    // 最多保留几天
		Compress:   true, // 压缩成gz
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encodeConfig := zapcore.EncoderConfig{
		LevelKey:    "level",
		TimeKey:     "time",
		MessageKey:  "message",
		NameKey:     "logger", // 可以放自定义x-api-id
		CallerKey:   "caller",
		FunctionKey: "func",
		// StacktraceKey: "trace",
		// LineEnding:     "\r\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// EncodeCaller:   MyCaller, // 自定义
	}

	encoder := zapcore.NewJSONEncoder(encodeConfig)

	// write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)

	// tee core
	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			stdoutSyncer, // 打印到console
			zapcore.DebugLevel,
		),
		zapcore.NewCore(
			encoder,
			writeSyncer, // 打印到server.log
			zapcore.InfoLevel,
		),
	)

	logger = zap.New(core, zap.AddCaller(),
		zap.AddCallerSkip(1),
		// zap.AddStacktrace(zap.DebugLevel),
	)

}

func initForAccess() {
	lumberJackLoggerForAccess := &lumberjack.Logger{
		Filename:   "../log/api/access.log",
		MaxSize:    1,    // 单一档案最大几M
		MaxBackups: 10,   // 最多保留几份
		MaxAge:     7,    // 最多保留几天
		Compress:   true, // 压缩成gz
	}
	writeSyncerForAccess := zapcore.AddSync(lumberJackLoggerForAccess)
	encodeConfigForAccess := zapcore.EncoderConfig{
		// LevelKey:    "level",
		// TimeKey:    "time",
		MessageKey: "message",
		NameKey:    "logger", // 可以放自定义x-api-id
		// CallerKey:   "caller",
		// FunctionKey: "func",
		// StacktraceKey: "trace",
		// LineEnding:     "\r\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// EncodeCaller:   MyCaller, // 自定义
	}
	encoderForAccess := zapcore.NewConsoleEncoder(encodeConfigForAccess)
	coreForAccess := zapcore.NewTee(
		zapcore.NewCore(
			encoderForAccess,
			zapcore.Lock(os.Stdout), // 打印到console
			zapcore.DebugLevel,
		),
		zapcore.NewCore(
			encoderForAccess,
			writeSyncerForAccess, // 打印到access.log
			zapcore.InfoLevel,
		),
	)
	loggerForAccess = zap.New(coreForAccess)
}

func Close() {
	logger.Sync()
	loggerForAccess.Sync()
}

// func Info(msg string, fields ...zapcore.Field) {
// 	logger.Info(msg, fields...)
// }

// func Error(msg string, fields ...zapcore.Field) {
// 	logger.Error(msg, fields...)
// }

// func Warn(msg string, fields ...zapcore.Field) {
// 	logger.Warn(msg, fields...)
// }

// func Debug(msg string, fields ...zapcore.Field) {
// 	logger.Debug(msg, fields...)
// }

// func InfoCtx(c *gin.Context, msg string, fields ...zapcore.Field) {
// 	logger.With(zap.String("X-Request-Id", requestid.Get(c))).Info(msg, fields...)
// }

// func ErrorCtx(c *gin.Context, msg string, fields ...zapcore.Field) {
// 	logger.With(zap.String("X-Request-Id", requestid.Get(c))).Error(msg, fields...)
// }

// func WarnCtx(c *gin.Context, msg string, fields ...zapcore.Field) {
// 	logger.With(zap.String("X-Request-Id", requestid.Get(c))).Warn(msg, fields...)
// }

// func DebugCtx(c *gin.Context, msg string, fields ...zapcore.Field) {
// 	logger.With(zap.String("X-Request-Id", requestid.Get(c))).Debug(msg, fields...)
// }

// func InfoName(name string, args ...interface{}) {
// 	logger.Sugar().Named(name).Info(args...)
// }

// func ErrorName(name string, args ...interface{}) {
// 	logger.Sugar().Named(name).Error(args...)
// }

// func WarnName(name string, args ...interface{}) {
// 	logger.Sugar().Named(name).Warn(args...)
// }

// func DebugName(name string, args ...interface{}) {
// 	logger.Sugar().Named(name).Debug(args...)
// }

// func InfoNameCtx(c *gin.Context, name string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Info(args...)
// }

// func ErrorNameCtx(c *gin.Context, name string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Error(args...)
// }

// func WarnNameCtx(c *gin.Context, name string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Warn(args...)
// }

// func DebugNameCtx(c *gin.Context, name string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Debug(args...)
// }

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

func InfofCtx(c *gin.Context, template string, args ...interface{}) {
	logger.Sugar().With("X-Request-Id", requestid.Get(c)).Infof(template, args...)
}

func ErrorfCtx(c *gin.Context, template string, args ...interface{}) {
	logger.Sugar().With("X-Request-Id", requestid.Get(c)).Errorf(template, args...)
}

func WarnfCtx(c *gin.Context, template string, args ...interface{}) {
	logger.Sugar().With("X-Request-Id", requestid.Get(c)).Warnf(template, args...)
}

func DebugfCtx(c *gin.Context, template string, args ...interface{}) {
	logger.Sugar().With("X-Request-Id", requestid.Get(c)).Debugf(template, args...)
}

// func InfofName(name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).Infof(template, args...)
// }

// func ErrorfName(name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).Errorf(template, args...)
// }

// func WarnfName(name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).Warnf(template, args...)
// }

// func DebugfName(name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).Debugf(template, args...)
// }

// func InfofNameCtx(c *gin.Context, name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Infof(template, args...)
// }

// func ErrorfNameCtx(c *gin.Context, name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Errorf(template, args...)
// }

// func WarnfNameCtx(c *gin.Context, name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Warnf(template, args...)
// }

// func DebugfNameCtx(c *gin.Context, name, template string, args ...interface{}) {
// 	logger.Sugar().Named(name).With("X-Request-Id", requestid.Get(c)).Debugf(template, args...)
// }

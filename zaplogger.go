package utilsgo

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var SugarLogger *zap.SugaredLogger

func InitLogger(logFileName, errFileName string) {
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, getLogWriter(logFileName), zap.ErrorLevel),
		zapcore.NewCore(encoder, getLogWriter(errFileName), zap.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zap.ErrorLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		//zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.WarnLevel),
	)

	logger := zap.New(core, zap.AddCaller())

	SugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./" + filename,
		MaxSize:    10,
		MaxBackups: 50,
		MaxAge:     60,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

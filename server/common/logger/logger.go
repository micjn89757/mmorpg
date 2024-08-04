package logger

import (
	"io"
	"os"
	"server/common"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger 		*zap.Logger
	loggerOnce 	sync.Once
)


func InitLogger(filename string) {
	loggerOnce.Do(func() {
		encoder := encoder()
		writer := writeSync(filename)
		core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

		Logger = zap.New(core)
	})
}


// 如何写入日志
func encoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写字母记录日志级别
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 写到哪里
func writeSync(filename string) zapcore.WriteSyncer {
	file, err := os.OpenFile(common.ProjectRootPath + "/log/" + filename, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		panic("file not exist")
	}

	writer := io.MultiWriter(os.Stderr, file)

	return zapcore.AddSync(writer)
}
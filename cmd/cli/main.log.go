package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterLog()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("This is an info message")
	logger.Error("This is an error message")
}

// getEncoderLog is used to create a zapcore.Encoder for logging.
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Set the time format to ISO8601
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Set time key to "time"
	encodeConfig.TimeKey = "ts"

	// Set the log level key to "level"
	encodeConfig.LevelKey = "level"

	// Set encode level to capital level
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Set encode caller to short caller
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterLog() zapcore.WriteSyncer {
	// Create a file writer for logging
	file, _ := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	// Create a multi-write syncer to write logs to both console and file
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

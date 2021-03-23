package main

import (
	"errors"
	"fmt"
	"go-tools/file"
	"go-tools/log"

	"github.com/zRedShift/mimemagic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = getLogger()

func getLogger() *zap.Logger {
	// l := log.NewLumberjackLogger{"log.log", 128, 30, 7, true}
	logger := log.NewLogger(nil, log.DefaultEncoderConfig, zapcore.DebugLevel, "json", "cloud")
	return logger
}

func function() {
	log.Info("msg")
}

func main() {
	logger.Debug("ssss")
	logger.Info("ssss")
	logger.Error("ssss")
	err := errors.New("test error")
	logger.Error("ssss", zap.Any("err", err.Error()))
	log.Debug("llllll")
	function()
	// filename := "/home/ubuntu/Projects/master.zip"
	dirname := "/home/ubuntu/Downloads"
	files, _ := file.GetAllFiles(dirname)
	for _, f := range files {
		mimeType, _ := mimemagic.MatchFilePath(f, -1)
		fmt.Println(mimeType.Subtype, mimeType.Extensions, file.IsCompressFile(f))
	}
	mimeType, _ := mimemagic.MatchFilePath("/home/ubuntu/Downloads/testdata/test1.tlz4", -1)
	fmt.Println("/home/ubuntu/Downloads/testdata/test1.tlz4", mimeType.Subtype, mimeType.Extensions)

	mimeType, _ = mimemagic.MatchFilePath("/home/ubuntu/Downloads/testdata.tgz", -1)
	fmt.Println("/home/ubuntu/Downloads/testdata.tgz", mimeType.Subtype, mimeType.Extensions)

	mimeType, _ = mimemagic.MatchFilePath("/home/ubuntu/Downloads/testdata/test4.tar.lz4", -1)
	fmt.Println("/home/ubuntu/Downloads/test4.tar.lz4", mimeType.Subtype, mimeType.Extensions)

}

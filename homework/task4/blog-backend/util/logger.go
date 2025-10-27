package util

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// 初始化日志
func InitLogger() {
	// 设置日志格式为 JSON 格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 创建日志目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		logrus.Fatalf("创建日志目录失败: %v", err)
	}

	// 设置日志文件
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logrus.Fatalf("打开日志文件失败: %v", err)
	}

	// 同时输出到控制台和文件
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(&fileHook{file: file})
}

// 文件钩子，将日志写入文件
type fileHook struct {
	file *os.File
}

func (hook *fileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.file.WriteString(line)
	return err
}

func (hook *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

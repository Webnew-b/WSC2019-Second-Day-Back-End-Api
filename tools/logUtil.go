package tools

import (
	"log"
	"os"
	"path/filepath"
)

var (
	Log *log.Logger
)

func GetWorkingDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前目录失败: %v", err)
	}
	return cwd
}

func makeLogDir(path string) string {
	logDir := filepath.Join(path, "log")
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("创建log目录失败: %v", err)
	}
	return logDir
}

func openLogFile(path string) *os.File {
	logFile := filepath.Join(path, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	return file
}

func closeLogger() {
	if file, ok := Log.Writer().(*os.File); ok {
		file.Close()
	}
}

func Close() {
	closeLogger()
}

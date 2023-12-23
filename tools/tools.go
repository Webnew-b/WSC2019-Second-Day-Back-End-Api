package tools

import (
	"io"
	"log"
	"os"
)

func createLogger() {
	log.Println("crating Log")
	cwd := GetWorkingDir()
	logDir := makeLogDir(cwd)
	file := openLogFile(logDir)

	multiWriter := io.MultiWriter(file, os.Stdout)
	Log = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("crate Log Success")
}

func init() {
	createLogger()
}

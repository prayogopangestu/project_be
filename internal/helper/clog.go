package helper

import (
	"log"
	"os"
)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func WriteError(errornya string) {
	fileError, err := OpenLogFile("./log_error.log")
	if err != nil {
		log.Fatal(err)
	}
	errorLog := log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Println(errornya)
}

func WriteInfo(infonya string) {
	fileInfo, err := OpenLogFile("./log_info.log")
	if err != nil {
		log.Fatal(err)
	}
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Println(infonya)
}

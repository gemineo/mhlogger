package mhlogger

import (
	"io"
	"log"
	"os"
	"os/user"
	fp "path/filepath"
)

var (
	// Info should be used for informative output
	Info *log.Logger

	// Warning should be used to raise attention
	Warning *log.Logger

	// Error should be used when a task is not doable
	Error *log.Logger
)

// Init initializes the logger
func Init(filepath string) (mw io.Writer, err error) {
	// If the path contains a folder structure, make sure it exists
	dir, _ := fp.Split(fp.Clean(filepath))
	if dir != "" {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return mw, err
		}
	}

	logFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	mw = io.MultiWriter(os.Stdout, logFile)

	username := "unknown user"
	user, err := user.Current()
	if err == nil {
		username = user.Username
	}

	Info = log.New(mw, "INFO: ["+username+"] - ", log.Ldate|log.Ltime)
	Warning = log.New(mw, "WARNING: ["+username+"] - ", log.Ldate|log.Ltime)
	Error = log.New(mw, "ERROR: ["+username+"] - ", log.Ldate|log.Ltime)

	return mw, err
}

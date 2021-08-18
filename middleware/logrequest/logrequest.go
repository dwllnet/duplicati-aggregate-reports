// Package logrequest provides an http.Handler that logs when a request is
// made to the application and lists the remote address, the HTTP method,
// and the URL.
package logrequest

import (
	"log"
	"net/http"
	"path/filepath"
	"time"
	"../../lib/flight"
	"github.com/kjk/dailyrotate"
)

var (
	logFile *dailyrotate.File
)

// Handler will log the HTTP requests.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		logFolder := c.Config.Server.LogsDirectory
		accessFile := c.Config.Server.AccessLog
		//errorFile := c.Config.Server.ErrorLog

		pathFormat := filepath.Join(logFolder, accessFile)
		err := openLogFile(pathFormat, onLogClose)
		if err != nil {
			log.Fatalf("openLogFile failed with '%s'\n", err)
		}
		defer closeLogFile()

		//fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)

		message := time.Now().Format("2006-01-02 03:04:05 PM") + " " + r.RemoteAddr + " " + r.Method + " " + r.URL.String() + "\n"
		err = writeToLog(message)
		if err != nil {
			log.Fatalf("writeToLog() failed with '%s'\n", err)
		}

		next.ServeHTTP(w, r)
	})
}

func openLogFile(pathFormat string, onClose func(string, bool)) error {
	w, err := dailyrotate.NewFile(pathFormat, onLogClose)
	if err != nil {
		return err
	}
	logFile = w
	return nil
}

func onLogClose(path string, didRotate bool) {
	//fmt.Printf("we just closed a file '%s', didRotate: %v\n", path, didRotate)
	if !didRotate {
		return
	}
	// process just closed file e.g. upload to backblaze storage for backup
	go func() {
		// if processing takes a long time, do it in background
	}()
}

func closeLogFile() error {
	return logFile.Close()
}

func writeToLog(msg string) error {
	_, err := logFile.Write([]byte(msg))
	return err
}





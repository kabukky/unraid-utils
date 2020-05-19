package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/kabukky/unraid-utils/cleaner"
	"github.com/kabukky/unraid-utils/scheduler"
	"github.com/kabukky/unraid-utils/utils"
)

func main() {
	if utils.Logfile != "" {
		logFile, err := os.OpenFile(utils.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Error opening log file", err)
		}
		defer logFile.Close()
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
		log.Println("Logging to file", utils.Logfile)
	}
	// Clean empty dirs
	if len(utils.DirsToClean) > 0 {
		go scheduler.Schedule(func() {
			cleaner.Clean(utils.DirsToClean, 24*time.Hour)
		}, 1*time.Hour)
	} else {
		log.Println("Please specify env var", utils.DirsToCleanEnvVar)
	}

	// Run forever
	select {}
}

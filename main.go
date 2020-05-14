package main

import (
	"log"
	"time"

	"github.com/kabukky/unraid-utils/cleaner"
	"github.com/kabukky/unraid-utils/scheduler"
	"github.com/kabukky/unraid-utils/utils"
)

func main() {

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

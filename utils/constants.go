package utils

import (
	"os"
	"strings"
)

var (
	// LogFileEnvVar is the env var for LogFile
	LogfileEnvVar = "LOGFILE"
	// DirsToCleanEnvVar is the env var for DirsToClean
	DirsToCleanEnvVar = "DIRS_TO_CLEAN"
	// LogFile is the path to the log file if set
	Logfile = os.Getenv(LogfileEnvVar)
	// DirsToClean is a list of all dirs that should be searched recursively for empty folders
	DirsToClean = readListEnvVar(DirsToCleanEnvVar)
)

func readListEnvVar(envVarName string) []string {
	var list []string
	envVar := os.Getenv(envVarName)
	if envVar != "" {
		for _, elem := range strings.Split(envVar, ",") {
			list = append(list, elem)
		}
	}
	return list
}

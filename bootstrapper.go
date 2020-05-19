package piutils

func BootStrap(logFile, errFile string) {
	InitLogger(logFile, errFile)
	defer SugarLogger.Sync()
}

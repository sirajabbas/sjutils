package utilsgo

func Configure(logFile, errFile string) {
	InitLogger(logFile, errFile)
	defer SugarLogger.Sync()
}

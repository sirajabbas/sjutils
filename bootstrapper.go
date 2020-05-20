package utilsgo

type Config struct {
	LogFileName   string
	ErrorFileName string
}

func Configure(c Config) {
	InitLogger(c.LogFileName, c.ErrorFileName)
	defer SugarLogger.Sync()
}

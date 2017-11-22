package config

func GetConfig() ServiceConfig {
	return Config
}

func WriteConfig() (error, bool) {
	return nil, true
}

func ReloadConfig() (error, bool) {
	return nil, true
}

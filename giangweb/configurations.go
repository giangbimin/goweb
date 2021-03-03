package giangweb

// configDatabase from env
type ConfigDatabase struct {
	AppName  string `env:APP_NAME env-default:"giangweb"`
	AppEnv   string `env:APP_ENV env-default:"development"`
	Port     string `env:GO_PORT env-default:"8080"`
	Host     string `env:GO_HOST env-default:"localhost"`
	LogLevel string `env:LOG_LEVEL env-default:"INFO"`
}

var cfg ConfigDatabase

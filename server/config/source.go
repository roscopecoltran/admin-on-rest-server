package config

type SourceConfig struct {
	Debug      bool
	Disabled   bool
	Connection ConnectionConfig
	Collection CollectionConfig
}

type ConnectionConfig struct {
	Name          string
	Adapter       string
	Host          string
	Port          string
	User          string
	Password      string
	PrefixPath    string
	MigrationsDir string
	Secured       bool
	SingularTable bool
	Debug         bool
	DSN           string
}

type CollectionConfig struct {
	Debug    bool
	Disabled bool
	Table    string
}

package router

type ServerMode string

const (
	Debug       ServerMode = "debug"
	Local       ServerMode = "local"
	Development ServerMode = "development"
	Production  ServerMode = "production"
)

type Config struct {
	Mode    ServerMode
	Version string
}

func New(mode ServerMode, version string) *Config {
	return &Config{
		Mode:    mode,
		Version: version,
	}
}

func NewDebug(version string) *Config {
	return &Config{
		Mode:    Debug,
		Version: version,
	}
}

func NewLocal(version string) *Config {
	return &Config{
		Mode:    Local,
		Version: version,
	}
}

func NewDevelopment(version string) *Config {
	return &Config{
		Mode:    Development,
		Version: version,
	}
}

func NewProduction(version string) *Config {
	return &Config{
		Mode:    Production,
		Version: version,
	}
}

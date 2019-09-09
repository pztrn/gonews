package configuration

// Represents configuration file structure.
type config struct {
	// Database represents database configuration.
	Database Database `yaml:"database"`
	// Network represents network stack configuration.
	Network []Network `yaml:"network"`
}

type Database struct {
	// DSN is a connection string in form "proto://user:password@host:port/database".
	DSN string `yaml:"dsn"`
	// Params is a DSN params that will be added to DSN after
	// connection string itself.
	Params string `yaml:"params"`
	// Timeout is a timeout for connection watcher. Every this
	// count of seconds connection watcher will check if database
	// connection is alive.
	Timeout int64 `yaml:"timeout"`
}

type Network struct {
	// Address represents address to bing in form of "ip:port".
	Address string `yaml:"address"`
	// Limit sets maximum simultaneous connections that can be
	// processed by worker.
	Limit int `yaml:"limit"`
	// Type sets connection type. See networker for available
	// types.
	Type string `yaml:"type"`
}

package configuration

// Represents configuration file structure.
type config struct {
	// Network represents network stack configuration.
	Network []Network `yaml:"network"`
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

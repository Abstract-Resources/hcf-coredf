package config

// ServerConfig is an extension of the Dragonfly server config to include fields specific to Vasar.
type ServerConfig struct {
	// Server contains fields specific to the server like Whitelisted or something like that.
	Server struct {
		// Whitelisted is true if the server is whitelisted.
		Whitelisted bool
	}
	// Provider contains fields related to the database provider like mongodb or mysql
	Provider struct {
		// ProviderName is the type of provider (MongoDB or MySQL)
		ProviderName string
		// Provider address, to use port add ":" like "127.0.0.1:3306"
		Address string
		// Database username used to login
		Username string
		// Database password used to login
		Password string
		// Database name
		Dbname string
	}
}

// DefaultConfig returns a default config for the server.
func DefaultConfig() ServerConfig {
	c := ServerConfig{}

	c.Server.Whitelisted = true

	c.Provider.ProviderName = "MongoDB"
	c.Provider.Address = "127.0.0.1"
	c.Provider.Username = "admin"
	c.Provider.Password = ""
	c.Provider.Dbname = "hcf_core"

	return c
}

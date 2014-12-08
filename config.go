package marathon

import (
	"github.com/Banno/go-marathon"
)

type Config struct {
	Host string
	Port int

	client *marathon.Client
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	c.client = marathon.NewClient(c.Host, c.Port)

	return nil
}

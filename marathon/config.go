package marathon

import (
	"github.com/Banno/go-marathon"
)

type Config struct {
	Url string

	client *marathon.Client
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	c.client = marathon.NewClientForUrl(c.Url)

	return nil
}

package marathon

import (
	"github.com/gambol99/go-marathon"
	"os"
)

type Config struct {
	Url string

	client *marathon.Client
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	config := marathon.NewDefaultConfig()
	config.URL = c.Url
	config.LogOutput = os.Stdout

	client, err := marathon.NewClient(config)
	c.client = client.(*marathon.Client)
	return err
}

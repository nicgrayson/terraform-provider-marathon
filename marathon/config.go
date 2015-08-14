package marathon

import (
	"github.com/gambol99/go-marathon"
	"os"
	"time"
)

type Config struct {
	Url                      string
	RequestTimeout           int
	DefaultDeploymentTimeout time.Duration

	client marathon.Marathon
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	config := marathon.NewDefaultConfig()
	config.URL = c.Url
	config.RequestTimeout = c.RequestTimeout
	config.DefaultDeploymentTimeout = c.DefaultDeploymentTimeout
	config.LogOutput = os.Stdout

	client, err := marathon.NewClient(config)
	c.client = client
	return err
}

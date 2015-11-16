package marathon

import (
	"github.com/gambol99/go-marathon"
	"time"
)

type Config struct {
	Url                      string
	RequestTimeout           int
	DefaultDeploymentTimeout time.Duration

	Client marathon.Marathon
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	marathonConfig := marathon.NewDefaultConfig()
	marathonConfig.URL = c.Url
	marathonConfig.RequestTimeout = c.RequestTimeout

	client, err := marathon.NewClient(marathonConfig)
	c.Client = client
	return err
}

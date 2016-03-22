package marathon

import (
	"github.com/gambol99/go-marathon"
	"net/http"
	"time"
)

type config struct {
	URL                      string
	RequestTimeout           int
	DefaultDeploymentTimeout time.Duration
	BasicAuthUser            string
	BasicAuthPassword        string

	Client marathon.Marathon
}

func (c *config) loadAndValidate() error {

	// this needs to return an err as well.
	marathonConfig := marathon.NewDefaultConfig()
	marathonConfig.URL = c.URL
	marathonConfig.HTTPClient = &http.Client{
		Timeout: time.Duration(c.RequestTimeout) * time.Second,
	}
	marathonConfig.HTTPBasicAuthUser = c.BasicAuthUser
	marathonConfig.HTTPBasicPassword = c.BasicAuthPassword

	client, err := marathon.NewClient(marathonConfig)
	c.Client = client
	return err
}

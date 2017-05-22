package main

import (
	"github.com/nicgrayson/terraform-provider-marathon/marathon"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: marathon.Provider,
	})
}

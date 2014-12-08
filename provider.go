package marathon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Marathon's Hostname or IP",
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8080,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"marathon_app": resourceMarathonApp(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host: d.Get("host").(string),
		Port: d.Get("port").(int),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return config.client, nil
}

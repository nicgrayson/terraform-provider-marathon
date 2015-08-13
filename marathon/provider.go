package marathon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MARATHON_URL", nil),
				Description: "Marathon's Base HTTP URL",
			},
			"request_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "'Request Timeout",
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
		Url:            d.Get("url").(string),
		RequestTimeout: d.Get("request_timeout").(int),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return config.client, nil
}

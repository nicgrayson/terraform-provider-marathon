package marathon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"time"
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
			"deployment_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     600,
				Description: "'Deployment Timeout",
			},
			"basic_auth_user": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTP basic auth user",
			},
			"basic_auth_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTP basic auth password",
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
		Url:                      d.Get("url").(string),
		RequestTimeout:           d.Get("request_timeout").(int),
		DefaultDeploymentTimeout: time.Duration(d.Get("deployment_timeout").(int)) * time.Second,
		BasicAuthUser:            d.Get("basic_auth_user").(string),
		BasicAuthPassword:        d.Get("basic_auth_password").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return config, nil
}

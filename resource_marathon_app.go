package marathon

import (
	"github.com/Banno/go-marathon"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
)

func resourceMarathonApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMarathonAppCreate,
		Read:   resourceMarathonAppRead,
		//		Update: resourceMarathonAppUpdate,
		Delete: resourceMarathonAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"container": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	// terraform doesn't have a float type yet
	cpus, _ := strconv.ParseFloat(d.Get("cpus").(string), 64)
	mem, _ := strconv.ParseFloat(d.Get("mem").(string), 64)

	appMutable := marathon.AppMutable{
		Id:   d.Get("name").(string),
		Cpus: cpus,
		Mem:  mem,
		Container: &marathon.Container{
			Docker: &marathon.Docker{
				Image: d.Get("container.0.docker.0.image").(string),
			},
			Type: d.Get("container.0.type").(string),
		},
	}

	app, err := c.AppCreate(appMutable)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(app.Id)

	return resourceMarathonAppRead(d, meta)
}

func resourceMarathonAppRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	// client should throw error if id is nil
	app, err := c.AppRead(d.Id())

	if app.Id == "" {
		d.SetId("")
	}

	// Add in computed values from App struct here.

	return nil
}

/*
// test mutating existing state of AppMutable fields

func resourceMarathonAppUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceMarathonAppCreate(d, meta)
}
*/

func resourceMarathonAppDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	if err := c.AppDelete(d.Id()); err != nil {
		return err
	}

	return nil
}

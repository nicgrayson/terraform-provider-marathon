package marathon

import (
	"github.com/Banno/go-marathon"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
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
			"name": &schema.Schema{ // represents 'id' field
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"args": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backoff_seconds": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"backoff_factor": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"cmd": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"constraints": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"constraint": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"operation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"parameter": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
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
									"network": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_mappings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_mapping": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: false,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"container_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"host_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"service_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"volumes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"container_path": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"host_path": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"mode": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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
			"cpus": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"dependencies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disk": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"health_checks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"path": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"grace_period_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interval_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"port_index": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"timeout_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"max_consecutive_failures": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"command": &schema.Schema{
										Type:     schema.TypeMap,
										Optional: true,
									},
									// incomplete computed values here
								},
							},
						},
					},
				},
			},
			"instances": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"upgrade_strategy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"minimum_health_capacity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"uris": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] >>>>>>>>>>>>>>>>>> ENTERING AppCreate")

	c := meta.(*marathon.Client)

	appMutable := marathon.AppMutable{}

	if v, ok := d.GetOk("name"); ok {
		appMutable.Id = v.(string)
	}

	if v, ok := d.GetOk("args.#"); ok {
		args := make([]string, v.(int))

		for i, _ := range args {
			args[i] = d.Get("args." + strconv.Itoa(i)).(string)
		}

		if len(args) != 0 {
			appMutable.Args = args
		}
	}

	if v, ok := d.GetOk("backoff_seconds"); ok {
		backoffSeconds, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.BackoffSeconds = backoffSeconds
	}

	if v, ok := d.GetOk("backoff_factor"); ok {
		backoffFactor, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.BackoffFactor = backoffFactor
	}

	if v, ok := d.GetOk("cmd"); ok {
		appMutable.Cmd = v.(string)
	}

	if v, ok := d.GetOk("constraints.0.constraint.#"); ok {
		constraints := make([][]string, v.(int))

		for i, _ := range constraints {
			cMap := d.Get(fmt.Sprintf("constraints.0.constraint.%d", i)).(map[string]interface{})

			if cMap["parameter"] == "" {
				constraints[i] = make([]string, 2)
				constraints[i][0] = cMap["attribute"].(string)
				constraints[i][1] = cMap["operation"].(string)
			} else {
				constraints[i] = make([]string, 3)
				constraints[i][0] = cMap["attribute"].(string)
				constraints[i][1] = cMap["operation"].(string)
				constraints[i][2] = cMap["parameter"].(string)
			}
		}

		appMutable.Constraints = constraints
	}

	// Container structure -- certainly not complete.
	appMutable.Container = &marathon.Container{
		Docker: &marathon.Docker{
			Image: d.Get("container.0.docker.0.image").(string),
		},
		Type: d.Get("container.0.type").(string),
	}

	if v, ok := d.GetOk("cpus"); ok {
		cpus, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Cpus = cpus
	}

	if v, ok := d.GetOk("dependencies.#"); ok {
		dependencies := make([]string, v.(int))

		for i, _ := range dependencies {
			dependencies[i] = d.Get("dependencies." + strconv.Itoa(i)).(string)
		}

		if len(dependencies) != 0 {
			appMutable.Dependencies = dependencies
		}
	}

	if v, ok := d.GetOk("disk"); ok {
		disk, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Disk = disk
	}

	if v, ok := d.GetOk("env"); ok {
		envMap := v.(map[string]interface{})
		env := make(map[string]string, len(envMap))

		for k, v := range envMap {
			env[k] = v.(string)
		}

		appMutable.Env = env
	}

	if v, ok := d.GetOk("health_checks.0.health_check.#"); ok {
		healthChecks := make([]marathon.HealthCheck, v.(int))

		for i, _ := range healthChecks {
			healthCheck := marathon.HealthCheck{}
			mapStruct := d.Get("health_checks.0.health_check." + strconv.Itoa(i)).(map[string]interface{})

			if prop, ok := mapStruct["grace_period_seconds"]; ok {
				healthCheck.GracePeriodSeconds = prop.(int)
			}

			if prop, ok := mapStruct["interval_seconds"]; ok {
				healthCheck.IntervalSeconds = prop.(int)
			}

			if prop, ok := mapStruct["max_consecutive_failures"]; ok {
				healthCheck.MaxConsecutiveFailures = prop.(int)
			}

			if prop, ok := mapStruct["path"]; ok {
				healthCheck.Path = prop.(string)
			}

			if prop, ok := mapStruct["port_index"]; ok {
				healthCheck.PortIndex = prop.(int)
			}

			if prop, ok := mapStruct["protocol"]; ok {
				healthCheck.Protocol = prop.(string)
			}

			if prop, ok := mapStruct["timeout_seconds"]; ok {
				healthCheck.TimeoutSeconds = prop.(int)
			}

			// config

			log.Printf("=====\n%#v\n", healthCheck)
			healthChecks[i] = healthCheck
		}

		appMutable.HealthChecks = healthChecks
	}

	if v, ok := d.GetOk("instances"); ok {
		appMutable.Instances = v.(int)
	}

	if v, ok := d.GetOk("mem"); ok {
		mem, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Mem = mem
	}

	if v, ok := d.GetOk("ports.#"); ok {
		ports := make([]int, v.(int))

		for i, _ := range ports {
			ports[i] = d.Get("ports." + strconv.Itoa(i)).(int)
		}

		if len(ports) != 0 {
			appMutable.Ports = ports
		}
	}

	if v, ok := d.GetOk("upgrade_strategy.minimum_health_capacity"); ok {
		capacity, _ := strconv.ParseFloat(v.(string), 64)

		upgradeStrategy := &marathon.UpgradeStrategy{
			MinimumHealthCapacity: capacity,
		}
		appMutable.UpgradeStrategy = upgradeStrategy

	}

	if v, ok := d.GetOk("uris.#"); ok {
		uris := make([]string, v.(int))

		for i, _ := range uris {
			uris[i] = d.Get("uris." + strconv.Itoa(i)).(string)
		}

		if len(uris) != 0 {
			appMutable.Uris = uris
		}
	}

	log.Printf("=====\n%#v\n", appMutable)

	app, err := c.AppCreate(appMutable)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(app.Id)
	d.Set("version", app.Version)

	return resourceMarathonAppRead(d, meta)
}

func resourceMarathonAppRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	// client should throw error if id is nil
	app, _ := c.AppRead(d.Id())

	log.Printf("== READ ==\n%#v\n", app)

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

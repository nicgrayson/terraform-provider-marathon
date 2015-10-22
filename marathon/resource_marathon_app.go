package marathon

import (
	"fmt"
	"github.com/gambol99/go-marathon"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceMarathonApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMarathonAppCreate,
		Read:   resourceMarathonAppRead,
		Update: resourceMarathonAppUpdate,
		Delete: resourceMarathonAppDelete,

		Schema: map[string]*schema.Schema{
			"accepted_resource_roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"app_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: false,
				Default:  1,
			},
			"backoff_factor": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: false,
				Default:  1.15,
			},
			"cmd": &schema.Schema{
				Type:     schema.TypeString,
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
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"force_pull_image": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"image": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"network": &schema.Schema{
										Type:     schema.TypeString,
										Default:  "HOST",
										Optional: true,
									},
									"privileged": &schema.Schema{
										Type:     schema.TypeBool,
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
																Default:  "tcp",
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
							Default:  "DOCKER",
						},
					},
				},
			},
			"cpus": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  0.1,
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
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
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
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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
				Default:  1,
				ForceNew: false,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  128,
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
			"require_ports": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"upgrade_strategy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"minimum_health_capacity": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
							Default:  1.0,
						},
						"maximum_over_capacity": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
							Default:  1.0,
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
			// many other "computed" values haven't been added.
		},
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	client := config.Client

	application := mutateResourceToApplication(d)

	application, err := client.CreateApplication(application)
	if err != nil {
		log.Println("[ERROR] creating application", err)
		return err
	}
	d.Partial(true)
	d.SetId(application.ID)
	setSchemaFieldsForApp(application, d)

	err = client.WaitOnApplication(application.ID, config.DefaultDeploymentTimeout)
	if err != nil {
		log.Println("[ERROR] waiting for application", err)
		return err
	}
	d.Partial(false)

	return resourceMarathonAppRead(d, meta)
}

func resourceMarathonAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	client := config.Client

	app, err := client.Application(d.Id())

	if err != nil {
		// Handle a deleted app
		if err == marathon.ErrDoesNotExist {
			d.SetId("")
			return nil
		}
		return err
	}

	if app.ID == "" {
		d.SetId("")
	}

	setSchemaFieldsForApp(app, d)

	return nil
}

func setSchemaFieldsForApp(app *marathon.Application, d *schema.ResourceData) {
	d.Set("app_id", app.ID)
	d.SetPartial("app_id")

	d.Set("accepted_resource_roles", app.AcceptedResourceRoles)
	d.SetPartial("accepted_resource_roles")

	d.Set("args", app.Args)
	d.SetPartial("args")

	d.Set("backoff_seconds", app.BackoffSeconds)
	d.SetPartial("backoff_seconds")

	d.Set("backoff_factor", app.BackoffFactor)
	d.SetPartial("backoff_factor")

	d.Set("cmd", app.Cmd)
	d.SetPartial("cmd")

	if len(app.Constraints) > 0 {
		cMaps := make([]map[string]string, len(app.Constraints))
		for idx, constraint := range app.Constraints {
			cMap := make(map[string]string)
			cMap["attribute"] = constraint[0]
			cMap["operation"] = constraint[1]
			if len(constraint) > 2 {
				cMap["parameter"] = constraint[2]
			}
			cMaps[idx] = cMap
		}
		d.Set("constraints", []interface{}{map[string]interface{}{"constraint": cMaps}})
	} else {
		d.Set("constraints", make([]interface{}, 0))
	}
	d.SetPartial("constraints")

	if app.Container != nil {
		container := app.Container

		containerMap := make(map[string]interface{})
		containerMap["type"] = container.Type

		if container.Type == "DOCKER" {
			docker := container.Docker
			dockerMap := make(map[string]interface{})
			containerMap["docker"] = []interface{}{dockerMap}

			dockerMap["image"] = docker.Image
			dockerMap["force_pull_image"] = docker.ForcePullImage
			dockerMap["network"] = docker.Network
			dockerMap["privileged"] = docker.Privileged

			if len(docker.PortMappings) > 0 {
				portMappings := make([]map[string]interface{}, len(docker.PortMappings))
				for idx, portMapping := range docker.PortMappings {
					pmMap := make(map[string]interface{})
					pmMap["container_port"] = portMapping.ContainerPort
					pmMap["host_port"] = portMapping.HostPort
					// pmMap["service_port"] = portMapping.ServicePort
					pmMap["protocol"] = portMapping.Protocol
					portMappings[idx] = pmMap
				}
				dockerMap["port_mappings"] = []interface{}{map[string]interface{}{"port_mapping": portMappings}}
			} else {
				dockerMap["port_mappings"] = make([]interface{}, 0)
			}

		}

		if len(container.Volumes) > 0 {
			volumes := make([]map[string]interface{}, len(container.Volumes))
			for idx, volume := range container.Volumes {
				volumeMap := make(map[string]interface{})
				volumeMap["container_path"] = volume.ContainerPath
				volumeMap["host_path"] = volume.HostPath
				volumeMap["mode"] = volume.Mode
				volumes[idx] = volumeMap
			}
			containerMap["volumes"] = []interface{}{map[string]interface{}{"volume": volumes}}
		} else {
			containerMap["volumes"] = make([]interface{}, 0)
		}

		d.Set("container", []interface{}{containerMap})
	}
	d.SetPartial("container")

	d.Set("cpus", app.CPUs)
	d.SetPartial("cpus")

	d.Set("dependencies", app.Dependencies)
	d.SetPartial("dependencies")

	d.Set("env", app.Env)
	d.SetPartial("env")

	if len(app.HealthChecks) > 0 {
		healthChecks := make([]map[string]interface{}, len(app.HealthChecks))
		for idx, healthCheck := range app.HealthChecks {
			hMap := make(map[string]interface{})
			if healthCheck.Command != nil {
				hMap["command"] = map[string]string{"value": healthCheck.Command.Value}
			}
			hMap["grace_period_seconds"] = healthCheck.GracePeriodSeconds
			hMap["interval_seconds"] = healthCheck.IntervalSeconds
			hMap["max_consecutive_failures"] = healthCheck.MaxConsecutiveFailures
			hMap["path"] = healthCheck.Path
			hMap["port_index"] = healthCheck.PortIndex
			hMap["timeout_seconds"] = healthCheck.TimeoutSeconds
			healthChecks[idx] = hMap
		}
		d.Set("health_checks", []interface{}{map[string]interface{}{"health_check": healthChecks}})
	} else {
		d.Set("health_checks", make([]interface{}, 0))
	}
	d.SetPartial("health_checks")

	d.Set("instances", app.Instances)
	d.SetPartial("instances")

	d.Set("mem", app.Mem)
	d.SetPartial("mem")

	if givenFreePortsDoesNotEqualAllocated(d, app) {
		d.Set("ports", app.Ports)
	}
	d.SetPartial("ports")

	d.Set("require_ports", app.RequirePorts)
	d.SetPartial("require_ports")

	if app.UpgradeStrategy != nil {
		usMap := make(map[string]interface{})
		usMap["minimum_health_capacity"] = app.UpgradeStrategy.MinimumHealthCapacity
		usMap["maximum_over_capacity"] = app.UpgradeStrategy.MaximumOverCapacity
		d.Set("upgrade_strategy", []interface{}{usMap})
	} else {
		d.Set("upgrade_strategy", make([]interface{}, 0))
	}
	d.SetPartial("upgrade_strategy")

	d.Set("uris", app.Uris)
	d.SetPartial("uris")

	// App
	d.Set("executor", app.Executor)
	d.SetPartial("executor")

	d.Set("disk", app.Disk)
	d.SetPartial("disk")

	d.Set("user", app.User)
	d.SetPartial("user")

	d.Set("version", app.Version)
	d.SetPartial("version")
}

func givenFreePortsDoesNotEqualAllocated(d *schema.ResourceData, app *marathon.Application) bool {
	marathonPorts := make([]int, len(app.Ports))
	for i, port := range app.Ports {
		if port >= 10000 && port <= 20000 {
			marathonPorts[i] = 0
		} else {
			marathonPorts[i] = port
		}
	}

	ports := getPorts(d)

	return !reflect.DeepEqual(marathonPorts, ports)
}

func resourceMarathonAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	client := config.Client

	application := mutateResourceToApplication(d)

	deploymentId, err := client.UpdateApplication(application)
	if err != nil {
		return err
	}

	err = client.WaitOnDeployment(deploymentId.DeploymentID, config.DefaultDeploymentTimeout)
	return err
}

func resourceMarathonAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	client := config.Client

	deploymentId, err := client.DeleteApplication(d.Id())
	if err != nil {
		return err
	}
	err = client.WaitOnDeployment(deploymentId.DeploymentID, config.DefaultDeploymentTimeout)
	return err
}

func mutateResourceToApplication(d *schema.ResourceData) *marathon.Application {

	application := new(marathon.Application)

	if v, ok := d.GetOk("accepted_resource_roles.#"); ok {
		accepted_resource_roles := make([]string, v.(int))

		for i, _ := range accepted_resource_roles {
			accepted_resource_roles[i] = d.Get("accepted_resource_roles." + strconv.Itoa(i)).(string)
		}

		if len(accepted_resource_roles) != 0 {
			application.AcceptedResourceRoles = accepted_resource_roles
		}
	}

	if v, ok := d.GetOk("app_id"); ok {
		application.ID = v.(string)
	}

	if v, ok := d.GetOk("args.#"); ok {
		args := make([]string, v.(int))

		for i, _ := range args {
			args[i] = d.Get("args." + strconv.Itoa(i)).(string)
		}

		if len(args) != 0 {
			application.Args = args
		}
	}

	if v, ok := d.GetOk("backoff_seconds"); ok {
		application.BackoffSeconds = v.(float64)
	}

	if v, ok := d.GetOk("backoff_factor"); ok {
		application.BackoffFactor = v.(float64)
	}

	if v, ok := d.GetOk("cmd"); ok {
		application.Cmd = v.(string)
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

		application.Constraints = constraints
	} else {
		application.Constraints = make([][]string, 0)
	}

	if v, ok := d.GetOk("container.0.type"); ok {
		container := new(marathon.Container)
		t := v.(string)

		container.Type = t

		if t == "DOCKER" {
			docker := new(marathon.Docker)

			if v, ok := d.GetOk("container.0.docker.0.image"); ok {
				docker.Image = v.(string)
			}

			if v, ok := d.GetOk("container.0.docker.0.force_pull_image"); ok {
				docker.ForcePullImage = v.(bool)
			}

			if v, ok := d.GetOk("container.0.docker.0.network"); ok {
				docker.Network = v.(string)
			}

			if v, ok := d.GetOk("container.0.docker.0.privileged"); ok {
				docker.Privileged = v.(bool)
			}

			if v, ok := d.GetOk("container.0.docker.0.port_mappings.0.port_mapping.#"); ok {
				portMappings := make([]*marathon.PortMapping, v.(int))

				for i, _ := range portMappings {
					portMappings[i] = new(marathon.PortMapping)

					pmMap := d.Get(fmt.Sprintf("container.0.docker.0.port_mappings.0.port_mapping.%d", i)).(map[string]interface{})

					if val, ok := pmMap["container_port"]; ok {
						portMappings[i].ContainerPort = val.(int)
					}
					if val, ok := pmMap["host_port"]; ok {
						portMappings[i].HostPort = val.(int)
					}
					if val, ok := pmMap["protocol"]; ok {
						portMappings[i].Protocol = val.(string)
					}
					if val, ok := pmMap["service_port"]; ok {
						portMappings[i].ServicePort = val.(int)
					}

				}
				docker.PortMappings = portMappings
			}
			container.Docker = docker

		}

		if v, ok := d.GetOk("container.0.volumes.0.volume.#"); ok {
			volumes := make([]*marathon.Volume, v.(int))

			for i, _ := range volumes {
				volumes[i] = new(marathon.Volume)

				volumeMap := d.Get(fmt.Sprintf("container.0.volumes.0.volume.%d", i)).(map[string]interface{})

				if val, ok := volumeMap["container_path"]; ok {
					volumes[i].ContainerPath = val.(string)
				}
				if val, ok := volumeMap["host_path"]; ok {
					volumes[i].HostPath = val.(string)
				}
				if val, ok := volumeMap["mode"]; ok {
					volumes[i].Mode = val.(string)
				}
			}
			container.Volumes = volumes
		}

		application.Container = container
	}

	if v, ok := d.GetOk("cpus"); ok {
		application.CPUs = v.(float64)
	}

	if v, ok := d.GetOk("dependencies.#"); ok {
		dependencies := make([]string, v.(int))

		for i, _ := range dependencies {
			dependencies[i] = d.Get("dependencies." + strconv.Itoa(i)).(string)
		}

		if len(dependencies) != 0 {
			application.Dependencies = dependencies
		}
	}

	if v, ok := d.GetOk("env"); ok {
		envMap := v.(map[string]interface{})
		env := make(map[string]string, len(envMap))

		for k, v := range envMap {
			env[k] = v.(string)
		}

		application.Env = env
	} else {
		application.Env = make(map[string]string, 0)
	}

	if v, ok := d.GetOk("health_checks.0.health_check.#"); ok {
		healthChecks := make([]*marathon.HealthCheck, v.(int))

		for i, _ := range healthChecks {
			healthCheck := new(marathon.HealthCheck)
			mapStruct := d.Get("health_checks.0.health_check." + strconv.Itoa(i)).(map[string]interface{})

			if prop, ok := d.GetOk("health_checks.0.health_check." + strconv.Itoa(i) + ".command.value"); ok {
				healthCheck.Command = &marathon.Command{Value: prop.(string)}
				healthCheck.Protocol = "COMMAND"
			}

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

			healthChecks[i] = healthCheck
		}

		application.HealthChecks = healthChecks
	}

	if v, ok := d.GetOk("instances"); ok {
		application.Instances = v.(int)
	}

	if v, ok := d.GetOk("mem"); ok {
		application.Mem = v.(float64)
	}

	if v, ok := d.GetOk("require_ports"); ok {
		application.RequirePorts = v.(bool)
	}

	application.Ports = getPorts(d)

	upgradeStrategy := &marathon.UpgradeStrategy{}

	if v, ok := d.GetOk("upgrade_strategy.0.minimum_health_capacity"); ok {
		upgradeStrategy.MinimumHealthCapacity = v.(float64)
	}

	if v, ok := d.GetOk("upgrade_strategy.0.maximum_over_capacity"); ok {
		upgradeStrategy.MaximumOverCapacity = v.(float64)
	}

	application.UpgradeStrategy = upgradeStrategy

	if v, ok := d.GetOk("uris.#"); ok {
		uris := make([]string, v.(int))

		for i, _ := range uris {
			uris[i] = d.Get("uris." + strconv.Itoa(i)).(string)
		}

		if len(uris) != 0 {
			application.Uris = uris
		}
	}

	return application
}

func getPorts(d *schema.ResourceData) []int {
	ports := make([]int, 0)
	if v, ok := d.GetOk("ports.#"); ok {
		ports = make([]int, v.(int))

		for i, _ := range ports {
			ports[i] = d.Get("ports." + strconv.Itoa(i)).(int)
		}
	}
	return ports
}

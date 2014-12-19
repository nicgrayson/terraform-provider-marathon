// use terraform.tfvars
variable "marathon_host" {}
variable "marathon_port" {}

provider "marathon" {
  host = "${var.marathon_host}"
  port = "${var.marathon_port}"
}

resource "marathon_app" "app-create-example" {
	name = "app-create-example"

	cmd = "env && python3 -m http.server $PORT0"
	
	constraints {
		constraint {
			attribute = "hostname"
			operation = "UNIQUE"
		}
		constraint {
			attribute = "hostname"
			operation = "UNIQUE"
			parameter = "test"
		}
	}

	container {
		// finish this
		type = "DOCKER"
		docker {
			image = "python:3"
			network = "BRIDGE"
			port_mappings {
				port_mapping {
					container_port = 8080
					host_port = 0
					service_port = 9000
					protocol = "tcp"
				}
				port_mapping {
					container_port = 161
					host_port = 0
					protocol = "udp"
				}
			}
		}

		volumes {
			volume {
				container_path = "/etc/a"
				host_path = "/var/data/a"
				mode = "RO"
			}
			volume {
				container_path = "/etc/b"
				host_path = "/var/data/b"
				mode = "RW"
			}
		}
	}

	cpus = "0.01"

	dependencies = ["/test"]

	env {
		TEST = "hey"
		OTHER_TEST = "nope"
	}
	
	health_checks {
		health_check {
			grace_period_seconds = 3
			interval_seconds = 10
			max_consecutive_failures = 3
			path = "/"
			port_index = 0
			protocol = "HTTP"
			timeout_seconds = 5
		}
		health_check {
			command { 
				value = "curl -f -X GET http://$HOST:$PORT0/"
			}
			max_consecutive_failures = 3
			protocol = "COMMAND"
		}
	}

	instances = 2
	mem = 50
	ports = [0]
	
	upgrade_strategy {
		minimum_health_capacity = "0.5"
	}

}

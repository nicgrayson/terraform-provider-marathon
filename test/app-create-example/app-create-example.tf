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
		docker {
			image = "python:3"
		}
		type = "DOCKER"
	}

	cpus = "0.25"
	
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
	}

	instances = 2
	mem = 50
	ports = [0]
	
	upgrade_strategy {
		minimum_health_capacity = "0.5"
	}

}

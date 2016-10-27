resource "marathon_app" "app-create-example" {
  app_id = "/app-create-example"

  cmd = "env && python3 -m http.server $PORT0"

  constraints {
    constraint {
      attribute = "hostname"
      operation = "UNIQUE"
    }
  }

  container {
    docker {
      image = "python:3"
      network = "BRIDGE"
      parameters {
        parameter {
          key = "hostname"
          value = "a.corp.org"
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

  env {
    TEST = "hey"
    OTHER_TEST = "nope"
  }

  health_checks {
    health_check {
      command {
        value = "curl -f -X GET http://$HOST:$PORT0/"
      }
      max_consecutive_failures = 0
      protocol = "COMMAND"
    }
  }

  instances = 1
  labels {
    test = "abc"
  }
  mem = 50
  ports = [0]

  upgrade_strategy {
    minimum_health_capacity = "0.5"
  }
}

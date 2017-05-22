resource "marathon_app" "ip-address-create-example" {
  app_id = "/app-create-example"
  cmd = "env && python3 -m http.server 8080"
  cpus = 0.01
  instances = 1
  mem = 50

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
      port_mappings {
        port_mapping {
          container_port = 8080
          host_port = 0
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

  env {
    TEST = "hey"
    OTHER_TEST = "nope"
  }

  health_checks {
     health_check {
       command {
         value = "ps aux |grep python"
       }
       max_consecutive_failures = 0
       protocol = "COMMAND"
     }
  }

  ipaddress {
    network_name = "default"
  }

  labels {
    test = "abc"
  }

  upgrade_strategy {
    minimum_health_capacity = "0.5"
  }
}

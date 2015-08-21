# Marathon Terraform Provider

## Install
```
$ go get github.com/Banno/terraform-provider-marathon
```

## Usage

### Provider Configuration
Use a [tfvar file](https://www.terraform.io/intro/getting-started/variables.html) or set the ENV variable

```bash
$ export TF_VAR_marathon_url="http://marthon.domain.tld:8080"
```

```hcl
variable "marathon_url" {}

provider "marathon" {
  url = "${var.marathon_url}"
}
```

### Basic Usage
```hcl
resource "marathon_app" "hello-world" {
  app_id= "/hello-world"
  cmd = "echo 'hello'; sleep 10000"
  cpus = 0.01
  instances = 1
  mem = 16
  ports = [0]
}
```

### Docker Usage
```hcl
resource "marathon_app" "docker-hello-world" {
  app_id = "/docker-hello-world"
  container {
    docker {
      image = "hello-world"
    }
  }
  cpus = 0.01
  instances = 1
  mem = 16
  ports = [0]
}
```

### Full Example

[terraform file](test/app-create-example.tf)

## Development

### Build
```bash
$ go install
```

### Test
```bash
$ export MARATHON_URL="http://marthon.domain.tld:8080"
$ ./test.sh
```

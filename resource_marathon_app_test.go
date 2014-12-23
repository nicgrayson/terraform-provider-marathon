package marathon

import (
	"github.com/Banno/go-marathon"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"testing"
)

func TestAccMarathonApp_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMarathonAppDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckMarathonAppConfig_basic,
				Check:  resource.ComposeTestCheckFunc(
				// fill in the basic test here
				),
			},
		},
	})
}

func testAccCheckMarathonAppDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*marathon.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "marathon_app" {
			continue
		}

		client.AppRead("/test")

		// make sure that it's properly destroyed
	}

	return nil
}

const testAccCheckMarathonAppConfig_basic = `
resource "marathon_app" "app-create-example" {
	name = "app-create-example"

	cmd = "env && python3 -m http.server $PORT0"
	
	container {
		type = "DOCKER"
		docker {
			image = "python:3"
                }
	}

	cpus = "0.01"
	instances = 1
	mem = 100

}

`

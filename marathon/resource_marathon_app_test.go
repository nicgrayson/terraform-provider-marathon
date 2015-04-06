package marathon

import (
	"fmt"
	"time"

	"github.com/Banno/go-marathon"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"testing"
)

const testAccCheckMarathonAppConfig_basic = `
resource "marathon_app" "app-create-example" {
	 app_id= "/app-create-example"

	cmd = "env && python3 -m http.server $PORT0"

	container {
		docker {
			image = "python:3"
                }
	}

	cpus = "0.01"
	instances = 1
	mem = 100

}
`

const testAccCheckMarathonAppConfig_update = `
resource "marathon_app" "app-create-example" {
	app_id = "/app-create-example"

	cmd = "env && python3 -m http.server $PORT0"

	container {
		docker {
			image = "python:3"
                }
	}

	cpus = "0.01"
	instances = 2
	mem = 100

}
`

func TestAccMarathonApp_basic(t *testing.T) {

	var a marathon.App

	testCheckCreate := func(app *marathon.App) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			time.Sleep(1 * time.Second)
			if a.Version == "" {
				return fmt.Errorf("Didn't return a version so something is broken: %#v", app)
			}
			if a.Instances != 1 {
				return fmt.Errorf("AppCreate: Wrong number of instances %#v", app)
			}
			return nil
		}
	}

	testCheckUpdate := func(app *marathon.App) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			if a.Instances != 2 {
				return fmt.Errorf("AppUpdate: Wrong number of instances %#v", app)

			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMarathonAppDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckMarathonAppConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccReadApp("marathon_app.app-create-example", &a),
					testCheckCreate(&a),
				),
			},
			resource.TestStep{
				Config: testAccCheckMarathonAppConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccReadApp("marathon_app.app-create-example", &a),
					testCheckUpdate(&a),
				),
			},
		},
	})
}

func testAccReadApp(name string, app *marathon.App) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("marathon_app resource not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("marathon_app resource id not set correctly: %s", name)
		}

		//log.Printf("=== testAccContainerExists: rs ===\n%#v\n", rs)

		client := testAccProvider.Meta().(*marathon.Client)

		appRead, _ := client.AppRead(rs.Primary.Attributes["app_id"])

		//		log.Printf("=== testAccContainerExists: appRead ===\n%#v\n", appRead)

		time.Sleep(5000 * time.Millisecond)

		*app = *appRead

		return nil
	}
}

func testAccCheckMarathonAppDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*marathon.Client)

	_, err := client.AppRead("/app-create-example")
	if err == nil {
		return fmt.Errorf("App not deleted! %#v", err)
	}

	return nil
}

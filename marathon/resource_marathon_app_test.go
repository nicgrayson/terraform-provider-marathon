package marathon

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/ContainerLabs/go-marathon"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"marathon": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MARATHON_URL"); v == "" {
		t.Fatal("MARATHON_URL must be set for the acceptance tests to work.")
	}
}

func readExampleAppConfiguration(config string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("../test/%s.tf", config))
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func readExampleAppConfigurationAndUpdateInstanceCount(count int) string {
	config := readExampleAppConfiguration("example")
	re := regexp.MustCompile("instances = \\d+")
	updated := re.ReplaceAllString(config, fmt.Sprintf("instances = %d", count))
	return updated
}

func TestAccMarathonApp_basic(t *testing.T) {
	var a marathon.Application

	testCheckCreate := func(app *marathon.Application) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			time.Sleep(1 * time.Second)
			if a.Version == "" {
				return fmt.Errorf("Didn't return a version so something is broken: %#v", app)
			}
			if *a.Instances != 1 {
				return fmt.Errorf("AppCreate: Wrong number of instances %#v", app)
			}
			return nil
		}
	}

	testCheckUpdate := func(app *marathon.Application) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			if *a.Instances != 2 {
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
				Config: readExampleAppConfiguration("example"),
				Check: resource.ComposeTestCheckFunc(
					testAccReadApp("marathon_app.app-create-example", &a),
					testCheckCreate(&a),
				),
			},
			resource.TestStep{
				Config: readExampleAppConfigurationAndUpdateInstanceCount(2),
				Check: resource.ComposeTestCheckFunc(
					testAccReadApp("marathon_app.app-create-example", &a),
					testCheckUpdate(&a),
				),
			},
		},
	})
}

func TestAccMarathonApp_ipAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: readExampleAppConfiguration("ip-address"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("marathon_app.ip-address-create-example", "ipaddress.0.network_name", "default"),
				),
			},
		},
	})
}

func testAccReadApp(name string, app *marathon.Application) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("marathon_app resource not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("marathon_app resource id not set correctly: %s", name)
		}

		//log.Printf("=== testAccContainerExists: rs ===\n%#v\n", rs)

		config := testAccProvider.Meta().(config)
		client := config.Client

		appRead, _ := client.Application(rs.Primary.Attributes["app_id"])

		//		log.Printf("=== testAccContainerExists: appRead ===\n%#v\n", appRead)

		time.Sleep(5000 * time.Millisecond)

		*app = *appRead

		return nil
	}
}

func testAccCheckMarathonAppDestroy(s *terraform.State) error {
	time.Sleep(5000 * time.Millisecond)

	config := testAccProvider.Meta().(config)
	client := config.Client

	_, err := client.Application("/app-create-example")
	if err == nil {
		return fmt.Errorf("App not deleted! %#v", err)
	}

	return nil
}

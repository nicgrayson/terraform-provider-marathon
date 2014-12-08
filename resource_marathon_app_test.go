package marathon

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccMarathonApp_basic(t *testing.T) {

}

const testAccCheckMarathonAppConfig_basic = `
resource "marathon_app" "foo" {
        name = "test"
}
`

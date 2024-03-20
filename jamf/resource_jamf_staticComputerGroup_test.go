package jamf

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJamfStaticComputerGroup_basic(t *testing.T) {
	staticGroupName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceStaticGroupWithName(staticGroupName),
				ExpectError: regexp.MustCompile("Computed attributes cannot be set, but a value was set for \"computer.0.name\""),
			},
			{
				Config:      testAccResourceStaticGroupMissingSerial(staticGroupName),
				ExpectError: regexp.MustCompile("must provide exactly one of \"serial_number\" or \"id\""),
			},
			{
				Config:      testAccResourceStaticGroupWithSerialAndId(staticGroupName),
				ExpectError: regexp.MustCompile("must provide exactly one of \"serial_number\" or \"id\""),
			},
		},
	})
}

func testAccResourceStaticGroupWithName(staticGroupName string) string {
	return fmt.Sprintf(`
resource "jamf_staticComputerGroup" "test" {
	name = "%s"
	computer {
		name = "test-hostname"
	}
}`, staticGroupName)
}

func testAccResourceStaticGroupMissingSerial(staticGroupName string) string {
	return fmt.Sprintf(`
resource "jamf_staticComputerGroup" "test" {
	name = "%s"
	computer {
	}
}`, staticGroupName)
}

func testAccResourceStaticGroupWithSerialAndId(staticGroupName string) string {
	return fmt.Sprintf(`
resource "jamf_staticComputerGroup" "test" {
	name = "%s"
	computer {
		id = 1
		serial_number = "test-serial"
	}
}`, staticGroupName)
}

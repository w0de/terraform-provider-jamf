package jamf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJamfStaticComputerGroup_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckJamfStaticComputerGroupConfigWithName,
				ExpectError: regexp.MustCompile("The argument \"serial_number\" is required, but no definition was found."),
			},
			{
				Config:      testAccCheckJamfStaticComputerGroupConfigMissingSerial,
				ExpectError: regexp.MustCompile("The argument \"serial_number\" is required, but no definition was found."),
			},
			{
				Config:      testAccCheckJamfStaticComputerGroupConfigWithSerialAndId,
				ExpectError: regexp.MustCompile("Computed attributes cannot be set, but a value was set for \"computer.0.id\""),
			},
		},
	})
}

const (
	testAccCheckJamfStaticComputerGroupConfigWithName = `
resource "jamf_staticComputerGroup" "test" {
	name = "test"
	computer {
		name = "test-hostname"
	}
}`

	testAccCheckJamfStaticComputerGroupConfigMissingSerial = `
resource "jamf_staticComputerGroup" "test" {
	name = "test"
	computer {
	}
}`

	testAccCheckJamfStaticComputerGroupConfigWithSerialAndId = `
resource "jamf_staticComputerGroup" "test" {
name = "test"
computer {
	id = 1
	serial_number = "test-serial"
}
}`
)

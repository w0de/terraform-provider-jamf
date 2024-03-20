package jamf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// {
// 	Config: testAccCheckJamfComputerExtensionAttributeScript(extensionAttributeName),
// },
// {
// 	ResourceName:      "jamf_computer_extension_attribute.extensionattribute-script-4",
// 	ImportState:       true,
// 	ImportStateVerify: true,
// },

func TestAccJamfComputerExtensionAttribute_basic(t *testing.T) {
	extensionAttributeName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckJamfComputerExtensionAttributeScript(extensionAttributeName),
			},
			{
				ResourceName:      "jamf_computer_extension_attribute.extensionattribute_script",
				ImportState:       false,
				ImportStateVerify: false,
			},
			{
				Config: testAccCheckJamfComputerExtensionAttributeTextField(extensionAttributeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJamfComputerExtensionAttributeExists("jamf_computer_extension_attribute.extensionattribute-textfield"),
				),
			},
			{
				Config: testAccCheckJamfComputerExtensionAttributePopup(extensionAttributeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJamfComputerExtensionAttributeExists("jamf_computer_extension_attribute.extensionattribute-popup"),
				),
			},
		},
	})
}

func testAccCheckJamfComputerExtensionAttributeScript(extensionAttributeName string) string {
	return fmt.Sprintf(`
resource "jamf_computer_extension_attribute" "extensionattribute_script" {
	name = "%s"
	description = "testing jamf extension attribute resource"
	data_type = "String"
	inventory_display = "Extension Attributes" 

	script {
		enabled = false
		script_contents = "#!/bin/bash\nprint(\"hello world\")"
	}
}`, extensionAttributeName)
}

func testAccCheckJamfComputerExtensionAttributeTextField(extensionAttributeName string) string {
	return fmt.Sprintf(`
resource "jamf_computer_extension_attribute" "extensionattribute-textfield" {
	name = "%s"
	text_field {}
}`, extensionAttributeName)
}

func testAccCheckJamfComputerExtensionAttributePopup(extensionAttributeName string) string {
	return fmt.Sprintf(`
resource "jamf_computer_extension_attribute" "extensionattribute-popup" {
	name = "%s"
	popup_menu {
		choices = ["choice1", "choice2"]
	}
}`, extensionAttributeName)
}

func testAccCheckJamfComputerExtensionAttributeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		extensionattribute, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if extensionattribute.Primary.ID == "" {
			return fmt.Errorf("No resource id set")
		}

		return nil
	}
}

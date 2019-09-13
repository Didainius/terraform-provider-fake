package fake

import (
	"testing"

	"github.com/hashicorp/terraform/terraform"

	"github.com/hashicorp/terraform/helper/resource"
)

var testAccProvider = Provider()
var testAccProviders = map[string]terraform.ResourceProvider{
	"fake": testAccProvider,
}

func TestAccNestedSet(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("fake_nested_set.first", "name", "set1"),

					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "org", "vcd_org_vdc."+vdcName, "org"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "allocation_model", "vcd_org_vdc."+vdcName, "allocation_model"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "network_pool_name", "vcd_org_vdc."+vdcName, "network_pool_name"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "provider_vdc_name", "vcd_org_vdc."+vdcName, "provider_vdc_name"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "enabled", "vcd_org_vdc."+vdcName, "enabled"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "enable_thin_provisioning", "vcd_org_vdc."+vdcName, "enable_thin_provisioning"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "storage_profile.0.enabled", "vcd_org_vdc."+vdcName, "storage_profile.0.enabled"),
					// resource.TestCheckResourceAttrPair("data."+datasourceVdc, "storage_profile.0.default", "vcd_org_vdc."+vdcName, "storage_profile.0.default"),

					// resource.ComposeTestCheckFunc(testAccDataSourceVcdOrgVdc("data."+datasourceVdc, vdcName)),
				),
			},
		},
	})
}

const configText = `
resource "fake_nested_set" "first" {
	name = "set1"
	nested_set {
		subset1 {
			field1 = "11"
			field2 = "12"
		}

		subset2 {
			field1 = "21"
			field2 = "22"
		}
	}
}

data "fake_nested_set" "first" {
	name = "set1"

	depends_on = ["fake_nested_set.first"]
}
`

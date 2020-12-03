package fake

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

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
					resource.TestCheckResourceAttrPair("data.fake_nested_set.first", "name", "fake_nested_set.first", "name"),
					resource.TestCheckResourceAttrPair("data.fake_nested_set.first", "nested_set.#", "fake_nested_set.first", "nested_set.#"),
					resource.ComposeTestCheckFunc(testAccSpewState()),
				),
			},
		},
	})
}

func testAccSpewState() resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rootModule := s.RootModule()
		fmt.Println("====rootmodule====")
		spew.Dump(rootModule)

		return nil
	}
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
	name = "${fake_nested_set.first.name}"

}
`

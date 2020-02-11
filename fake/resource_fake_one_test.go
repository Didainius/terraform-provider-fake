package fake

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFakeOne(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("fake_one.first-example", "field1", "field1"),
					resource.TestCheckResourceAttr("fake_one.first-example", "field2", "field2"),
				),
			},
		},
	})
}

func TestAccFakeOne2(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("fake_one.first-example", "field1", "field1"),
					resource.TestCheckResourceAttr("fake_one.first-example", "field2", "field2"),
				),
			},
		},
	})
}

func TestAccFakeOne3(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("fake_one.first-example", "field1", "field1"),
					resource.TestCheckResourceAttr("fake_one.first-example", "field2", "field2"),
				),
			},
		},
	})
}

func TestAccFakeOne4(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("fake_one.first-example", "field1", "field1"),
					resource.TestCheckResourceAttr("fake_one.first-example", "field2", "field2"),
				),
			},
		},
	})
}

func TestAccFakeOne5(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configText,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("fake_one.first-example", "field1", "field1"),
					resource.TestCheckResourceAttr("fake_one.first-example", "field2", "field2"),
				),
			},
		},
	})
}

const configText = `
resource "fake_one" "first-example" {
	field1 = "field1"
	field2 = "field2"
}
`

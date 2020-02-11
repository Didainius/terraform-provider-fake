// Package fake is an implementation for fake terraform provider which is mainly used for testing
package fake

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"fake_one": resourceFakeOne(),
		},
	}
}

package main

import (
	"github.com/Didainius/terraform-provider-fake/fake"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return fake.Provider()
		},
	})
}

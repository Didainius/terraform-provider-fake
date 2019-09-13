package fake

import (
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
)

// globalLock is used to prevent multiple resources writing to the same file
var globalLock struct {
	sync.Mutex
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"fake_nested_set": resourceNestedSet(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fake_nested_set": dsNestedSet(),
		},
	}
}

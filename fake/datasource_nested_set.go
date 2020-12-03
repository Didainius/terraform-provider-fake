package fake

import "github.com/hashicorp/terraform/helper/schema"

func dsSubsetResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field1": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"field2": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dsSubsetSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem:     dsSubsetResource(),
	}
}

func dsSubset1Subset2Resource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subset1": dsSubsetSchema(),
			"subset2": dsSubsetSchema(),
		},
	}
}

func dsNestedSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceNestedSetRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name defines filename for writing",
			},
			"nested_set": &schema.Schema{
				Computed:    true,
				MinItems:    1,
				MaxItems:    1,
				Type:        schema.TypeSet,
				Elem:        dsSubset1Subset2Resource(),
				Description: "The compute capacity allocated to this VDC.",
			},
		},
	}
}

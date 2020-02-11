package fake

import (
	"math/rand"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFakeOne() *schema.Resource {
	return &schema.Resource{
		Create: resourceFakeOneCreate,
		Read:   resourceFakeOneRead,
		Update: resourceFakeOneUpdate,
		Delete: resourceFakeOneDelete,
		Importer: &schema.ResourceImporter{
			State: resourceFakeOneImport,
		},

		Schema: map[string]*schema.Schema{
			"field1": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"field2": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceFakeOneCreate
func resourceFakeOneCreate(d *schema.ResourceData, meta interface{}) error {
	value := rand.Int()
	string := strconv.Itoa(value)
	d.SetId(string)
	return nil
}

// resourceFakeOneUpdate
func resourceFakeOneUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceFakeOneRead
func resourceFakeOneRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceFakeOneDelete
func resourceFakeOneDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceFakeOneImport
func resourceFakeOneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}

package fake

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSimpleSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSimpleSetCreate,
		Read:   resourceSimpleSetRead,
		// Update: resourceSimpleSetUpdate,
		Delete: resourceSimpleSetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name defines filename for writing",
			},
			"nested_set": &schema.Schema{
				ForceNew: true,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					// Required: true,
					Type: schema.TypeSet,
					// MinItems: 2,
					// MaxItems: 2,
				},
			},
		},
	}
}

func resourceSimpleSetCreate(d *schema.ResourceData, m interface{}) error {
	globalLock.Lock()
	defer globalLock.Unlock()

	nestedSet := d.Get("nested_set")

	spew.Fdump(GetTerraformStdout(), nestedSet)

	err := setSimpleToFile(d)
	if err != nil {
		return fmt.Errorf("could not set data to file: %w", err)
	}

	d.SetId(d.Get("name").(string))

	return resourceNestedSetRead(d, m)
}

func resourceSimpleSetDelete(d *schema.ResourceData, m interface{}) error {
	globalLock.Lock()
	defer globalLock.Unlock()

	err := os.Remove(d.Get("name").(string))

	if err != nil {
		d.SetId("")
		return fmt.Errorf("could not remove file: %w", err)
	}

	return nil
}

func resourceSimpleSetRead(d *schema.ResourceData, m interface{}) error {

	// globalLock.Lock()
	// defer globalLock.Unlock()

	// err := getFromFile(d.Get("name").(string), d)
	// if err != nil {
	// 	d.SetId("")
	// 	return fmt.Errorf("could nor read data from file: %w", err)
	// }

	d.SetId(d.Get("name").(string))

	return nil
}

func setSimpleToFile(d *schema.ResourceData) error {
	name := d.Get("name").(string)
	nestedSetList := d.Get("nested_set").(*schema.Set).List()

	if len(nestedSetList) != 1 {
		return fmt.Errorf("not one 'nested_set' specified")
	}

	nestedSetMap := nestedSetList[0].(map[string]interface{})

	subset1List := nestedSetMap["subset1"].(*schema.Set).List()
	if len(subset1List) != 1 {
		return fmt.Errorf("not one 'subset1' specified")
	}

	subset2List := nestedSetMap["subset2"].(*schema.Set).List()
	if len(subset2List) != 1 {
		return fmt.Errorf("not one 'subset2' specified")
	}

	subset1Map := subset1List[0].(map[string]interface{})
	subset2Map := subset2List[0].(map[string]interface{})

	subset1Field1 := subset1Map["field1"].(int)
	subset1Field2 := subset1Map["field2"].(int)
	subset2Field1 := subset2Map["field1"].(int)
	subset2Field2 := subset2Map["field2"].(int)

	nestedSetStruct := NestedSet{
		Subset1: Subset{
			Field1: subset1Field1,
			Field2: subset1Field2,
		},
		Subset2: Subset{
			Field1: subset2Field1,
			Field2: subset2Field2,
		},
	}

	byteSlice, err := json.Marshal(&nestedSetStruct)
	if err != nil {
		return fmt.Errorf("unable to marshal nested_set: %w", err)
	}

	err = ioutil.WriteFile(name, byteSlice, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to file %s : %w", name, err)
	}

	return nil
}

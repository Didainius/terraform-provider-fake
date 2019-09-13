package fake

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/hashicorp/terraform/helper/schema"
)

// type JsonStruct struct {
type NestedSet struct {
	Subset1 Subset `json:"subset1"`
	Subset2 Subset `json:"subset2"`
}

type Subset struct {
	Field1 int `json:"field1"`
	Field2 int `json:"field2"`
}

func subsetResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field1": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"field2": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func subsetSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		//Set:      hashMapStringForCapacityElements,
		Elem: subsetResource(),
	}
}

func subset1Subset2Resource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subset1": subsetSchema(),
			"subset2": subsetSchema(),
		},
	}
}

func resourceNestedSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceNestedSetCreate,
		Read:   resourceNestedSetRead,
		// Update: resourceNestedSetUpdate,
		Delete: resourceNestedSetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name defines filename for writing",
			},
			"nested_set": &schema.Schema{
				ForceNew:    true,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Type:        schema.TypeSet,
				Elem:        subset1Subset2Resource(),
				Description: "The compute capacity allocated to this VDC.",
			},
		},
	}
}

func resourceNestedSetCreate(d *schema.ResourceData, m interface{}) error {
	globalLock.Lock()
	defer globalLock.Unlock()

	nestedSet := d.Get("nested_set")

	spew.Fdump(GetTerraformStdout(), nestedSet)

	err := setToFile(d)
	if err != nil {
		return fmt.Errorf("could not set data to file: %w", err)
	}

	d.SetId(d.Get("name").(string))

	return resourceNestedSetRead(d, m)
}

func resourceNestedSetRead(d *schema.ResourceData, m interface{}) error {

	// globalLock.Lock()
	// defer globalLock.Unlock()

	err := getFromFile(d.Get("name").(string), d)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("could nor read data from file: %w", err)
	}

	return nil
}

func resourceNestedSetDelete(d *schema.ResourceData, m interface{}) error {
	globalLock.Lock()
	defer globalLock.Unlock()

	return nil
}

func setToFile(d *schema.ResourceData) error {
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

func getFromFile(fileName string, d *schema.ResourceData) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", fileName, err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", fileName, err)
	}
	nestedSetStruct := NestedSet{}

	err = json.Unmarshal(b, &nestedSetStruct)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	// Subset 1 fields
	subset1Fields := make(map[string]interface{})
	subset1Fields["field1"] = nestedSetStruct.Subset1.Field1
	subset1Fields["field2"] = nestedSetStruct.Subset1.Field2

	subset1FieldsInterface := make([]interface{}, 0)
	subset1FieldsInterface = append(subset1FieldsInterface, subset1Fields)

	// subset1FieldsSet := *schema.NewSet(schema.HashResource(subsetResource()), subset1FieldsInterface)

	// EOF layer 1 hashing

	// Create a parent "subset1" for fields 1 and 2
	subset1Map := make(map[string]interface{})
	subset1Map["subset1"] = subset1FieldsInterface

	subset1 := *schema.NewSet(schema.HashResource(subsetResource()), subset1FieldsInterface)

	subset1Interface := make([]interface{}, 0)
	subset1Interface = append(subset1Interface, subset1)

	// Subset 2 fields
	subset2Fields := make(map[string]interface{})
	subset2Fields["field1"] = nestedSetStruct.Subset2.Field1
	subset2Fields["field2"] = nestedSetStruct.Subset2.Field2

	subset2FieldsInterface := make([]interface{}, 0)
	subset2FieldsInterface = append(subset2FieldsInterface, subset2Fields)

	// subset2FieldsSet := *schema.NewSet(schema.HashResource(subsetResource()), subset2FieldsInterface)

	// Create a parent "subset2" for fields 1 and 2
	subset2Map := make(map[string]interface{})
	subset2Map["subset2"] = subset2Fields

	// Insert child subset2Fields into parent subset2
	// subset2Interface = append(subset2Interface, subset2Fields)
	// Make TypeSet for subset2
	subset2 := *schema.NewSet(schema.HashResource(subsetResource()), subset2FieldsInterface)

	subset2Interface := make([]interface{}, 0)
	subset2Interface = append(subset2Interface, subset2)

	// Create a parent "nested_set" for "subset1" and "subset2"

	rootMap := make(map[string]interface{})
	rootMap["subset1"] = &subset1
	rootMap["subset2"] = &subset2

	nestedSchemaInterface := make([]interface{}, 0)
	nestedSchemaInterface = append(nestedSchemaInterface, rootMap)

	nestedSetSchema := *schema.NewSet(schema.HashResource(subset1Subset2Resource()), nestedSchemaInterface)

	spew.Fdump(GetTerraformStdout(), nestedSetSchema)

	return d.Set("nested_set", &nestedSetSchema)
}

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 * FirestoreIndex api apends __name__ as an item to the
 * end of the fields list if not present. We are suppressing
 * this server generated field.
 */
func FirestoreIFieldsDiffSuppressFunc(k, old, new string, d TerraformResourceDataChange) bool {
	kLength := "fields.#"
	oldLength, newLength := d.GetChange(kLength)
	oldInt, ok := oldLength.(int)
	if !ok {
		return false
	}
	newInt, ok := newLength.(int)
	if !ok {
		return false
	}

	if oldInt == newInt+1 {
		kold := fmt.Sprintf("fields.%v.field_path", oldInt-1)
		knew := fmt.Sprintf("fields.%v.field_path", newInt-1)

		oldLastIndexName, _ := d.GetChange(kold)
		_, newLastIndexName := d.GetChange(knew)
		if oldLastIndexName == "__name__" && newLastIndexName != "__name__" {
			oldBase := fmt.Sprintf("fields.%v", oldInt-1)
			if strings.HasPrefix(k, oldBase) || k == kLength {
				return true
			}
		}
	}
	return false
}

func firestoreIFieldsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return FirestoreIFieldsDiffSuppressFunc(k, old, new, d)
}

func ResourceFirestoreIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreIndexCreate,
		Read:   resourceFirestoreIndexRead,
		Delete: resourceFirestoreIndexDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreIndexImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"collection": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The collection being indexed.`,
			},
			"fields": {
				Type:             schema.TypeList,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: firestoreIFieldsDiffSuppress,
				Description: `The fields supported by this index. The last field entry is always for
the field path '__name__'. If, on creation, '__name__' was not
specified as the last field, it will be added automatically with the
same direction as that of the last field defined. If the final field
in a composite index is not directional, the '__name__' will be
ordered '"ASCENDING"' (unless explicitly specified otherwise).`,
				MinItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"array_config": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validateEnum([]string{"CONTAINS", ""}),
							Description: `Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can
be specified. Possible values: ["CONTAINS"]`,
						},
						"field_path": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Name of the field.`,
						},
						"order": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validateEnum([]string{"ASCENDING", "DESCENDING", ""}),
							Description: `Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"]`,
						},
					},
				},
			},
			"database": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The Firestore database id. Defaults to '"(default)"'.`,
				Default:     "(default)",
			},
			"query_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"COLLECTION", "COLLECTION_GROUP", ""}),
				Description:  `The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"]`,
				Default:      "COLLECTION",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A server defined name for this index. Format:
'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirestoreIndexCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	databaseProp, err := expandFirestoreIndexDatabase(d.Get("database"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database"); !isEmptyValue(reflect.ValueOf(databaseProp)) && (ok || !reflect.DeepEqual(v, databaseProp)) {
		obj["database"] = databaseProp
	}
	collectionProp, err := expandFirestoreIndexCollection(d.Get("collection"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collection"); !isEmptyValue(reflect.ValueOf(collectionProp)) && (ok || !reflect.DeepEqual(v, collectionProp)) {
		obj["collection"] = collectionProp
	}
	queryScopeProp, err := expandFirestoreIndexQueryScope(d.Get("query_scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("query_scope"); !isEmptyValue(reflect.ValueOf(queryScopeProp)) && (ok || !reflect.DeepEqual(v, queryScopeProp)) {
		obj["queryScope"] = queryScopeProp
	}
	fieldsProp, err := expandFirestoreIndexFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	obj, err = resourceFirestoreIndexEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Index: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Index: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Index", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Index: %s", err)
	}

	if err := d.Set("name", flattenFirestoreIndexName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// The operation for this resource contains the generated name that we need
	// in order to perform a READ.
	metadata := res["metadata"].(map[string]interface{})
	name := metadata["index"].(string)
	log.Printf("[DEBUG] Setting Index name, id to %s", name)
	if err := d.Set("name", name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)

	log.Printf("[DEBUG] Finished creating Index %q: %#v", d.Id(), res)

	return resourceFirestoreIndexRead(d, meta)
}

func resourceFirestoreIndexRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirestoreIndex %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	if err := d.Set("name", flattenFirestoreIndexName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("query_scope", flattenFirestoreIndexQueryScope(res["queryScope"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("fields", flattenFirestoreIndexFields(res["fields"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	return nil
}

func resourceFirestoreIndexDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Index %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Index")
	}

	err = FirestoreOperationWaitTime(
		config, res, project, "Deleting Index", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Index %q: %#v", d.Id(), res)
	return nil
}

func resourceFirestoreIndexImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) != 8 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("name"),
			"projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}",
		)
	}

	if err := d.Set("project", stringParts[1]); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("database", stringParts[3]); err != nil {
		return nil, fmt.Errorf("Error setting database: %s", err)
	}
	if err := d.Set("collection", stringParts[5]); err != nil {
		return nil, fmt.Errorf("Error setting collection: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreIndexName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreIndexQueryScope(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreIndexFields(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"field_path":   flattenFirestoreIndexFieldsFieldPath(original["fieldPath"], d, config),
			"order":        flattenFirestoreIndexFieldsOrder(original["order"], d, config),
			"array_config": flattenFirestoreIndexFieldsArrayConfig(original["arrayConfig"], d, config),
		})
	}
	return transformed
}
func flattenFirestoreIndexFieldsFieldPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreIndexFieldsOrder(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirestoreIndexFieldsArrayConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirestoreIndexDatabase(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexCollection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexQueryScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFieldPath, err := expandFirestoreIndexFieldsFieldPath(original["field_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFieldPath); val.IsValid() && !isEmptyValue(val) {
			transformed["fieldPath"] = transformedFieldPath
		}

		transformedOrder, err := expandFirestoreIndexFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedArrayConfig, err := expandFirestoreIndexFieldsArrayConfig(original["array_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArrayConfig); val.IsValid() && !isEmptyValue(val) {
			transformed["arrayConfig"] = transformedArrayConfig
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirestoreIndexFieldsFieldPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsArrayConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceFirestoreIndexEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// We've added project / database / collection as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	return obj, nil
}

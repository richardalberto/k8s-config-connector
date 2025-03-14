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

func ResourceApiGatewayApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayApiCreate,
		Read:   resourceApiGatewayApiRead,
		Update: resourceApiGatewayApiUpdate,
		Delete: resourceApiGatewayApiDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApiGatewayApiImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Identifier to assign to the API. Must be unique within scope of the parent resource(project)`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `A user-visible name for the API.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"managed_service": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
If not specified, a new Service will automatically be created in the same project as this API.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'`,
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

func resourceApiGatewayApiCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayApiDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	managedServiceProp, err := expandApiGatewayApiManagedService(d.Get("managed_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("managed_service"); !isEmptyValue(reflect.ValueOf(managedServiceProp)) && (ok || !reflect.DeepEqual(v, managedServiceProp)) {
		obj["managedService"] = managedServiceProp
	}
	labelsProp, err := expandApiGatewayApiLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis?apiId={{api_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Api: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Api: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Api: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApiGatewayOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Api", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Api: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Api %q: %#v", d.Id(), res)

	return resourceApiGatewayApiRead(d, meta)
}

func resourceApiGatewayApiRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Api: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApiGatewayApi %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}

	if err := d.Set("name", flattenApiGatewayApiName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}
	if err := d.Set("display_name", flattenApiGatewayApiDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}
	if err := d.Set("managed_service", flattenApiGatewayApiManagedService(res["managedService"], d, config)); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}
	if err := d.Set("create_time", flattenApiGatewayApiCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}
	if err := d.Set("labels", flattenApiGatewayApiLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Api: %s", err)
	}

	return nil
}

func resourceApiGatewayApiUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Api: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayApiDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandApiGatewayApiLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Api %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Api %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Api %q: %#v", d.Id(), res)
	}

	err = ApiGatewayOperationWaitTime(
		config, res, project, "Updating Api", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApiGatewayApiRead(d, meta)
}

func resourceApiGatewayApiDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Api: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Api %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Api")
	}

	err = ApiGatewayOperationWaitTime(
		config, res, project, "Deleting Api", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Api %q: %#v", d.Id(), res)
	return nil
}

func resourceApiGatewayApiImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/apis/(?P<api_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<api_id>[^/]+)",
		"(?P<api_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApiGatewayApiName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApiGatewayApiDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApiGatewayApiManagedService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApiGatewayApiCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApiGatewayApiLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApiGatewayApiDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiManagedService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

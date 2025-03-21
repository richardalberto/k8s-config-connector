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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var ApigeeEnvironmentIamSchema = map[string]*schema.Schema{
	"org_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"env_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ApigeeEnvironmentIamUpdater struct {
	orgId  string
	envId  string
	d      TerraformResourceData
	Config *Config
}

func ApigeeEnvironmentIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("org_id"); ok {
		values["org_id"] = v.(string)
	}

	if v, ok := d.GetOk("env_id"); ok {
		values["env_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"(?P<org_id>.+)/environments/(?P<env_id>[^/]+)", "(?P<env_id>[^/]+)"}, d, config, d.Get("env_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ApigeeEnvironmentIamUpdater{
		orgId:  values["org_id"],
		envId:  values["env_id"],
		d:      d,
		Config: config,
	}

	if err := d.Set("org_id", u.orgId); err != nil {
		return nil, fmt.Errorf("Error setting org_id: %s", err)
	}
	if err := d.Set("env_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting env_id: %s", err)
	}

	return u, nil
}

func ApigeeEnvironmentIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	m, err := getImportIdQualifiers([]string{"(?P<org_id>.+)/environments/(?P<env_id>[^/]+)", "(?P<env_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ApigeeEnvironmentIamUpdater{
		orgId:  values["org_id"],
		envId:  values["env_id"],
		d:      d,
		Config: config,
	}
	if err := d.Set("env_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting env_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ApigeeEnvironmentIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyEnvironmentUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "GET", "", url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *ApigeeEnvironmentIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyEnvironmentUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", "", url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ApigeeEnvironmentIamUpdater) qualifyEnvironmentUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ApigeeBasePath}}%s:%s", fmt.Sprintf("%s/environments/%s", u.orgId, u.envId), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ApigeeEnvironmentIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s/environments/%s", u.orgId, u.envId)
}

func (u *ApigeeEnvironmentIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-apigee-environment-%s", u.GetResourceId())
}

func (u *ApigeeEnvironmentIamUpdater) DescribeResource() string {
	return fmt.Sprintf("apigee environment %q", u.GetResourceId())
}

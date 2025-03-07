package google

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2/google"
)

func TestHandleSDKDefaults_ImpersonateServiceAccount(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"impersonate_service_account value set in the provider schema is not overridden by ENVs": {
			ConfigValue: "value-from-config@example.com",
			EnvVariables: map[string]string{
				"GOOGLE_IMPERSONATE_SERVICE_ACCOUNT": "value-from-env@example.com",
			},
			ExpectedValue: "value-from-config@example.com",
		},
		"impersonate_service_account value can be set by environment variable": {
			EnvVariables: map[string]string{
				"GOOGLE_IMPERSONATE_SERVICE_ACCOUNT": "value-from-env@example.com",
			},
			ExpectedValue: "value-from-env@example.com",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("impersonate_service_account", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("impersonate_service_account")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected impersonate_service_account to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected impersonate_service_account to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestHandleSDKDefaults_Project(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"project value set in the provider config is not overridden by ENVs": {
			ConfigValue: "my-project-from-config",
			EnvVariables: map[string]string{
				"GOOGLE_PROJECT": "my-project-from-env",
			},
			ExpectedValue: "my-project-from-config",
		},
		"project can be set by environment variable, when no value supplied via the config": {
			EnvVariables: map[string]string{
				"GOOGLE_PROJECT": "my-project-from-env",
			},
			ExpectedValue: "my-project-from-env",
		},
		"when multiple project environment variables are provided, `GOOGLE_PROJECT` is used first": {
			EnvVariables: map[string]string{
				"GOOGLE_PROJECT":        "project-from-GOOGLE_PROJECT",
				"GOOGLE_CLOUD_PROJECT":  "project-from-GOOGLE_CLOUD_PROJECT",
				"GCLOUD_PROJECT":        "project-from-GCLOUD_PROJECT",
				"CLOUDSDK_CORE_PROJECT": "project-from-CLOUDSDK_CORE_PROJECT",
			},
			ExpectedValue: "project-from-GOOGLE_PROJECT",
		},
		"when multiple project environment variables are provided, `GOOGLE_CLOUD_PROJECT` is used second": {
			EnvVariables: map[string]string{
				// GOOGLE_PROJECT unset
				"GOOGLE_CLOUD_PROJECT":  "project-from-GOOGLE_CLOUD_PROJECT",
				"GCLOUD_PROJECT":        "project-from-GCLOUD_PROJECT",
				"CLOUDSDK_CORE_PROJECT": "project-from-CLOUDSDK_CORE_PROJECT",
			},
			ExpectedValue: "project-from-GOOGLE_CLOUD_PROJECT",
		},
		"when multiple project environment variables are provided, `GCLOUD_PROJECT` is used third": {
			EnvVariables: map[string]string{
				// GOOGLE_PROJECT unset
				// GOOGLE_CLOUD_PROJECT unset
				"GCLOUD_PROJECT":        "project-from-GCLOUD_PROJECT",
				"CLOUDSDK_CORE_PROJECT": "project-from-CLOUDSDK_CORE_PROJECT",
			},
			ExpectedValue: "project-from-GCLOUD_PROJECT",
		},
		"when multiple project environment variables are provided, `CLOUDSDK_CORE_PROJECT` is the last-used ENV": {
			EnvVariables: map[string]string{
				// GOOGLE_PROJECT unset
				// GOOGLE_CLOUD_PROJECT unset
				// GCLOUD_PROJECT unset
				"CLOUDSDK_CORE_PROJECT": "project-from-CLOUDSDK_CORE_PROJECT",
			},
			ExpectedValue: "project-from-CLOUDSDK_CORE_PROJECT",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("project", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("project")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected project to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected project to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestHandleSDKDefaults_BillingProject(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"billing project value set in the provider config is not overridden by ENVs": {
			ConfigValue: "my-billing-project-from-config",
			EnvVariables: map[string]string{
				"GOOGLE_BILLING_PROJECT": "my-billing-project-from-env",
			},
			ExpectedValue: "my-billing-project-from-config",
		},
		"billing project can be set by environment variable, when no value supplied via the config": {
			EnvVariables: map[string]string{
				"GOOGLE_BILLING_PROJECT": "my-billing-project-from-env",
			},
			ExpectedValue: "my-billing-project-from-env",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("billing_project", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("billing_project")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected billing_project to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected billing_project to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestHandleSDKDefaults_Region(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"region value set in the provider config is not overridden by ENVs": {
			ConfigValue: "region-from-config",
			EnvVariables: map[string]string{
				"GOOGLE_REGION": "region-from-env",
			},
			ExpectedValue: "region-from-config",
		},
		"region can be set by environment variable, when no value supplied via the config": {
			EnvVariables: map[string]string{
				"GOOGLE_REGION": "region-from-env",
			},
			ExpectedValue: "region-from-env",
		},
		"when multiple region environment variables are provided, `GOOGLE_REGION` is used first": {
			EnvVariables: map[string]string{
				"GOOGLE_REGION":           "project-from-GOOGLE_REGION",
				"GCLOUD_REGION":           "project-from-GCLOUD_REGION",
				"CLOUDSDK_COMPUTE_REGION": "project-from-CLOUDSDK_COMPUTE_REGION",
			},
			ExpectedValue: "project-from-GOOGLE_REGION",
		},
		"when multiple region environment variables are provided, `GCLOUD_REGION` is used second": {
			EnvVariables: map[string]string{
				// GOOGLE_REGION unset
				"GCLOUD_REGION":           "project-from-GCLOUD_REGION",
				"CLOUDSDK_COMPUTE_REGION": "project-from-CLOUDSDK_COMPUTE_REGION",
			},
			ExpectedValue: "project-from-GCLOUD_REGION",
		},
		"when multiple region environment variables are provided, `CLOUDSDK_COMPUTE_REGION` is the last-used ENV": {
			EnvVariables: map[string]string{
				// GOOGLE_REGION unset
				// GCLOUD_REGION unset
				"CLOUDSDK_COMPUTE_REGION": "project-from-CLOUDSDK_COMPUTE_REGION",
			},
			ExpectedValue: "project-from-CLOUDSDK_COMPUTE_REGION",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("region", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("region")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected region to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected region to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestHandleSDKDefaults_Zone(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"region value set in the provider config is not overridden by ENVs": {
			ConfigValue: "zone-from-config",
			EnvVariables: map[string]string{
				"GOOGLE_ZONE": "zone-from-env",
			},
			ExpectedValue: "zone-from-config",
		},
		"zone can be set by environment variable, when no value supplied via the config": {
			EnvVariables: map[string]string{
				"GOOGLE_ZONE": "zone-from-env",
			},
			ExpectedValue: "zone-from-env",
		},
		"when multiple zone environment variables are provided, `GOOGLE_ZONE` is used first": {
			EnvVariables: map[string]string{
				"GOOGLE_ZONE":           "zone-from-GOOGLE_ZONE",
				"GCLOUD_ZONE":           "zone-from-GCLOUD_ZONE",
				"CLOUDSDK_COMPUTE_ZONE": "zone-from-CLOUDSDK_COMPUTE_ZONE",
			},
			ExpectedValue: "zone-from-GOOGLE_ZONE",
		},
		"when multiple zone environment variables are provided, `GCLOUD_ZONE` is used second": {
			EnvVariables: map[string]string{
				// GOOGLE_ZONE unset
				"GCLOUD_ZONE":           "zone-from-GCLOUD_ZONE",
				"CLOUDSDK_COMPUTE_ZONE": "zone-from-CLOUDSDK_COMPUTE_ZONE",
			},
			ExpectedValue: "zone-from-GCLOUD_ZONE",
		},
		"when multiple zone environment variables are provided, `CLOUDSDK_COMPUTE_ZONE` is the last-used ENV": {
			EnvVariables: map[string]string{
				// GOOGLE_ZONE unset
				// GCLOUD_ZONE unset
				"CLOUDSDK_COMPUTE_ZONE": "zone-from-CLOUDSDK_COMPUTE_ZONE",
			},
			ExpectedValue: "zone-from-CLOUDSDK_COMPUTE_ZONE",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("zone", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("zone")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected zone to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected zone to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

// The `user_project_override` field is an odd one out, as other provider schema fields tend to be strings
// and `user_project_override` is a boolean
func TestHandleSDKDefaults_UserProjectOverride(t *testing.T) {
	cases := map[string]struct {
		SetViaConfig     bool // Awkward, but necessary as zero value of ConfigValue could be intended
		ConfigValue      bool
		ValueNotProvided bool
		EnvVariables     map[string]string
		ExpectedValue    bool
		ExpectError      bool
	}{
		"user_project_override value set in the provider schema is not overridden by ENVs": {
			SetViaConfig: true,
			ConfigValue:  false,
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "true",
			},
			ExpectedValue: false,
		},
		"user_project_override can be set by environment variable: true": {
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "true",
			},
			ExpectedValue: true,
		},
		"user_project_override can be set by environment variable: false": {
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "false",
			},
			ExpectedValue: false,
		},
		"user_project_override can be set by environment variable: 1": {
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "1",
			},
			ExpectedValue: true,
		},
		"user_project_override can be set by environment variable: 0": {
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "0",
			},
			ExpectedValue: false,
		},
		"error returned due to non-boolean environment variables": {
			EnvVariables: map[string]string{
				"USER_PROJECT_OVERRIDE": "I'm not a boolean",
			},
			ExpectError: true,
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.SetViaConfig {
				d.Set("user_project_override", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			v, ok := d.GetOkExists("user_project_override")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected user_project_override to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected user_project_override to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestHandleSDKDefaults_RequestReason(t *testing.T) {
	cases := map[string]struct {
		ConfigValue      string
		EnvVariables     map[string]string
		ExpectedValue    string
		ValueNotProvided bool
		ExpectError      bool
	}{
		"request_reason value set in the provider config is not overridden by ENVs": {
			ConfigValue: "request-reason-from-config",
			EnvVariables: map[string]string{
				"CLOUDSDK_CORE_REQUEST_REASON": "request-reason-from-env",
			},
			ExpectedValue: "request-reason-from-config",
		},
		"request_reason can be set by environment variable, when no value supplied via the config": {
			EnvVariables: map[string]string{
				"CLOUDSDK_CORE_REQUEST_REASON": "request-reason-from-env",
			},
			ExpectedValue: "request-reason-from-env",
		},
		"when no values are provided via config or environment variables, the field remains unset without error": {
			ValueNotProvided: true,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {

			// Arrange
			// Create empty schema.ResourceData using the SDK Provider schema
			emptyConfigMap := map[string]interface{}{}
			d := schema.TestResourceDataRaw(t, Provider().Schema, emptyConfigMap)

			// Set config value(s)
			if tc.ConfigValue != "" {
				d.Set("request_reason", tc.ConfigValue)
			}

			// Set ENVs
			if len(tc.EnvVariables) > 0 {
				for k, v := range tc.EnvVariables {
					t.Setenv(k, v)
				}
			}

			// Act
			err := HandleSDKDefaults(d)

			// Assert
			if err != nil {
				if !tc.ExpectError {
					t.Fatalf("error: %v", err)
				}
				return
			}

			// Assert
			v, ok := d.GetOk("request_reason")
			if !ok && !tc.ValueNotProvided {
				t.Fatal("expected request_reason to be set in the provider data")
			}
			if ok && tc.ValueNotProvided {
				t.Fatal("expected request_reason to not be set in the provider data")
			}

			if v != tc.ExpectedValue {
				t.Fatalf("unexpected value: wanted %v, got, %v", tc.ExpectedValue, v)
			}
		})
	}
}

func TestConfigLoadAndValidate_accountFilePath(t *testing.T) {
	config := &Config{
		Credentials: testFakeCredentialsPath,
		Project:     "my-gce-project",
		Region:      "us-central1",
	}

	ConfigureBasePaths(config)

	err := config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestConfigLoadAndValidate_accountFileJSON(t *testing.T) {
	contents, err := ioutil.ReadFile(testFakeCredentialsPath)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	config := &Config{
		Credentials: string(contents),
		Project:     "my-gce-project",
		Region:      "us-central1",
	}

	ConfigureBasePaths(config)

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestConfigLoadAndValidate_accountFileJSONInvalid(t *testing.T) {
	config := &Config{
		Credentials: "{this is not json}",
		Project:     "my-gce-project",
		Region:      "us-central1",
	}

	ConfigureBasePaths(config)

	if config.LoadAndValidate(context.Background()) == nil {
		t.Fatalf("expected error, but got nil")
	}
}

func TestAccConfigLoadValidate_credentials(t *testing.T) {
	if os.Getenv(TestEnvVar) == "" {
		t.Skipf("Network access not allowed; use %s=1 to enable", TestEnvVar)
	}
	AccTestPreCheck(t)

	creds := GetTestCredsFromEnv()
	proj := GetTestProjectFromEnv()

	config := &Config{
		Credentials: creds,
		Project:     proj,
		Region:      "us-central1",
	}

	ConfigureBasePaths(config)

	err := config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	_, err = config.NewComputeClient(config.UserAgent).Zones.Get(proj, "us-central1-a").Do()
	if err != nil {
		t.Fatalf("expected call with loaded config client to work, got error: %s", err)
	}
}

func TestAccConfigLoadValidate_impersonated(t *testing.T) {
	if os.Getenv(TestEnvVar) == "" {
		t.Skipf("Network access not allowed; use %s=1 to enable", TestEnvVar)
	}
	AccTestPreCheck(t)

	serviceaccount := MultiEnvSearch([]string{"IMPERSONATE_SERVICE_ACCOUNT_ACCTEST"})
	creds := GetTestCredsFromEnv()
	proj := GetTestProjectFromEnv()

	config := &Config{
		Credentials:               creds,
		ImpersonateServiceAccount: serviceaccount,
		Project:                   proj,
		Region:                    "us-central1",
	}

	ConfigureBasePaths(config)

	err := config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	_, err = config.NewComputeClient(config.UserAgent).Zones.Get(proj, "us-central1-a").Do()
	if err != nil {
		t.Fatalf("expected API call with loaded config to work, got error: %s", err)
	}
}

func TestAccConfigLoadValidate_accessTokenImpersonated(t *testing.T) {
	if os.Getenv(TestEnvVar) == "" {
		t.Skipf("Network access not allowed; use %s=1 to enable", TestEnvVar)
	}
	AccTestPreCheck(t)

	creds := GetTestCredsFromEnv()
	proj := GetTestProjectFromEnv()
	serviceaccount := MultiEnvSearch([]string{"IMPERSONATE_SERVICE_ACCOUNT_ACCTEST"})

	c, err := google.CredentialsFromJSON(context.Background(), []byte(creds), DefaultClientScopes...)
	if err != nil {
		t.Fatalf("invalid test credentials: %s", err)
	}

	token, err := c.TokenSource.Token()
	if err != nil {
		t.Fatalf("Unable to generate test access token: %s", err)
	}

	config := &Config{
		AccessToken:               token.AccessToken,
		ImpersonateServiceAccount: serviceaccount,
		Project:                   proj,
		Region:                    "us-central1",
	}

	ConfigureBasePaths(config)

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	_, err = config.NewComputeClient(config.UserAgent).Zones.Get(proj, "us-central1-a").Do()
	if err != nil {
		t.Fatalf("expected API call with loaded config to work, got error: %s", err)
	}
}

func TestAccConfigLoadValidate_accessToken(t *testing.T) {
	if os.Getenv(TestEnvVar) == "" {
		t.Skipf("Network access not allowed; use %s=1 to enable", TestEnvVar)
	}
	AccTestPreCheck(t)

	creds := GetTestCredsFromEnv()
	proj := GetTestProjectFromEnv()

	c, err := google.CredentialsFromJSON(context.Background(), []byte(creds), testOauthScope)
	if err != nil {
		t.Fatalf("invalid test credentials: %s", err)
	}

	token, err := c.TokenSource.Token()
	if err != nil {
		t.Fatalf("Unable to generate test access token: %s", err)
	}

	config := &Config{
		AccessToken: token.AccessToken,
		Project:     proj,
		Region:      "us-central1",
	}

	ConfigureBasePaths(config)

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	_, err = config.NewComputeClient(config.UserAgent).Zones.Get(proj, "us-central1-a").Do()
	if err != nil {
		t.Fatalf("expected API call with loaded config to work, got error: %s", err)
	}
}

func TestConfigLoadAndValidate_customScopes(t *testing.T) {
	config := &Config{
		Credentials: testFakeCredentialsPath,
		Project:     "my-gce-project",
		Region:      "us-central1",
		Scopes:      []string{"https://www.googleapis.com/auth/compute"},
	}

	ConfigureBasePaths(config)

	err := config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(config.Scopes) != 1 {
		t.Fatalf("expected 1 scope, got %d scopes: %v", len(config.Scopes), config.Scopes)
	}
	if config.Scopes[0] != "https://www.googleapis.com/auth/compute" {
		t.Fatalf("expected scope to be %q, got %q", "https://www.googleapis.com/auth/compute", config.Scopes[0])
	}
}

func TestConfigLoadAndValidate_defaultBatchingConfig(t *testing.T) {
	// Use default batching config
	batchCfg, err := ExpandProviderBatchingConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	config := &Config{
		Credentials:    testFakeCredentialsPath,
		Project:        "my-gce-project",
		Region:         "us-central1",
		BatchingConfig: batchCfg,
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedDur := time.Second * DefaultBatchSendIntervalSec
	if config.RequestBatcherServiceUsage.SendAfter != expectedDur {
		t.Fatalf("expected SendAfter to be %d seconds, got %v",
			DefaultBatchSendIntervalSec,
			config.RequestBatcherServiceUsage.SendAfter)
	}
}

func TestConfigLoadAndValidate_customBatchingConfig(t *testing.T) {
	batchCfg, err := ExpandProviderBatchingConfig([]interface{}{
		map[string]interface{}{
			"send_after":      "1s",
			"enable_batching": false,
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if batchCfg.SendAfter != time.Second {
		t.Fatalf("expected batchCfg SendAfter to be 1 second, got %v", batchCfg.SendAfter)
	}
	if batchCfg.EnableBatching {
		t.Fatalf("expected EnableBatching to be false")
	}

	config := &Config{
		Credentials:    testFakeCredentialsPath,
		Project:        "my-gce-project",
		Region:         "us-central1",
		BatchingConfig: batchCfg,
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedDur := time.Second * 1
	if config.RequestBatcherServiceUsage.SendAfter != expectedDur {
		t.Fatalf("expected SendAfter to be %d seconds, got %v",
			1,
			config.RequestBatcherServiceUsage.SendAfter)
	}

	if config.RequestBatcherServiceUsage.EnableBatching {
		t.Fatalf("expected EnableBatching to be false")
	}
}

func TestRemoveBasePathVersion(t *testing.T) {
	cases := []struct {
		BaseURL  string
		Expected string
	}{
		{"https://www.googleapis.com/compute/version_v1/", "https://www.googleapis.com/compute/"},
		{"https://runtimeconfig.googleapis.com/v1beta1/", "https://runtimeconfig.googleapis.com/"},
		{"https://www.googleapis.com/compute/v1/", "https://www.googleapis.com/compute/"},
		{"https://staging-version.googleapis.com/", "https://staging-version.googleapis.com/"},
		// For URLs with any parts, the last part is always removed- it's assumed to be the version.
		{"https://runtimeconfig.googleapis.com/runtimeconfig/", "https://runtimeconfig.googleapis.com/"},
	}

	for _, c := range cases {
		if c.Expected != RemoveBasePathVersion(c.BaseURL) {
			t.Errorf("replace url failed: got %s wanted %s", RemoveBasePathVersion(c.BaseURL), c.Expected)
		}
	}
}

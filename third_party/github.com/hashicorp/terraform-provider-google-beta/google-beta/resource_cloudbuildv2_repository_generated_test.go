// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudbuildv2 "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccCloudbuildv2Repository_GheRepository(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  GetTestProjectFromEnv(),
		"region":        GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck: func() { AccTestPreCheck(t) },

		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2RepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Repository_GheRepository(context),
			},
			{
				ResourceName:      "google_cloudbuildv2_repository.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccCloudbuildv2Repository_GithubRepository(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  GetTestProjectFromEnv(),
		"region":        GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck: func() { AccTestPreCheck(t) },

		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2RepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Repository_GithubRepository(context),
			},
			{
				ResourceName:      "google_cloudbuildv2_repository.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudbuildv2Repository_GheRepository(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuildv2_repository" "primary" {
  name              = "tf-test-repository%{random_suffix}"
  parent_connection = google_cloudbuildv2_connection.ghe_complete.name
  remote_uri        = "https://ghe.proctor-staging-test.com/proctorteam/regional_test.git"

  annotations = {
    some-key = "some-value"
  }

  location = "%{region}"
  project  = "%{project_name}"
  provider          = google-beta
}
resource "google_cloudbuildv2_connection" "ghe_complete" {
  location    = "%{region}"
  name        = "tf-test-connection%{random_suffix}"
  annotations = {}

  github_enterprise_config {
    host_uri                      = "https://ghe.proctor-staging-test.com"
    app_id                        = 516
    app_installation_id           = 243
    app_slug                      = "myapp"
    private_key_secret_version    = "projects/gcb-terraform-creds/secrets/ghe-private-key/versions/latest"
    webhook_secret_secret_version = "projects/gcb-terraform-creds/secrets/ghe-webhook-secret/versions/latest"
  }

  project = "%{project_name}"
  provider    = google-beta
}

`, context)
}

func testAccCloudbuildv2Repository_GithubRepository(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuildv2_repository" "primary" {
  name              = "tf-test-repository%{random_suffix}"
  parent_connection = google_cloudbuildv2_connection.github_update.name
  remote_uri        = "https://github.com/gcb-repos-robot/tf-demo.git"
  annotations       = {}
  location          = "%{region}"
  project           = "%{project_name}"
  provider          = google-beta
}
resource "google_cloudbuildv2_connection" "github_update" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  annotations = {
    otherkey = "othervalue"

    somekey = "somevalue"
  }

  disabled = false

  github_config {
    app_installation_id = 31300675

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/latest"
    }
  }

  project = "%{project_name}"
  provider = google-beta
}

`, context)
}

func testAccCheckCloudbuildv2RepositoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_cloudbuildv2_repository" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &cloudbuildv2.Repository{
				Name:       dcl.String(rs.Primary.Attributes["name"]),
				Connection: dcl.String(rs.Primary.Attributes["parent_connection"]),
				RemoteUri:  dcl.String(rs.Primary.Attributes["remote_uri"]),
				Location:   dcl.StringOrNil(rs.Primary.Attributes["location"]),
				Project:    dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime: dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				Etag:       dcl.StringOrNil(rs.Primary.Attributes["etag"]),
				UpdateTime: dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
			}

			client := NewDCLCloudbuildv2Client(config, config.UserAgent, billingProject, 0)
			_, err := client.GetRepository(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_cloudbuildv2_repository still exists %v", obj)
			}
		}
		return nil
	}
}

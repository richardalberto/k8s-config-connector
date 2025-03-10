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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayGatewayIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayGatewayIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccApiGatewayGatewayIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayGatewayIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccApiGatewayGatewayIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayGatewayIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayGatewayIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccApiGatewayGatewayIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccApiGatewayGatewayIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-my-gateway%{random_suffix}"
}

resource "google_api_gateway_gateway_iam_member" "foo" {
  provider = google-beta
  project = google_api_gateway_gateway.api_gw.project
  region = google_api_gateway_gateway.api_gw.region
  gateway = google_api_gateway_gateway.api_gw.gateway_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccApiGatewayGatewayIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-my-gateway%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_api_gateway_gateway_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_gateway.api_gw.project
  region = google_api_gateway_gateway.api_gw.region
  gateway = google_api_gateway_gateway.api_gw.gateway_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayGatewayIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-my-gateway%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_api_gateway_gateway_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_gateway.api_gw.project
  region = google_api_gateway_gateway.api_gw.region
  gateway = google_api_gateway_gateway.api_gw.gateway_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayGatewayIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-my-gateway%{random_suffix}"
}

resource "google_api_gateway_gateway_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_gateway.api_gw.project
  region = google_api_gateway_gateway.api_gw.region
  gateway = google_api_gateway_gateway.api_gw.gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccApiGatewayGatewayIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-my-gateway%{random_suffix}"
}

resource "google_api_gateway_gateway_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_gateway.api_gw.project
  region = google_api_gateway_gateway.api_gw.region
  gateway = google_api_gateway_gateway.api_gw.gateway_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

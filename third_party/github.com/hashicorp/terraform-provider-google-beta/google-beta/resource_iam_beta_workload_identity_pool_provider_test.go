package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIAMBetaWorkloadIdentityPoolProvider_aws(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_full(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_enabled(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_aws_basic(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_oidc(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_full(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_update(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_oidc_basic(context),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool_provider.my_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_full(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "AWS identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "attribute.aws_role==\"arn:aws:sts::999999999999:assumed-role/stack-eu-central-1-lambdaRole\""
  attribute_mapping                  = {
    "google.subject"        = "assertion.arn"
    "attribute.aws_account" = "assertion.account"
    "attribute.environment" = "assertion.arn.contains(\":instance-profile/Production\") ? \"prod\" : \"test\""
  }
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_enabled(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "AWS identity pool provider for automated test"
  disabled                           = false
  attribute_condition                = "attribute.aws_role==\"arn:aws:sts::999999999999:assumed-role/stack-eu-central-1-lambdaRole\""
  attribute_mapping                  = {
    "google.subject"        = "assertion.arn"
    "attribute.aws_account" = "assertion.account"
    "attribute.environment" = "assertion.arn.contains(\":instance-profile/Production\") ? \"prod\" : \"test\""
  }
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_full(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-full"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation-update", "example.com/gcp-oidc-federation-update"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-update"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_aws_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolProvider_oidc_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "my_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.my_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "my-provider-%{random_suffix}"
  attribute_mapping                  = {
	"google.subject"                  = "assertion.sub"
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id-full"
  }
}
`, context)
}

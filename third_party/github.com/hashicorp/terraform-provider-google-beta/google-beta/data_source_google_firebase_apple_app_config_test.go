package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceGoogleFirebaseAppleAppConfig(t *testing.T) {
	// TODO: https://github.com/hashicorp/terraform-provider-google/issues/14158
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   GetTestProjectFromEnv(),
		"bundle_id":    "apple.app." + RandString(t, 5),
		"display_name": "tf-test Display Name AppleAppConfig DataSource",
		"app_store_id": 12345,
		"team_id":      1234567890,
	}

	VcrTest(t, resource.TestCase{
		PreCheck: func() { AccTestPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				Config: testAccDataSourceGoogleFirebaseAppleAppConfig(context),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceFirebaseAppleAppConfigCheck("data.google_firebase_apple_app_config.my_app_config"),
				),
			},
			{
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Config:                   testAccDataSourceGoogleFirebaseAppleAppConfig(context),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceFirebaseAppleAppConfigCheck("data.google_firebase_apple_app_config.my_app_config"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseAppleAppConfig(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_apple_app" "my_app_config" {
  project = "%{project_id}"
  bundle_id = "%{bundle_id}"
  display_name = "%{display_name}"
  app_store_id = "%{app_store_id}"
  team_id = "%{team_id}"
}

data "google_firebase_apple_app_config" "my_app_config" {
  app_id = google_firebase_apple_app.my_app_config.app_id
}
`, context)
}

func testAccDataSourceFirebaseAppleAppConfigCheck(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", datasourceName)
		}

		if ds.Primary.Attributes["config_filename"] == "" {
			return fmt.Errorf("config filename not found in data source")
		}

		if ds.Primary.Attributes["config_file_contents"] == "" {
			return fmt.Errorf("config file contents not found in data source")
		}

		return nil
	}
}

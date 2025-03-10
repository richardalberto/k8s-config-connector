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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotebooksRuntimeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntimeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_runtime_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/runtimes/%s roles/viewer", GetTestProjectFromEnv(), "us-central1", fmt.Sprintf("tf-test-notebooks-runtime%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccNotebooksRuntimeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_runtime_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/runtimes/%s roles/viewer", GetTestProjectFromEnv(), "us-central1", fmt.Sprintf("tf-test-notebooks-runtime%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNotebooksRuntimeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccNotebooksRuntimeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_runtime_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/runtimes/%s roles/viewer user:admin@hashicorptest.com", GetTestProjectFromEnv(), "us-central1", fmt.Sprintf("tf-test-notebooks-runtime%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNotebooksRuntimeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntimeIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_runtime_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/runtimes/%s", GetTestProjectFromEnv(), "us-central1", fmt.Sprintf("tf-test-notebooks-runtime%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebooksRuntimeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_notebooks_runtime_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/runtimes/%s", GetTestProjectFromEnv(), "us-central1", fmt.Sprintf("tf-test-notebooks-runtime%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNotebooksRuntimeIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}

resource "google_notebooks_runtime_iam_member" "foo" {
  project = google_notebooks_runtime.runtime.project
  location = google_notebooks_runtime.runtime.location
  runtime_name = google_notebooks_runtime.runtime.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccNotebooksRuntimeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_notebooks_runtime_iam_policy" "foo" {
  project = google_notebooks_runtime.runtime.project
  location = google_notebooks_runtime.runtime.location
  runtime_name = google_notebooks_runtime.runtime.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNotebooksRuntimeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_notebooks_runtime_iam_policy" "foo" {
  project = google_notebooks_runtime.runtime.project
  location = google_notebooks_runtime.runtime.location
  runtime_name = google_notebooks_runtime.runtime.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNotebooksRuntimeIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}

resource "google_notebooks_runtime_iam_binding" "foo" {
  project = google_notebooks_runtime.runtime.project
  location = google_notebooks_runtime.runtime.location
  runtime_name = google_notebooks_runtime.runtime.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccNotebooksRuntimeIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}

resource "google_notebooks_runtime_iam_binding" "foo" {
  project = google_notebooks_runtime.runtime.project
  location = google_notebooks_runtime.runtime.location
  runtime_name = google_notebooks_runtime.runtime.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_bigquery_dataset" "test" {
	dataset_id = "dataset_id"
}

resource "google_bigquery_routine" "sproc" {
  dataset_id      = google_bigquery_dataset.test.dataset_id
  routine_id      = "tf_test_routine_id%{random_suffix}"
  routine_type    = "TABLE_VALUED_FUNCTION"
  language        = "SQL"
  definition_body = <<-EOS
    SELECT 1 + value AS value
  EOS
  arguments {
    name          = "value"
    argument_kind = "FIXED_TYPE"
    data_type     = jsonencode({ "typeKind" : "INT64" })
  }
  return_table_type = jsonencode({"columns" : [
    { "name" : "value", "type" : { "typeKind" : "INT64" } },
  ] })
}
```

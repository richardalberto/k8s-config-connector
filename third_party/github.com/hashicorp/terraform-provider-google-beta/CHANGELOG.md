## 4.62.0 (Unreleased)

FEATURES:
* **New Data Source:** `google_compute_region_instance_template` ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* **New Resource:** `google_compute_region_instance_template` ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* **New Resource:** `google_logging_linked_dataset` ([#5459](https://github.com/hashicorp/terraform-provider-google-beta/pull/5459))

IMPROVEMENTS:
* cloudasset: added `OS_INVENTORY` value to `content_type` for `google_cloud_asset_*_feed` ([#5465](https://github.com/hashicorp/terraform-provider-google-beta/pull/5465))
* clouddeploy: added canary deployment fields for resource `google_clouddeploy_delivery_pipeline` ([#5451](https://github.com/hashicorp/terraform-provider-google-beta/pull/5451))
* compute: supported region instance template in`source_instance_template` field of `google_compute_instance_from_template` resource ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* container: added `pod_cidr_overprovision_config` field to `google_container_cluster` and  `google_container_node_pool` resources. ([#5468](https://github.com/hashicorp/terraform-provider-google-beta/pull/5468))
* orgpolicy: accepted variable cases for booleans such as true, True, and TRUE in `google_org_policy_policy` ([#5443](https://github.com/hashicorp/terraform-provider-google-beta/pull/5443))

BUG FIXES:
* cloudidentity: fixed immutability issue on `initialGroupConfig` field for resource `google_cloud_identity_group` ([#5456](https://github.com/hashicorp/terraform-provider-google-beta/pull/5456))
* provider: fixed an error resulting from leaving `batching.send_after` unspecified and `batching` specified ([#5460](https://github.com/hashicorp/terraform-provider-google-beta/pull/5460))
* provider: fixed bug where `credentials` field could not be set as an empty string ([#5466](https://github.com/hashicorp/terraform-provider-google-beta/pull/5466))
* vertex: increased the default timeout for `google_vertex_ai_index` to 180m ([#5450](https://github.com/hashicorp/terraform-provider-google-beta/pull/5450))

## 4.61.0 (April 10, 2023)

BREAKING CHANGES:
* cloudrunv2: set a default value of 3 for `max_retries` in `google_cloud_run_v2_job`. This should match the API's existing default, but may show a diff at plan time in limited circumstances as drift is now detected ([#5432](https://github.com/hashicorp/terraform-provider-google-beta/pull/5432))

FEATURES:
* **New Data Source:** `google_firebase_android_app_config` ([#5425](https://github.com/hashicorp/terraform-provider-google-beta/pull/5425))
* **New Resource:** `google_apigee_keystores_aliases_pkcs12` ([#5411](https://github.com/hashicorp/terraform-provider-google-beta/pull/5411))
* **New Resource:** `google_apigee_keystores_aliases_self_signed_cert` ([#5394](https://github.com/hashicorp/terraform-provider-google-beta/pull/5394))
* **New Resource:** `google_network_security_url_lists` ([#5439](https://github.com/hashicorp/terraform-provider-google-beta/pull/5439))
* **New Resource:** `google_network_services_mesh` ([#5393](https://github.com/hashicorp/terraform-provider-google-beta/pull/5393))
* **New Resource:** `google_network_security_gateway_security_policy` (beta) ([#5434](https://github.com/hashicorp/terraform-provider-google-beta/pull/5434))
* **New Resource:** `google_network_security_gateway_security_policy_rule` (beta) ([#5434](https://github.com/hashicorp/terraform-provider-google-beta/pull/5434))

IMPROVEMENTS:
* alloydb: added update support for `initial_user` and `automated_backup_policy.weekly_schedule` to `google_alloydb_cluster` ([#5420](https://github.com/hashicorp/terraform-provider-google-beta/pull/5420))
* artifactregistry: added support for tag immutability ([#5427](https://github.com/hashicorp/terraform-provider-google-beta/pull/5427))
* artifactregistry: promoted `mode`, `virtual_repository_config`, and `remote_repository_config` to GA ([#5426](https://github.com/hashicorp/terraform-provider-google-beta/pull/5426))
* bigqueryreservation: added `edition` and `autoscale` to `google_bigquery_reservation` and `edition` to `bigquery_capacity_commitment` ([#5399](https://github.com/hashicorp/terraform-provider-google-beta/pull/5399))
* compute: added support for `SEV_LIVE_MIGRATABLE` to `guest_os_features.type` in `google_compute_image` ([#5424](https://github.com/hashicorp/terraform-provider-google-beta/pull/5424))
* compute: added support for `stack_type` to `google_compute_ha_vpn_gateway` ([#5395](https://github.com/hashicorp/terraform-provider-google-beta/pull/5395))
* container: added support for `ephemeral_storage_local_ssd_config` to `google_container_cluster.node_config`, `google_container_cluster.node_pools.node_config`, `google_container_node_pool.node_config` ([#5400](https://github.com/hashicorp/terraform-provider-google-beta/pull/5400))
* dlp: Changed `dictionary`, `regex`, `regex.group_indexes` and `large_custom_dictionary` fields in `google_data_loss_prevention_stored_info_type` to be update-in-place ([#5428](https://github.com/hashicorp/terraform-provider-google-beta/pull/5428))
* logging: added support for `disabled` to `google_logging_metric` ([#5423](https://github.com/hashicorp/terraform-provider-google-beta/pull/5423))
* networkservices: increased the max count for `route_rule` to 200 on `google_network_services_edge_cache_service` ([#5433](https://github.com/hashicorp/terraform-provider-google-beta/pull/5433))
* storagetransfer: added support for 'last_modified_since' and 'last_modified_before' fields to 'google_storage_transfer_job' resource ([#5398](https://github.com/hashicorp/terraform-provider-google-beta/pull/5398))

BUG FIXES:
* bigquery: fixed the import logic in `google_bigquery_capacity_commitment` ([#5435](https://github.com/hashicorp/terraform-provider-google-beta/pull/5435))
* cloudrunv2: fixed the bug where setting `max_retries` to 0 in `google_cloud_run_v2_job` was not respected. ([#5432](https://github.com/hashicorp/terraform-provider-google-beta/pull/5432))
* container: fixed a bug creating a diff adding a `stack_type` when GKE omitted `stackType` in API responses from older GKE clusters ([#5429](https://github.com/hashicorp/terraform-provider-google-beta/pull/5429))
* dataproc: fixed validation of `optional_components` ([#5410](https://github.com/hashicorp/terraform-provider-google-beta/pull/5410))
* provider: fixed an issue where the `USER_PROJECT_OVERRIDE` environment variable was not being read ([#5441](https://github.com/hashicorp/terraform-provider-google-beta/pull/5441))
* provider: fixed an issue where the provider crashed when "batching" was set in `4.60.0`/`4.60.1` ([#5440](https://github.com/hashicorp/terraform-provider-google-beta/pull/5440))

## 4.60.2 (April 6, 2023)

BUG FIXES:
* provider: fixed an issue where the provider crashed when "batching" was set in `4.60.0`/`4.60.1`
* provider: fixed an issue where the `USER_PROJECT_OVERRIDE` environment variable was not being read

## 4.60.1 (April 5, 2023)

BUG FIXES:
* container: fixed a bug creating a diff adding a `stack_type` when GKE omitted `stackType` in API responses from older GKE clusters

## 4.60.0 (April 4, 2023)

FEATURES:
* **New Resource:** `google_apigee_keystores_aliases_key_cert_file` ([#5386](https://github.com/hashicorp/terraform-provider-google-beta/pull/5386))

IMPROVEMENTS:
* compute: added `address_type`, `network`, `network_tier`, `prefix_length`, `purpose`, `subnetwork` and `users` field for `google_compute_address` and `google_compute_global_address` datasource ([#5363](https://github.com/hashicorp/terraform-provider-google-beta/pull/5363))
* compute: added `network_firewall_policy_enforcement_order` field to `google_compute_network` resource ([#5375](https://github.com/hashicorp/terraform-provider-google-beta/pull/5375))
* compute: added output-only attribute `self_link_unique` for `google_compute_instance_template` to point to the unique id of the resource instead of its name ([#5384](https://github.com/hashicorp/terraform-provider-google-beta/pull/5384))
* container: added `stack_type` field to `google_container_cluster` resource ([#5364](https://github.com/hashicorp/terraform-provider-google-beta/pull/5364))
* container: added `advanced_machine_features` field to `google_container_cluster` resource ([#5371](https://github.com/hashicorp/terraform-provider-google-beta/pull/5371))
* networkservice: updated the max number of `host_rule` on `google_network_services_edge_cache_service` ([#5376](https://github.com/hashicorp/terraform-provider-google-beta/pull/5376))
* sql: added support of single-database-recovery for SQL Server PITR with `database_names` attribute to `google_sql_instance` ([#5366](https://github.com/hashicorp/terraform-provider-google-beta/pull/5366))

BUG FIXES:
* cloudrun: fixed race condition when polling for status during an update of a `google_cloud_run_service` ([#5365](https://github.com/hashicorp/terraform-provider-google-beta/pull/5365))
* cloudsql: fixed the error in any subsequent apply on `google_sql_user` after its `google_sql_database_instance` is deleted ([#5369](https://github.com/hashicorp/terraform-provider-google-beta/pull/5369))
* datacatalog: fixed `google_data_catalog_tag` only allowing 10 tags by increasing the page size to 1000 ([#5362](https://github.com/hashicorp/terraform-provider-google-beta/pull/5362))
* firebase: fixed `google_firebase_project` to succeed on apply when the project already has firebase enabled ([#5379](https://github.com/hashicorp/terraform-provider-google-beta/pull/5379))


## 4.59.0 (March 28, 2023)

FEATURES:
* **New Resource:** `google_dataplex_asset_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_dataplex_lake_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_dataplex_zone_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_network_services_gateway` ([#5355](https://github.com/hashicorp/terraform-provider-google-beta/pull/5355))

IMPROVEMENTS:
* auth: added support for oauth2 token exchange over mTLS ([#5343](https://github.com/hashicorp/terraform-provider-google-beta/pull/5343))
* bigquery: added `is_case_insensitive` and `default_collation` fields to `google_bigquery_dataset` resource ([#5342](https://github.com/hashicorp/terraform-provider-google-beta/pull/5342))
* compute: added `scratch_disk.size` field on `google_compute_instance` ([#5358](https://github.com/hashicorp/terraform-provider-google-beta/pull/5358))
* compute: added 3000 as allowable value for `disk_size_gb` for SCRATCH disks in `google_compute_instance_template` ([#5358](https://github.com/hashicorp/terraform-provider-google-beta/pull/5358))
* compute: added `WEIGHED_MAGLEV` to `locality_lb_policy` enum for backend service resources ([#5353](https://github.com/hashicorp/terraform-provider-google-beta/pull/5353))
* container: added `local_nvme_ssd_block` to `node_config` block in the `google_container_node_pool` ([#5335](https://github.com/hashicorp/terraform-provider-google-beta/pull/5335))
* logging: added `enable_analytics` field to `google_logging_project_bucket_config` ([#5347](https://github.com/hashicorp/terraform-provider-google-beta/pull/5347))
* networkservices: updated max allowed items to 25 for `expose_headers`, `allow_headers`, `request_header_to_remove`, `request_header_to_add`, `response_header_to_add` and `response_header_to_remove` of `google_network_services_edge_cache_service` ([#5346](https://github.com/hashicorp/terraform-provider-google-beta/pull/5346))
* networkservices: updated max allowed items to 25 for `request_headers_to_add` of `google_network_services_edge_cache_origin` ([#5346](https://github.com/hashicorp/terraform-provider-google-beta/pull/5346))

BUG FIXES:
* certificatemanager: fixed `managed.dns_authorizations` not being included during import of `google_certificate_manager_certificate` ([#5325](https://github.com/hashicorp/terraform-provider-google-beta/pull/5325))
* certificatemanager: fixed a bug where modifying non-updatable fields `hostname` and `matcher` in `google_certificate_manager_certificate_map_entry` would fail with API errors; now updating them will recreate the resource ([#5327](https://github.com/hashicorp/terraform-provider-google-beta/pull/5327))
* compute: fixed bug where `enforce_on_key_name` could not be unset on `google_compute_security_policy` ([#5326](https://github.com/hashicorp/terraform-provider-google-beta/pull/5326))
* datastream: fixed bug where field `dataset_id` could not utilize the id from bigquery directly ([#5331](https://github.com/hashicorp/terraform-provider-google-beta/pull/5331))
* workstations: fixed permadiff on `service_account` of `google_workstations_workstation_config` ([#5323](https://github.com/hashicorp/terraform-provider-google-beta/pull/5323))

## 4.58.0 (March 21, 2023)

FEATURES:
* **New Resource:** `google_apigee_sharedflow` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))
* **New Resource:** `google_apigee_sharedflow_deployment` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))
* **New Resource:** `google_apigee_flowhook` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))

IMPROVEMENTS:
* datafusion: added support for `accelerators` field to `google_datafusion_instance` resource. ([#5304](https://github.com/hashicorp/terraform-provider-google-beta/pull/5304))
* privateca: added support for X.509 name constraints to `google_privateca_pool`, `google_privateca_certificate`, and `google_privateca_certificate_authority` ([#5317](https://github.com/hashicorp/terraform-provider-google-beta/pull/5317))

BUG FIXES:
* alloydb: fixed permadiff on `automated_backup_policy.weekly_schedule` of `google_alloydb_cluster` ([#5305](https://github.com/hashicorp/terraform-provider-google-beta/pull/5305))
* bigquery: fixed a permadiff when `friendly_name` is removed from `google_bigquery_dataset` ([#5319](https://github.com/hashicorp/terraform-provider-google-beta/pull/5319))
* redis: fixed a bug causing diff detection on `reserved_ip_range` in `google_redis_instance` ([#5310](https://github.com/hashicorp/terraform-provider-google-beta/pull/5310))

## 4.57.0 (March 13, 2023)

FEATURES:
* **New Resource:** `google_access_context_manager_authorized_orgs_desc` ([#5292](https://github.com/hashicorp/terraform-provider-google-beta/pull/5292))
* **New Resource:** `google_bigquery_capacity_commitment` ([#5282](https://github.com/hashicorp/terraform-provider-google-beta/pull/5282))
* **New Resource:** `google_workstations_workstation` ([#5273](https://github.com/hashicorp/terraform-provider-google-beta/pull/5273))
* **New Resource:** `google_apigee_env_keystore` ([#5267](https://github.com/hashicorp/terraform-provider-google-beta/pull/5267))
* **New Resource:** `google_apigee_env_references` ([#5267](https://github.com/hashicorp/terraform-provider-google-beta/pull/5267))

IMPROVEMENTS:
* artifactregistry: added field `virtual_repository_config` and `remote_repository_config` to `google_artifact_registry_repository` ([#5289](https://github.com/hashicorp/terraform-provider-google-beta/pull/5289))
* bigquerydatapolicy: updated api version from v1beta1 to v1 and made it possible to use additional data policies. ([#5291](https://github.com/hashicorp/terraform-provider-google-beta/pull/5291))
* compute: added `maintenance_interval` field to `google_compute_instance_template` and `google_compute_instance` ([#5297](https://github.com/hashicorp/terraform-provider-google-beta/pull/5297))

BUG FIXES:
* cloudidentity: fixed an issue on `google_cloud_identity_group` `initial_group_config` field when importing ([#5266](https://github.com/hashicorp/terraform-provider-google-beta/pull/5266))
* compute: fixed the error of invalid value for field `failover_policy` when UDP is selected on `google_compute_region_backend_service` ([#5280](https://github.com/hashicorp/terraform-provider-google-beta/pull/5280))
* firebase: allowed specifying a `project` field on datasources for `google_firebase_android_app`, `google_firebase_web_app`, and `google_firebase_apple_app`. ([#5293](https://github.com/hashicorp/terraform-provider-google-beta/pull/5293))
* tags: fixed a bug preventing use of `google_tags_location_tag_binding` with zonal parent resources ([#5269](https://github.com/hashicorp/terraform-provider-google-beta/pull/5269))

## 4.56.0 (March 6, 2023)

DEPRECATIONS:
* gkehub: deprecated `mesh.control_plane` in `google_gke_hub_feature_membership`. Use `mesh.management` instead ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))

FEATURES:
* **New Resource:** google_scc_mute_config ([#5241](https://github.com/hashicorp/terraform-provider-google-beta/pull/5241))
* **New Resource:** google_workstations_workstation_config ([#5246](https://github.com/hashicorp/terraform-provider-google-beta/pull/5246))

IMPROVEMENTS:
* cloudbuild: added `peered_network_ip_range` field to `google_cloudbuild_worker_pool` resource ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* clouddeploy: added `multi_target` field to `google_clouddeploy_target` resource ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* cloudrun: added `template.0.containers0.liveness_probe.grpc`, `template.0.containers0.startup_probe.grpc` fields to `google_cloud_run_v2_service` resource ([#5259](https://github.com/hashicorp/terraform-provider-google-beta/pull/5259))
* compute: added `internal_ip` and `external_ip` to `google_compute_per_instance_config` and `google_compute_region_per_instance_config` (beta) ([#5256](https://github.com/hashicorp/terraform-provider-google-beta/pull/5256))
* compute: added `max_distance` field to `resource-policy` resource ([#5257](https://github.com/hashicorp/terraform-provider-google-beta/pull/5257))
* compute: added field `deletion_policy` to resource `google_compute_shared_vpc_service_project` ([#5243](https://github.com/hashicorp/terraform-provider-google-beta/pull/5243))
* container: added field `protect_config` to `google_container_cluster` (beta) ([#5255](https://github.com/hashicorp/terraform-provider-google-beta/pull/5255))
* containerazure: added `azure_services_authentication` to `google_container_azure_cluster` ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* networkservices: increased maximum `allow_origins` from 5 to 25 on `network_services_edge_cache_service` ([#5239](https://github.com/hashicorp/terraform-provider-google-beta/pull/5239))
* storagetransfer: added general field `sink_agent_pool_name` and `source_agent_pool_name` to `google_storage_transfer_job` ([#5262](https://github.com/hashicorp/terraform-provider-google-beta/pull/5262))

BUG FIXES:
* artifactregistry: fixed an issue where `google-beta` used an outdated beta API rather than the GA service API. New format values like "KFP" will now be accepted by both providers. ([#5247](https://github.com/hashicorp/terraform-provider-google-beta/pull/5247))
* cloudfunctions: fixed no diff found on `event_trigger.resource` of `google_cloudfunctions_function` ([#5261](https://github.com/hashicorp/terraform-provider-google-beta/pull/5261))
* dataproc: fixed an issue where `master_config.num_instances` would not force recreation when changed in `google_dataproc_cluster` ([#5251](https://github.com/hashicorp/terraform-provider-google-beta/pull/5251))
* spanner: fixed the error when updating `deletion_protection` on `google_spanner_database` ([#5242](https://github.com/hashicorp/terraform-provider-google-beta/pull/5242))
* spanner: fixed the error when updating `force_destroy` on `google_spanner_instance` ([#5242](https://github.com/hashicorp/terraform-provider-google-beta/pull/5242))

## 4.55.0 (February 27, 2023)

FEATURES:
* **New Resource:** `google_cloudbuild_bitbucket_server_config` ([#5218](https://github.com/hashicorp/terraform-provider-google-beta/pull/5218))
* **New Resource:** `google_firebase_hosting_release` ([#5229](https://github.com/hashicorp/terraform-provider-google-beta/pull/5229))
* **New Resource:** `google_firebase_hosting_version` ([#5229](https://github.com/hashicorp/terraform-provider-google-beta/pull/5229))

IMPROVEMENTS:
* container: added support for `node_config.kubelet_config.pod_pids_limit` on `google_container_node_pool` ([#5217](https://github.com/hashicorp/terraform-provider-google-beta/pull/5217))
* storage: changed the default create timeout of `google_storage_bucket` to 10m from 4m ([#5222](https://github.com/hashicorp/terraform-provider-google-beta/pull/5222))

BUG FIXES:
* container: fixed a crash when leaving `placement_policy` blank on `google_container_node_pool` ([#5233](https://github.com/hashicorp/terraform-provider-google-beta/pull/5233))
* serviceusage: removed unneeded Optional schema behaviour from `email` field in `google_project_service_identity` ([#5226](https://github.com/hashicorp/terraform-provider-google-beta/pull/5226))

## 4.54.0 (February 23, 2023)

FEATURES:
* **New Data Source:** `google_firebase_hosting_channel` ([#5188](https://github.com/hashicorp/terraform-provider-google-beta/pull/5188))
* **New Data Source:** `google_logging_sink` ([#5207](https://github.com/hashicorp/terraform-provider-google-beta/pull/5207))
* **New Data Source:** `google_sql_databases` ([#5204](https://github.com/hashicorp/terraform-provider-google-beta/pull/5204))

IMPROVEMENTS:
* cloudbuild: added `bitbucket_server_trigger_config` field to `google_cloudbuild_trigger` resource ([#5198](https://github.com/hashicorp/terraform-provider-google-beta/pull/5198))
* cloudbuild: added `github.enterprise_config_resource_name` field to `google_cloudbuild_trigger` resource ([#5205](https://github.com/hashicorp/terraform-provider-google-beta/pull/5205))
* compute: added field `rsa_encrypted_key` to `google_compute_disk` resource ([#5187](https://github.com/hashicorp/terraform-provider-google-beta/pull/5187))
* compute: added additional fields to `rules[].rate_limit_options` in `google_compute_security_policy` to support Cloud Armor Rate Limit Options ([#5193](https://github.com/hashicorp/terraform-provider-google-beta/pull/5193))
* sql: added replica promotion support to `google_sql_database_instance`. This change will allow users to promote read replica as stand alone primary instance. ([#5184](https://github.com/hashicorp/terraform-provider-google-beta/pull/5184))

BUG FIXES:
* bigquery: fixed permadiff on `max_time_travel_hours` of `google_bigquery_dataset` ([#5190](https://github.com/hashicorp/terraform-provider-google-beta/pull/5190))
* compute: added possibility to remove `stateful_disk` in `compute_instance_group_manager` and `compute_region_instance_group_manager`. ([#5203](https://github.com/hashicorp/terraform-provider-google-beta/pull/5203))
* sql: fixed an issue with updating the `settings.activation_policy` field in `google_sql_database_instance` ([#5202](https://github.com/hashicorp/terraform-provider-google-beta/pull/5202))

## 4.53.1 (February 14, 2023)

BUG FIXES:
* provider: fixed crash when trying to configure the provider with invalid credentials

## 4.53.0 (February 13, 2023)

FEATURES:
* **New Resource:** `google_apigee_addons_config` ([#5171](https://github.com/hashicorp/terraform-provider-google-beta/pull/5171))
* **New Resource:** `google_cloudbuildv2_connection_iam_binding` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_cloudbuildv2_connection_iam_member` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_cloudbuildv2_connection_iam_policy` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_firestore_database` ([#5181](https://github.com/hashicorp/terraform-provider-google-beta/pull/5181))
* **New Resource:** `google_workstations_workstation_cluster` ([#5154](https://github.com/hashicorp/terraform-provider-google-beta/pull/5154))

IMPROVEMENTS:
* compute: added `resource_policies` field to `google_compute_instance_template` ([#5182](https://github.com/hashicorp/terraform-provider-google-beta/pull/5182))
* compute: added field `force_update_on_repair` to `instance_lifecycle_policy` ([#5172](https://github.com/hashicorp/terraform-provider-google-beta/pull/5172))
* compute: added field `instance_lifecycle_policy` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#5172](https://github.com/hashicorp/terraform-provider-google-beta/pull/5172))
* compute: added the `labels` field to the `google_compute_external_vpn_gateway` resource ([#5162](https://github.com/hashicorp/terraform-provider-google-beta/pull/5162))
* datastream: added `postgresql_source_config` & `oracle_source_config` in `google_datastream_stream` ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* datastream: added support for creating `google_datastream_stream` with `desired_state=RUNNING` ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* datastream: exposed validation errors during `google_datastream_stream` creation ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* firebase: marked `deletion_policy` as updatable without recreation on `google_firebase_android_app` and `google_firebase_apple_app` ([#5163](https://github.com/hashicorp/terraform-provider-google-beta/pull/5163))
* sql: added `enable_private_path_for_google_cloud_services` field to `google_sql_database_instance` resource ([#5177](https://github.com/hashicorp/terraform-provider-google-beta/pull/5177))
* vertex_ai: added `offline_storage_ttl_days` to `google_vertex_ai_featurestore_entitytype` resource ([#5178](https://github.com/hashicorp/terraform-provider-google-beta/pull/5178))
* vertex_ai: added `online_storage_ttl_days` to `google_vertex_ai_featurestore` resource ([#5178](https://github.com/hashicorp/terraform-provider-google-beta/pull/5178))
* vertex_ai: added the field `description` to `google_vertex_ai_featurestore_entitytype` ([#5161](https://github.com/hashicorp/terraform-provider-google-beta/pull/5161))

BUG FIXES:
* composer: fixed an issue with cleaning up environments created in an error state ([#5164](https://github.com/hashicorp/terraform-provider-google-beta/pull/5164))
* compute: fixed wrong maximum limit description for possible VPC MTUs ([#5180](https://github.com/hashicorp/terraform-provider-google-beta/pull/5180))
* datafusion: fixed `version` can't be updated on `google_data_fusion_instance` ([#5175](https://github.com/hashicorp/terraform-provider-google-beta/pull/5175))

## 4.52.0 (February 6, 2023)

FEATURES:
* **New Data Source:** `google_secret_manager_secret_version_access` ([#5147](https://github.com/hashicorp/terraform-provider-google-beta/pull/5147))
* **New Resource:** `google_cloudbuildv2_connection` ([#5140](https://github.com/hashicorp/terraform-provider-google-beta/pull/5140))
* **New Resource:** `google_cloudbuildv2_repository` ([#5140](https://github.com/hashicorp/terraform-provider-google-beta/pull/5140))
* **New Resource:** `google_workstations_workstation_cluster` ([#5154](https://github.com/hashicorp/terraform-provider-google-beta/pull/5154))

IMPROVEMENTS:
* bigquery: added support for federated Azure identities to BigQuery Omni connections. ([#5150](https://github.com/hashicorp/terraform-provider-google-beta/pull/5150))
* bigquery: added `cloud_spanner.use_serverless_analytics` field ([#5139](https://github.com/hashicorp/terraform-provider-google-beta/pull/5139))
* bigquery: added `cloud_sql.service_account_id` and `azure.identity` output fields ([#5139](https://github.com/hashicorp/terraform-provider-google-beta/pull/5139))
* cloudbuild: added field `repository_event_config` to resource `trigger` ([#5142](https://github.com/hashicorp/terraform-provider-google-beta/pull/5142))
* compute: added `locality_lb_policies` field to `google_compute_backend_service` ([#5146](https://github.com/hashicorp/terraform-provider-google-beta/pull/5146))
* sql: updated the `settings.deletion_protection_enabled` property documentation. ([#13581](https://github.com/hashicorp/terraform-provider-google-beta/pull/13581))
* sql: made `root_password` field updatable in `google_sql_database_instance` ([#5133](https://github.com/hashicorp/terraform-provider-google-beta/pull/5133))

BUG FIXES:
* cloudfunctions: updated max_instances field to take API's result as default value ([#13575](https://github.com/hashicorp/terraform-provider-google-beta/pull/13575))
* container: fixed an issue with resuming failed cluster creation ([#5136](https://github.com/hashicorp/terraform-provider-google-beta/pull/5136))
* datacatalog: fixed the import failure when the `project` is different from the default on `google_data_catalog_taxonomy` ([#5145](https://github.com/hashicorp/terraform-provider-google-beta/pull/5145))
* secretmanager: fixed incorrect required_with for topics in `google_secret_managed_secret` ([#5149](https://github.com/hashicorp/terraform-provider-google-beta/pull/5149))

## 4.51.0 (January 30, 2023)

DEPRECATIONS:
* cloudrunv2: deprecated `liveness_probe.tcp_socket` field from `google_cloud_run_v2_service` resource as it is not supported by the API and it will be removed in a future major release ([#5128](https://github.com/hashicorp/terraform-provider-google-beta/pull/5128))
* cloudrunv2: deprecated `startup_probe` and `liveness_probe` fields from `google_cloud_run_v2_job` resource as they are not supported by the API and they will be removed in a future major release ([#5118](https://github.com/hashicorp/terraform-provider-google-beta/pull/5118))

FEATURES:
* **New Resource:** `google_iam_access_boundary_policy` ([#5130](https://github.com/hashicorp/terraform-provider-google-beta/pull/5130))
* **New Resource:** `google_tags_location_tag_bindings` ([#5115](https://github.com/hashicorp/terraform-provider-google-beta/pull/5115))

IMPROVEMENTS:
* cloudbuild: added `github_enterprise_config` fields to `google_cloudbuild_trigger` resource. ([#5110](https://github.com/hashicorp/terraform-provider-google-beta/pull/5110))
* cloudrunV2: added `annotations` to `google_cloud_run_v2_service` resource ([#5108](https://github.com/hashicorp/terraform-provider-google-beta/pull/5108))
* composer: Added field `cloud_data_lineage_integration` to resource `google_composer_environment` (beta) ([#5109](https://github.com/hashicorp/terraform-provider-google-beta/pull/5109))
* compute: added `tcp_time_wait_timeout_sec` field to `google_compute_router_nat` resource ([#5123](https://github.com/hashicorp/terraform-provider-google-beta/pull/5123))
* compute: Added fields to resource `google_compute_security_policy` to support Cloud Armor Auto Deploy (beta) ([#5116](https://github.com/hashicorp/terraform-provider-google-beta/pull/5116))
* compute: added `share_settings` field to the `google_compute_node_group` resource. ([#5113](https://github.com/hashicorp/terraform-provider-google-beta/pull/5113))
* containerattached: added `deletion_policy` field to `google_container_attached_cluster` resource. ([#5121](https://github.com/hashicorp/terraform-provider-google-beta/pull/5121))
* datastream: added `customer_managed_encryption_key` and `destination_config.bigquery_destination_config.source_hierarchy_datasets.dataset_template.kms_key_name` fields to `google_datastream_stream` resource ([#5120](https://github.com/hashicorp/terraform-provider-google-beta/pull/5120))
* dlp: added `publish_findings_to_cloud_data_catalog` and `publish_summary_to_cscc` to `google_data_loss_prevention_job_trigger` resource ([#5127](https://github.com/hashicorp/terraform-provider-google-beta/pull/5127))
* sql: added point_in_time_recovery_enabled for SQLServer in `google_sql_database_instance` ([#5124](https://github.com/hashicorp/terraform-provider-google-beta/pull/5124))
* spanner: added support for IAM conditions with `google_spanner_database_iam_member` and `google_spanner_instance_iam_member` ([#5125](https://github.com/hashicorp/terraform-provider-google-beta/pull/5125))
* sql: added additional fields to `google_sql_source_representation_instance`([#5114](https://github.com/hashicorp/terraform-provider-google-beta/pull/5114))

BUG FIXES:
* bigquery: fixed bug where valid iam member values for bigquery were prevented from actuation by validation ([#5111](https://github.com/hashicorp/terraform-provider-google-beta/pull/5111))
* bigquery: fixed permadiff on `external_data_configuration.connection_id` of `google_bigquery_table` ([#5126](https://github.com/hashicorp/terraform-provider-google-beta/pull/5126))
* gke: fixed the error of Invalid address to set on `config_connector_config` of the data source `google_container_cluster` ([#5131](https://github.com/hashicorp/terraform-provider-google-beta/pull/5131))
* google_project: fixes misleading examples that could cause `firebase:enabled` label to be accidentally removed. ([#5122](https://github.com/hashicorp/terraform-provider-google-beta/pull/5122))

## 4.50.0 (January 23, 2023)

FEATURES:
* **New Data Source:** `google_compute_network_peering` ([#5092](https://github.com/hashicorp/terraform-provider-google-beta/pull/5092))
* **New Data Source:** `google_compute_router_nat` ([#5091](https://github.com/hashicorp/terraform-provider-google-beta/pull/5091))
* **New Resource:** `google_cloud_run_v2_job_iam_binding` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_job_iam_member` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_job_iam_policy` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_binding` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_member` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_policy` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_gke_backup_backup_plan_iam_binding` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))
* **New Resource:** `google_gke_backup_backup_plan_iam_member` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))
* **New Resource:** `google_gke_backup_backup_plan_iam_policy` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))

IMPROVEMENTS:
* bigquery_table - added `reference_file_schema_uri` ([#5100](https://github.com/hashicorp/terraform-provider-google-beta/pull/5100))
* billingbudget: made fields `credit_types` and `subaccounts` updatable for `google_billing_budget` ([#5087](https://github.com/hashicorp/terraform-provider-google-beta/pull/5087))
* cloudrunV2: added `annotations` to `CloudRunV2_service` resource ([#5108](https://github.com/hashicorp/terraform-provider-google-beta/pull/5108))
* composer: added `recovery_config` in `google_composer_environment` resource ([#5105](https://github.com/hashicorp/terraform-provider-google-beta/pull/5105))
* compute: added support for 'edge_security_policy' field to 'google_compute_backend_service' resource. ([#5101](https://github.com/hashicorp/terraform-provider-google-beta/pull/5101))
* compute: added `max_run_duration` field to `google_compute_instance` and `google_compute_instance_template` resource (beta) ([#5096](https://github.com/hashicorp/terraform-provider-google-beta/pull/5096))
* dataproc: added support for `dataproc_metric_config` to resource `google_dataproc_cluster` ([#5093](https://github.com/hashicorp/terraform-provider-google-beta/pull/5093))
* dlp: added all subfields under `deidentify_template.record_transformations.field_transformations.primitive_transformation` to `google_data_loss_prevention_deidentify_template` ([#5104](https://github.com/hashicorp/terraform-provider-google-beta/pull/5104))
* sql: changed the default create timeout of `google_sql_database_instance` to 40m from 30m ([#5094](https://github.com/hashicorp/terraform-provider-google-beta/pull/5094))

BUG FIXES:
* certificatemanager: removed incorrect indication that the `self_managed` field in `google_certificate_manager_certificate` was treated as sensitive, and marked `self_managed.pem_private_key` as sensitive ([#5106](https://github.com/hashicorp/terraform-provider-google-beta/pull/5106))
* cloudplatform: fixed the error with header `X-Goog-User-Project` on `google_client_openid_userinfo` ([#5090](https://github.com/hashicorp/terraform-provider-google-beta/pull/5090))
* cloudsql: fixed `disk_type` can't be updated on `google_sql_database_instance` ([#5095](https://github.com/hashicorp/terraform-provider-google-beta/pull/5095))
* vertexai: fixed updating value_type in google_vertex_ai_featurestore_entitytype_feature ([#5098](https://github.com/hashicorp/terraform-provider-google-beta/pull/5098))

## 4.49.0 (January 17, 2023)

FEATURES:
* **New Data Source:** `google_project_service` ([#5067](https://github.com/hashicorp/terraform-provider-google-beta/pull/5067))
* **New Data Source:** `google_sql_database_instances` ([#5066](https://github.com/hashicorp/terraform-provider-google-beta/pull/5066))
* **New Data Source:** `google_container_attached_install_manifest` ([#5073](https://github.com/hashicorp/terraform-provider-google-beta/pull/5073))
* **New Data Source:** `google_container_attached_install_manifest` ([#5080](https://github.com/hashicorp/terraform-provider-google-beta/pull/5080))
* **New Data Source:** `google_container_attached_versions` ([#5073](https://github.com/hashicorp/terraform-provider-google-beta/pull/5073))
* **New Resource:** `google_datastream_stream` ([#5045](https://github.com/hashicorp/terraform-provider-google-beta/pull/5045))

IMPROVEMENTS:
* android_app: added general fields `sha1_hashes`, `sha256_hashes` and `etag` to `google_firebase_android_app`. ([#5074](https://github.com/hashicorp/terraform-provider-google-beta/pull/5074))
* cloudids: added `threat_exception` field to `google_cloud_ids_endpoint` ([#5072](https://github.com/hashicorp/terraform-provider-google-beta/pull/5072))
* composer: added `triggerer` field in `google_composer_environment` ([#5055](https://github.com/hashicorp/terraform-provider-google-beta/pull/5055))
* compute: enabled deletion for `stateful_ips` fields in `instance_group_manager` and `region_instance_group_manager`. ([#5064](https://github.com/hashicorp/terraform-provider-google-beta/pull/5064))
* compute: added field `expire_time` to resource `google_compute_region_ssl_certificate` ([#5049](https://github.com/hashicorp/terraform-provider-google-beta/pull/5049))
* compute: added field `expire_time` to resource `google_compute_ssl_certificate` ([#5049](https://github.com/hashicorp/terraform-provider-google-beta/pull/5049))
* container: added `release_channel_latest_version` in `google_container_engine_versions` datasource ([#5044](https://github.com/hashicorp/terraform-provider-google-beta/pull/5044))
* container: added `google_container_aws_node_pool` `autoscaling_metrics_collection` field ([#5084](https://github.com/hashicorp/terraform-provider-google-beta/pull/5084))
* container: added update support for `google_container_aws_node_pool` `tags` field ([#5084](https://github.com/hashicorp/terraform-provider-google-beta/pull/5084))
* dataproc: added support for `node_group_affinity.` in `google_dataproc_cluster` ([#5053](https://github.com/hashicorp/terraform-provider-google-beta/pull/5053))
* dataproc: added support for `reservation_affinity` in `google_dataproc_cluster` ([#5050](https://github.com/hashicorp/terraform-provider-google-beta/pull/5050))
* dlp: added field 'identifyingFields' to 'bigQueryOptions' for creating DLP jobs. ([#5085](https://github.com/hashicorp/terraform-provider-google-beta/pull/5085))
* metastore: added `telemetry_config` field to `google_dataproc_metastore_service` ([#5065](https://github.com/hashicorp/terraform-provider-google-beta/pull/5065))
* sql: added the ability to set `point_in_time_recovery_enabled` flag in `google_sql_database_instance` for `SQLSERVER` instance, since the API supports it now. ([#5079](https://github.com/hashicorp/terraform-provider-google-beta/pull/5079))
* sql: added `instance_type` field to `google_sql_database_instance` resource ([#5057](https://github.com/hashicorp/terraform-provider-google-beta/pull/5057))
* vertexai: added `scaling` field in `google_vertex_ai_featurestore` ([#5081](https://github.com/hashicorp/terraform-provider-google-beta/pull/5081))

BUG FIXES:
* android_app: modified the `package_name` field suffix to always start with a letter in `google_firebase_android_app`. ([#5074](https://github.com/hashicorp/terraform-provider-google-beta/pull/5074))
* bigqueryconnection: fixed a bug where `aws.access_role.iam_role_id` cannot be updated on `google_bigquery_connection` ([#5083](https://github.com/hashicorp/terraform-provider-google-beta/pull/5083))
* container: fixed a bug preventing updates to `master_global_access_config` in `google_container_cluster` ([#5043](https://github.com/hashicorp/terraform-provider-google-beta/pull/5043))
* container: fixed perma-diff based on a server set taint `kubernetes.io/arch` on `google_container_node_pool` ([#5054](https://github.com/hashicorp/terraform-provider-google-beta/pull/5054))
* spanner: fixed crash when `google_spanner_database.ddl` item was nil ([#5071](https://github.com/hashicorp/terraform-provider-google-beta/pull/5071))

## 4.48.0 (January 9, 2023)

FEATURES:
* **New Data Source:** `google_beyondcorp_app_connection` ([#5025](https://github.com/hashicorp/terraform-provider-google-beta/pull/5025))
* **New Data Source:** `google_beyondcorp_app_connector` ([#5008](https://github.com/hashicorp/terraform-provider-google-beta/pull/5008))
* **New Data Source:** `google_beyondcorp_app_gateway` ([#5008](https://github.com/hashicorp/terraform-provider-google-beta/pull/5008))
* **New Data Source:** `google_cloudbuild_trigger` ([#5017](https://github.com/hashicorp/terraform-provider-google-beta/pull/5017))
* **New Data Source:** `google_compute_instance_group_manager` ([#5002](https://github.com/hashicorp/terraform-provider-google-beta/pull/5002))
* **New Data Source:** `google_firebase_apple_app_config` ([#5031](https://github.com/hashicorp/terraform-provider-google-beta/pull/5031))
* **New Data Source:** `google_firebase_apple_app` ([#4978](https://github.com/hashicorp/terraform-provider-google-beta/pull/4978))
* **New Data Source:** `google_pubsub_subscription` ([#5001](https://github.com/hashicorp/terraform-provider-google-beta/pull/5001))
* **New Data Source:** `google_sql_database` ([#5038](https://github.com/hashicorp/terraform-provider-google-beta/pull/5038))
* **New Resource:** `google_apigee_sync_authorization` ([#5015](https://github.com/hashicorp/terraform-provider-google-beta/pull/5015))
* **New Resource:** `google_beyondcorp_app_connection` ([#5012](https://github.com/hashicorp/terraform-provider-google-beta/pull/5012))
* **New Resource:** `google_container_attached_cluster` ([#5037](https://github.com/hashicorp/terraform-provider-google-beta/pull/5037))
* **New Resource:** `google_dns_managed_zone_iam_*` ([#5007](https://github.com/hashicorp/terraform-provider-google-beta/pull/5007))
* **New Resource:** `google_firebase_database_instance`([#5019](https://github.com/hashicorp/terraform-provider-google-beta/pull/5019))

IMPROVEMENTS:
* cloudfunctions2: added `available_cpu` and `max_instance_request_concurrency` to support concurrency in `google_cloudfunctions2_function` ([#5011](https://github.com/hashicorp/terraform-provider-google-beta/pull/5011))
* gkehub: added support for `configmanagement.config_sync.oci` field to `google_gke_hub_feature_membership` resource([#5013](https://github.com/hashicorp/terraform-provider-google-beta/pull/5013))
* compute: added support for local IP ranges in `google_compute_firewall` ([#4979](https://github.com/hashicorp/terraform-provider-google-beta/pull/4979))
* compute: added `router_appliance_instance` field to `google_compute_router_bgp_peer` ([#5035](https://github.com/hashicorp/terraform-provider-google-beta/pull/5035))
* compute: added support for `generated_id` field in `google_compute_backend_service` to get the value of `id` defined by the server ([#4981](https://github.com/hashicorp/terraform-provider-google-beta/pull/4981))
* compute: added support for `image_encryption_key` to `google_compute_image` ([#4989](https://github.com/hashicorp/terraform-provider-google-beta/pull/4989))
* compute: added support for `source_snapshot`, `source_snapshot_encyption_key`, and `source_image_encryption_key` to `google_compute_instance_template` ([#4989](https://github.com/hashicorp/terraform-provider-google-beta/pull/4989))
* compute: added stateful ip fields `stateful_internal_ip` and `stateful_external_ip` to `google_compute_instance_group_manager` ([#4992](https://github.com/hashicorp/terraform-provider-google-beta/pull/4992))
* container: added `gateway_api_config` block to `google_container_cluster` resource for supporting the gke gateway api controller ([#4976](https://github.com/hashicorp/terraform-provider-google-beta/pull/4976))
* container: supported in-place update for `labels` in `google_container_node_pool` ([#4998](https://github.com/hashicorp/terraform-provider-google-beta/pull/4998))
* dataproc: added support for `SPOT` option for `preemptibility` in `google_dataproc_cluster` ([#5024](https://github.com/hashicorp/terraform-provider-google-beta/pull/5024))
* dlp: added field `deidentify_config.record_transformations.field_transformations` to `google_data_loss_prevention_deidentify_template` ([#4996](https://github.com/hashicorp/terraform-provider-google-beta/pull/4996))
* dlp: added field `deidentify_config.record_transformations.record_suppressions` to `google_data_loss_prevention_deidentify_template` ([#5004](https://github.com/hashicorp/terraform-provider-google-beta/pull/5004))
* dlp: added `version` field to `google_data_loss_prevention_inspect_template` resource ([#5032](https://github.com/hashicorp/terraform-provider-google-beta/pull/5032))
* osconfig: added support for `skip_await_rollout` in `google_os_config_os_policy_assignment` ([#5026](https://github.com/hashicorp/terraform-provider-google-beta/pull/5026))
* sql: added [new deletion protection](https://cloud.google.com/sql/docs/mysql/deletion-protection) feature `deletion_protection_enabled` in `google_sql_database_instance` to guard against deletion from all surfaces ([#4987](https://github.com/hashicorp/terraform-provider-google-beta/pull/4987))
* sql: made `settings.sql_server_audit_config.bucket` field in `google_sql_database_instance` to be optional. ([#4988](https://github.com/hashicorp/terraform-provider-google-beta/pull/4988))
* storagetransfer: supported in-place update for `schedule` in `google_storage_transfer_job` ([#4993](https://github.com/hashicorp/terraform-provider-google-beta/pull/4993))

BUG FIXES:
* bigquery: fixed a permadiff on `labels` of `google_bigquery_dataset` when it is referenced in `google_dataplex_asset` ([#5022](https://github.com/hashicorp/terraform-provider-google-beta/pull/5022))
* compute: fixed a permadiff on `private_ip_google_access` of `google_compute_subnetwork` ([#4983](https://github.com/hashicorp/terraform-provider-google-beta/pull/4983))
* compute: fixed an issue where `enable_dynamic_port_allocation` was not able to set to `false` in `google_compute_router_nat` ([#4982](https://github.com/hashicorp/terraform-provider-google-beta/pull/4982))
* container: fixed a permadiff on `location_policy` of `google_container_cluster` and `google_container_node_pool` ([#4997](https://github.com/hashicorp/terraform-provider-google-beta/pull/4997))
* identityplatform: fixed issues with `google_identity_platform_config` creation  ([#5005](https://github.com/hashicorp/terraform-provider-google-beta/pull/5005))
* resourcemanager: fixed the `google_project` datasource silently returning empty results when the project was not found or not in the ACTIVE state. Now, an error will be surfaced instead. ([#5029](https://github.com/hashicorp/terraform-provider-google-beta/pull/5029))
* sql: fixed `sql_database_instance` leaking root users ([#4991](https://github.com/hashicorp/terraform-provider-google-beta/pull/4991))

## 4.47.0 (December 21, 2022)

NOTES:
* sql: fixed an issue where `google_sql_database` was abandoned by default as of version `4.45.0`. Users who have upgraded to `4.45.0` or `4.46.0` will see a diff when running their next `terraform apply` after upgrading this version, indicating the `deletion_policy` field's value has changed from `"ABANDON"` to `"DELETE"`. This will create a no-op call against the API, but can otherwise be safely applied. ([#4972](https://github.com/hashicorp/terraform-provider-google-beta/pull/4972))

FEATURES:
* **New Resource:** `google_alloydb_backup` ([#4959](https://github.com/hashicorp/terraform-provider-google-beta/pull/4959))
* **New Resource:** `google_filestore_backup` ([#4963](https://github.com/hashicorp/terraform-provider-google-beta/pull/4963))

IMPROVEMENTS:
* bigtable: added `deletion_protection` field to `google_bigtable_table` ([#4975](https://github.com/hashicorp/terraform-provider-google-beta/pull/4975))
* compute: Made `google_compute_subnetwork.ipv6_access_type` field updatable in-place ([#4965](https://github.com/hashicorp/terraform-provider-google-beta/pull/4965))
* container: added `auto_provisioning_defaults.cluster_autoscaling.upgrade_settings` in `google_container_cluster` ([#4958](https://github.com/hashicorp/terraform-provider-google-beta/pull/4958))
* container: added `gateway_api_config` block to `google_container_cluster` resource for supporting the gke gateway api controller ([#4976](https://github.com/hashicorp/terraform-provider-google-beta/pull/4976))
* datacatalog: added update support for `fields` in `google_data_catalog_tag_template` ([#4968](https://github.com/hashicorp/terraform-provider-google-beta/pull/4968))
* iam: added plan-time validation for IAM members ([#4960](https://github.com/hashicorp/terraform-provider-google-beta/pull/4960))
* logging: added `bucket_name` field to `google_logging_metric` ([#4964](https://github.com/hashicorp/terraform-provider-google-beta/pull/4964))
* logging: made `metric_descriptor` field optional for `google_logging_metric` ([#4971](https://github.com/hashicorp/terraform-provider-google-beta/pull/4971))

BUG FIXES:
* composer: fixed a crash when updating `ip_allocation_policy` of `google_composer_environment` ([#4956](https://github.com/hashicorp/terraform-provider-google-beta/pull/4956))
* sql: fixed an issue where `google_sql_database` was abandoned by default as of version `4.45.0`. Users who have upgraded to `4.45.0` or `4.46.0` will see a diff when running their next `terraform apply` after upgrading this version, indicating the `deletion_policy` field's value has changed from `"ABANDON"` to `"DELETE"`. This will create a no-op call against the API, but can otherwise be safely applied. ([#4972](https://github.com/hashicorp/terraform-provider-google-beta/pull/4972))

## 4.46.0 (December 12, 2022)

FEATURES:
* **New Data Source:** `google_firebase_android_app` ([#4955](https://github.com/hashicorp/terraform-provider-google-beta/pull/4955))
* **New Resource:** `google_cloud_run_v2_job` ([#4937](https://github.com/hashicorp/terraform-provider-google-beta/pull/4937))
* **New Resource:** `google_cloud_run_v2_service` ([#4942](https://github.com/hashicorp/terraform-provider-google-beta/pull/4942))
* **New Resource:** `google_gke_backup_backup_plan` ([#4948](https://github.com/hashicorp/terraform-provider-google-beta/pull/4948))
* **New Resource:** `google_firebase_storage_bucket` ([#4951](https://github.com/hashicorp/terraform-provider-google-beta/pull/4951))

IMPROVEMENTS:
* network_services: added `origin_override_action` and `origin_redirect` to `google_network_services_edge_cache_origin` ([#4936](https://github.com/hashicorp/terraform-provider-google-beta/pull/4936))
* bigquerydatatransfer: recreate `google_bigquery_data_transfer_config` for Cloud Storage transfers when immutable params `data_path_template` and `destination_table_name_template` are changed ([#4929](https://github.com/hashicorp/terraform-provider-google-beta/pull/4929))
* compute: added fields to resource `google_compute_security_policy` to support Cloud Armor bot management ([#4938](https://github.com/hashicorp/terraform-provider-google-beta/pull/4938))
* container: added support for concurrent node pool mutations on a cluster. Previously, node pool mutations were restricted to run synchronously clientside. NOTE: While this feature is supported in Terraform from this release onwards, only a limited number of GCP projects will support this behavior initially. The provider will automatically process mutations concurrently as the feature rolls out generally. ([#4947](https://github.com/hashicorp/terraform-provider-google-beta/pull/4947))
* metastore: added general field `network_config` to `google_dataproc_metastore_service` ([#4952](https://github.com/hashicorp/terraform-provider-google-beta/pull/4952))
* storage: added support for `autoclass` in `google_storage_bucket` resource ([#4953](https://github.com/hashicorp/terraform-provider-google-beta/pull/4953))

BUG FIXES:
* alloydb: made `machine_config.cpu_count` updatable on `google_alloydb_instance` ([#4930](https://github.com/hashicorp/terraform-provider-google-beta/pull/4930))
* composer: fixed a crash when updating `ip_allocation_policy` of `google_composer_environment` ([#4956](https://github.com/hashicorp/terraform-provider-google-beta/pull/4956))
* container: fixed GKE permadiff/thrashing when `update_settings. max_surge` or `update_settings. max_unavailable` values are updating on `google_container_node_pool` ([#4945](https://github.com/hashicorp/terraform-provider-google-beta/pull/4945))
* datastream: fixed `google_datastream_private_connection` ignoring failures during creation ([#4939](https://github.com/hashicorp/terraform-provider-google-beta/pull/4939))
* firebase: fixed permadiff on the field `deletion_policy` of `google_firebase_apple_app` ([#4954](https://github.com/hashicorp/terraform-provider-google-beta/pull/4954))
* kms: fixed issues with deleting crypto key versions in states other than ENABLED ([#4943](https://github.com/hashicorp/terraform-provider-google-beta/pull/4943))

## 4.45.0 (December 5, 2022)

FEATURES:
* **New Data Source:** `google_logging_project_cmek_settings` ([#4902](https://github.com/hashicorp/terraform-provider-google-beta/pull/4902))
* **New Resource:** `google_iam_workforce_pool_provider` ([#4922](https://github.com/hashicorp/terraform-provider-google-beta/pull/4922))
* **New Resource:** `google_vertex_ai_tensorboard` ([#4896](https://github.com/hashicorp/terraform-provider-google-beta/pull/4896))
* **New Resource:** `google_data_fusion_instance_iam_binding` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_data_fusion_instance_iam_member` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_data_fusion_instance_iam_policy` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_eventarc_google_channel_config` ([#4905](https://github.com/hashicorp/terraform-provider-google-beta/pull/4905))
* **New Resource:** `google_vertex_ai_index` ([#4923](https://github.com/hashicorp/terraform-provider-google-beta/pull/4923))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_binding` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_member` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_policy` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))


IMPROVEMENTS:
* anthos-fleet-management: added option `mesh: control_plane` to resource `google_gke_hub_feature_membership`. ([#4927](https://github.com/hashicorp/terraform-provider-google-beta/pull/4927))
* bigquerydatatransfer: made `google_bigquery_data_transfer_config` recreate for Cloud Storage transfers when immutable params `data_path_template` and `destination_table_name_template` are changed ([#4929](https://github.com/hashicorp/terraform-provider-google-beta/pull/4929))
* bigtable: added support for abandoning GC policy ([#4897](https://github.com/hashicorp/terraform-provider-google-beta/pull/4897))
* cloudsql: added `connector_enforcement` field to `google_sql_database_instance` resource ([#4894](https://github.com/hashicorp/terraform-provider-google-beta/pull/4894))
* compute: added `default_route_action.cors_policy` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.fault_injection_policy` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.timeout` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.url_rewrite` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `include_http_headers` field to the `cdn_policy` field of `google_compute_backend_service` resource ([#4912](https://github.com/hashicorp/terraform-provider-google-beta/pull/4912))
* compute: added field `list_managed_instances_results` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4903](https://github.com/hashicorp/terraform-provider-google-beta/pull/4903))
* compute: added subnetwork and private_ip_address arguments to resource_compute_router_interface ([#4914](https://github.com/hashicorp/terraform-provider-google-beta/pull/4914))
* container: added `resource_labels` field to `node_config` resource ([#4913](https://github.com/hashicorp/terraform-provider-google-beta/pull/4913))
* container: added field `enable_private_nodes` in `network_config` to `google_container_node_pool` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* container: added field `gcp_public_cidrs_access_enabled` and `private_endpoint_subnetwork` to `google_container_cluster` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* container: added update support for `enable_private_endpoint` and `enable_private_nodes` in `google_container_cluster` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* datafusion: added `api_endpoint` and `p4_service_account ` attributes to `google_data_fusion_instance` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* datafusion: added `zone`, `display_name`, `crypto_key_config`, `event_publish_config`, and `enable_rbac` args to `google_data_fusion_instance` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* logging: added `cmek_settings` field to `google_logging_project_bucket_config` resource ([#4902](https://github.com/hashicorp/terraform-provider-google-beta/pull/4902))
* sql: added 'deny_maintenance_period' field for 'google_sql_database_instance' within which 'end_date', 'start_date' and 'time' fields are present. ([#4915](https://github.com/hashicorp/terraform-provider-google-beta/pull/4915))
* sql: added field `deletion_policy` to resource `google_sql_database` ([#4916](https://github.com/hashicorp/terraform-provider-google-beta/pull/4916))

BUG FIXES:
* alloydb: fixed permdiff on `display_name` of `google_alloydb_instance` ([#4925](https://github.com/hashicorp/terraform-provider-google-beta/pull/4925))
* compute: fixed a failure in updating `most_disruptive_allowed_action` on `google_compute_per_instance_config` and `google_compute_region_per_instance_config` ([#4898](https://github.com/hashicorp/terraform-provider-google-beta/pull/4898))
* compute: fixed the error when `metadata` and `machine_type` are updated while `metadata_startup_script` was already provided on `google_compute_instance` ([#4901](https://github.com/hashicorp/terraform-provider-google-beta/pull/4901))
* container: fixed the inability to update `authenticator_groups_config` on `google_container_cluster` ([#4918](https://github.com/hashicorp/terraform-provider-google-beta/pull/4918))
* container: fixed the data source `google_container_cluster` to return an error if it does not exist ([#4900](https://github.com/hashicorp/terraform-provider-google-beta/pull/4900))
* sql: fixed `googe_sql_database_instance` to include `backup_configuration` in initial create request ([#4911](https://github.com/hashicorp/terraform-provider-google-beta/pull/4911))
* storage: fixed permdiff when `website`, `website.main_page_suffix`, `website.not_found_page` are removed on `google_storage_bucket` ([#4899](https://github.com/hashicorp/terraform-provider-google-beta/pull/4899))



## 4.44.1 (November 22, 2022)

NOTES: No changes, only released to keep this provider in sync with the `google` provider

## 4.44.0 (November 21, 2022)

FEATURES:
* **New Data Source:** `google_cloud_asset_resources_search_all` ([#4891](https://github.com/hashicorp/terraform-provider-google-beta/pull/4891))
* **New Resource:** `google_alloydb_instance` ([#4857](https://github.com/hashicorp/terraform-provider-google-beta/pull/4857))
* **New Resource:** `google_beyondcorp_app_connector` ([#4866](https://github.com/hashicorp/terraform-provider-google-beta/pull/4866))
* **New Resource:** `google_beyondcorp_app_gateway` ([#4866](https://github.com/hashicorp/terraform-provider-google-beta/pull/4866))
* **New Resource:** `google_compute_network_firewall_policy_association` ([#4868](https://github.com/hashicorp/terraform-provider-google-beta/pull/4868))
* **New Resource:** `google_compute_network_firewall_policy_rule` ([#4880](https://github.com/hashicorp/terraform-provider-google-beta/pull/4880))
* **New Resource:** `google_compute_network_firewall_policy` ([#4851](https://github.com/hashicorp/terraform-provider-google-beta/pull/4851))
* **New Resource:** `google_compute_region_network_firewall_policy_association` ([#4868](https://github.com/hashicorp/terraform-provider-google-beta/pull/4868))
* **New Resource:** `google_compute_region_network_firewall_policy_rule` ([#4880](https://github.com/hashicorp/terraform-provider-google-beta/pull/4880))
* **New Resource:** `google_compute_region_network_firewall_policy` ([#4851](https://github.com/hashicorp/terraform-provider-google-beta/pull/4851))
* **New Resource:** `google_eventarc_channel` ([#4876](https://github.com/hashicorp/terraform-provider-google-beta/pull/4876))
* **New Resource:** `google_firebase_apple_app` ([#4887](https://github.com/hashicorp/terraform-provider-google-beta/pull/4887))
* **New Resource:** `google_firebase_hosting_channel` ([#4890](https://github.com/hashicorp/terraform-provider-google-beta/pull/4890))
* **New Resource:** `google_firebase_hosting_site` ([#4846](https://github.com/hashicorp/terraform-provider-google-beta/pull/4846))
* **New Resource:** `google_identity_platform_project_default_config` ([#4853](https://github.com/hashicorp/terraform-provider-google-beta/pull/4853))
* **New Resource:** `google_kms_crypto_key_versions` ([#4831](https://github.com/hashicorp/terraform-provider-google-beta/pull/4831))
* **New Resource:** `google_storage_transfer_agent_pool` ([#4835](https://github.com/hashicorp/terraform-provider-google-beta/pull/4835))

IMPROVEMENTS:
* bigquery: supported authorized routines on resource `bigquery_dataset` and `bigquery_dataset_access` ([#4855](https://github.com/hashicorp/terraform-provider-google-beta/pull/4855))
* clouddeploy: added execution_configs.execution_timeout in target resource. ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* clouddeploy: added support for Cloud Run Targets ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* clouddeploy: added support for Deployment Verification standard strategy ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* cloudidentity: made security label settable by making labels updatable in `google_cloud_identity_groups` ([#4834](https://github.com/hashicorp/terraform-provider-google-beta/pull/4834))
* cloudrun: added field `liveness_probe.grpc` and `startup_probe.grpc` to resource `google_cloud_run_service` ([#4863](https://github.com/hashicorp/terraform-provider-google-beta/pull/4863))
* cloudsql: added `connector_enforcement` field to `google_sql_database_instance` resource ([#4894](https://github.com/hashicorp/terraform-provider-google-beta/pull/4894))
* compute: added optional `redundant_interface` argument to `google_compute_router_interface` resource ([#4881](https://github.com/hashicorp/terraform-provider-google-beta/pull/4881))
* compute: added `default_route_action.request_mirror_policy` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `default_route_action.retry_policy` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `default_route_action.weighted_backend_services` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `preconfigured_waf_config` block to `google_compute_security_policy` resource ([#4852](https://github.com/hashicorp/terraform-provider-google-beta/pull/4852))
* compute: modified machine_type field in compute instance resource to accept short name. ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* container: added `node_config.logging_variant` to `google_container_node_pool`. ([#4889](https://github.com/hashicorp/terraform-provider-google-beta/pull/4889))
* container: added `node_pool_defaults.node_config_defaults.logging_variant`, `node_pool.node_config.logging_variant`, and `node_config.logging_variant` to `google_container_cluster`. ([#4889](https://github.com/hashicorp/terraform-provider-google-beta/pull/4889))
* container: added support for Shielded Instance configuration for node auto-provisioning to `google_container_cluster` ([#4833](https://github.com/hashicorp/terraform-provider-google-beta/pull/4833))
* container: added management attribute to the google_container_cluster resource ([#4862](https://github.com/hashicorp/terraform-provider-google-beta/pull/4862))
* container: added field `blue_green_settings` to `google_container_node_pool` ([#4860](https://github.com/hashicorp/terraform-provider-google-beta/pull/4860))
* container: added field `strategy` to `google_container_node_pool` ([#4860](https://github.com/hashicorp/terraform-provider-google-beta/pull/4860))
* container: added support for additional values `APISERVER`, `CONTROLLER_MANAGER`, and `SCHEDULER` in `google_container_cluster.monitoring_config` ([#4854](https://github.com/hashicorp/terraform-provider-google-beta/pull/4854))
* datafusion: added `enable_rbac` field to `google_data_fusion_instance` resource ([#4864](https://github.com/hashicorp/terraform-provider-google-beta/pull/4864))
* dlp: added fields `rows_limit`, `rows_limit_percent`, and `sample_method` to `big_query_options` in `google_data_loss_prevention_job_trigger` ([#4856](https://github.com/hashicorp/terraform-provider-google-beta/pull/4856))
* dlp: added pubsub action to `google_data_loss_prevention_job_trigger` ([#4832](https://github.com/hashicorp/terraform-provider-google-beta/pull/4832))
* dns: added `gke_clusters` field to `google_dns_managed_zone` resource ([#4888](https://github.com/hashicorp/terraform-provider-google-beta/pull/4888))
* dns: added `gke_clusters` field to `google_dns_response_policy` resource ([#4888](https://github.com/hashicorp/terraform-provider-google-beta/pull/4888))
* eventarc: added field `channel` to `google_eventarc_trigger` ([#4876](https://github.com/hashicorp/terraform-provider-google-beta/pull/4876))
* gkehub: added `mesh` field and `management` subfield to resource `feature_membership` ([#4867](https://github.com/hashicorp/terraform-provider-google-beta/pull/4867))
* networkservices: added `aws_v4_authentication ` field to `google_network_services_edge_cache_origin ` to support S3-compatible Origins ([#4875](https://github.com/hashicorp/terraform-provider-google-beta/pull/4875))
* networkservices: added `signed_token_options` and `add_signatures` field to `google_network_services_edge_cache_service` and `validation_shared_keys` to `google_network_services_edge_cache_keyset` to support dual-token authentication ([#4884](https://github.com/hashicorp/terraform-provider-google-beta/pull/4884))
* sql: added `query_plan_per_minute` field to `insights_config` in `google_sql_database_instance` resource ([#4840](https://github.com/hashicorp/terraform-provider-google-beta/pull/4840))
* vertexai: added fields to `vertex_ai_featurestore_entitytype` to support feature value monitoring ([#4859](https://github.com/hashicorp/terraform-provider-google-beta/pull/4859))

BUG FIXES:
* apigee: fixed permadiff on consumer_accept_list for `google_apigee_instance` ([#4883](https://github.com/hashicorp/terraform-provider-google-beta/pull/4883))
* appengine: fixed permadiff on serviceaccount for 'google_app_engine_flexible_app_version' ([#4858](https://github.com/hashicorp/terraform-provider-google-beta/pull/4858))
* bigtable: updated ForceNew logic for `kms_key_name` ([#4873](https://github.com/hashicorp/terraform-provider-google-beta/pull/4873))
* bigtable: updated the error handling logic to remove the resource on resource not found error only ([#4841](https://github.com/hashicorp/terraform-provider-google-beta/pull/4841))
* billingbudget: fixed a bug where `budget_filter.credit_types_treatment` in `google_billing_budget` resource was not updating. ([#4836](https://github.com/hashicorp/terraform-provider-google-beta/pull/4836))
* cloudbuild: fixed a failure when BITBUCKET is provided for `repo_type` on `google_cloudbuild_trigger` ([#4878](https://github.com/hashicorp/terraform-provider-google-beta/pull/4878))
* cloudids: fixed `endpoint_forwarding_rule` and `endpoint_ip` attributes for `google_cloud_ids_endpoint` ([#4843](https://github.com/hashicorp/terraform-provider-google-beta/pull/4843))
* compute: fixed perma-diff on `google_compute_disk` for new amd64 images ([#4847](https://github.com/hashicorp/terraform-provider-google-beta/pull/4847))
* compute: made `target_https_proxy` possible to set `ssl_certificates` and `certificate_map` in `google_compute_target_https_proxy`  at the same time ([#4839](https://github.com/hashicorp/terraform-provider-google-beta/pull/4839))
* container: fixed a bug where `cluster_autoscaling.auto_provisioning_defaults.service_account` can not be set when `enable_autopilot = true` for `google_container_cluster` ([#4877](https://github.com/hashicorp/terraform-provider-google-beta/pull/4877))
* dialogflowcx: fixed a deployment issue for `google_dialogflow_cx_version` and `google_dialogflow_cx_environment` when they are deployed to a non-global location ([#4869](https://github.com/hashicorp/terraform-provider-google-beta/pull/4869))
* dns: fixed apply failure when `description` is set to empty string on `google_dns_managed_zone` ([#4837](https://github.com/hashicorp/terraform-provider-google-beta/pull/4837))
* provider: fixed a crash during provider authentication for certain environments ([#4892](https://github.com/hashicorp/terraform-provider-google-beta/pull/4892))
* storage: fixed a crash when `log_bucket` is updated with empty body on `google_storage_bucket` ([#4893](https://github.com/hashicorp/terraform-provider-google-beta/pull/4893))
* vertexai: made google_vertex_ai_featurestore_entitytype always use regional endpoint corresponding to parent's region ([#4845](https://github.com/hashicorp/terraform-provider-google-beta/pull/4845))

## 4.43.0 (November 7, 2022)

FEATURES:
* **New Resource:** `google_kms_crypto_key_version`([#4831](https://github.com/hashicorp/terraform-provider-google-beta/pull/4831))

## 4.42.1 (November 2, 2022)

BUG FIXES:
* storage: fixed a crash in `google_storage_bucket` when upgrading provider to version `4.42.0` with `lifecycle_rule.condition.age` unset ([#4828](https://github.com/hashicorp/terraform-provider-google-beta/pull/4828))

## 4.42.0 (October 31, 2022)

FEATURES:
* **New Data Source:** `google_compute_addresses` ([#4802](https://github.com/hashicorp/terraform-provider-google-beta/pull/4802))
* **New Data Source:** `google_compute_region_network_endpoint_group` ([#4811](https://github.com/hashicorp/terraform-provider-google-beta/pull/4811))
* **New Resource:** `google_alloydb_cluster` ([#4780](https://github.com/hashicorp/terraform-provider-google-beta/pull/4780))
* **New Resource:** `google_dataform_repository` (beta) ([#4801](https://github.com/hashicorp/terraform-provider-google-beta/pull/4801))
* **New Resource:** `google_firebase_android_app` ([#4814](https://github.com/hashicorp/terraform-provider-google-beta/pull/4814))
* **New Resource:** `google_iam_workforce_pool` ([#4818](https://github.com/hashicorp/terraform-provider-google-beta/pull/4818))
* **New Resource:** `google_monitoring_generic_service` ([#4789](https://github.com/hashicorp/terraform-provider-google-beta/pull/4789))
* **New Resource:** `google_scc_source_iam_binding` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_scc_source_iam_member` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_scc_source_iam_policy` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_vertex_ai_endpoint` ([#4815](https://github.com/hashicorp/terraform-provider-google-beta/pull/4815))
* **New Resource:** `google_vertex_ai_featurestore_iam_binding` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))
* **New Resource:** `google_vertex_ai_featurestore_iam_member` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))
* **New Resource:** `google_vertex_ai_featurestore_iam_policy` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))

IMPROVEMENTS:
* appengine: added `member` field to `google_app_engine_default_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* bigquery: added `max_time_travel_hours` field in `google_bigquery_dataset` resource ([#4803](https://github.com/hashicorp/terraform-provider-google-beta/pull/4803))
* bigquery: added `member` field to `google_bigquery_default_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* cloudbuild: added `script` field to `google_cloudbuild_trigger` resource ([#4807](https://github.com/hashicorp/terraform-provider-google-beta/pull/4807))
* cloudplatform: validated `project_id` for `google_project` data-source ([#4810](https://github.com/hashicorp/terraform-provider-google-beta/pull/4810))
* cloudrun: added field `liveness_probe` to resource `google_cloud_run_service` ([#4788](https://github.com/hashicorp/terraform-provider-google-beta/pull/4788))
* cloudrun: added field `startup_probe` to resource `google_cloud_run_service` ([#4773](https://github.com/hashicorp/terraform-provider-google-beta/pull/4773))
* compute: added `source_disk` field to `google_compute_disk` and `google_compute_region_disk` resource ([#4783](https://github.com/hashicorp/terraform-provider-google-beta/pull/4783))
* compute: added general field `rules` to `google_compute_router_nat` ([#4797](https://github.com/hashicorp/terraform-provider-google-beta/pull/4797))
* container: added `disk_size` and `disk_type` fields to `google_container_cluster.cluster_autoscaling.auto_provisioning_defaults` ([#4786](https://github.com/hashicorp/terraform-provider-google-beta/pull/4786))
* container: added support for in-place update of `node_config.0.tags` for `google_container_node_pool` resource ([#4781](https://github.com/hashicorp/terraform-provider-google-beta/pull/4781))
* datastream: added `private_connectivity` field to `google_datastream_connection_profile` ([#4808](https://github.com/hashicorp/terraform-provider-google-beta/pull/4808))
* dns: added `enable_geo_fencing` to `routing_policy` block of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* dns: added `health_checked_targets` to `wrr` and `geo` blocks of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* dns: added `primary_backup` to `routing_policy` block of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* firebase: added `app_urls` field to `google_firebase_web_app` ([#4798](https://github.com/hashicorp/terraform-provider-google-beta/pull/4798))
* firebase: added deletion support and new field `deletion_policy` for `google_firebase_web_app` ([#4796](https://github.com/hashicorp/terraform-provider-google-beta/pull/4796))
* privateca: added a new field `skip_grace_period` to skip the grace period when deleting a CertificateAuthority. ([#4784](https://github.com/hashicorp/terraform-provider-google-beta/pull/4784))
* serviceaccount: added `member` field to `google_service_account` resource and datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* sql: added `time_zone` field in `google_sql_database_instance` ([#4774](https://github.com/hashicorp/terraform-provider-google-beta/pull/4774))
* storage: added `member` field to `google_storage_project_service_account` and `google_storage_transfer_project_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))

BUG FIXES:
* compute: made `vm_count` in `google_compute_resource_policy` optional ([#4792](https://github.com/hashicorp/terraform-provider-google-beta/pull/4792))
* container: fixed inability to update `datapath_provider` on `google_container_cluster` by making field changes trigger resource recreation ([#4824](https://github.com/hashicorp/terraform-provider-google-beta/pull/4824))
* pubsub: ensured topics are recreated when their schemas change. ([#4791](https://github.com/hashicorp/terraform-provider-google-beta/pull/4791))
* redis: updated `persistence_config.rdb_snapshot_period` to optional in the `google_redis_instance` resource. ([#4821](https://github.com/hashicorp/terraform-provider-google-beta/pull/4821))

## 4.41.0 (October 17, 2022)

KNOWN ISSUES:
* container: This release introduced a new field, `node_config.0.guest_accelerator.0.gpu_sharing_config`, to an https://www.terraform.io/language/attr-as-blocks field (`node_config.0.guest_accelerator`). As detailed on the linked page, this may cause issues for modules and/or formats other than HCL.

BREAKING CHANGES:
* sql: updated `google_sql_user.sql_server_user_details` to be read only. Any configuration attempting to set this field is invalid and will cause the provider to fail during plan time. ([#4764](https://github.com/hashicorp/terraform-provider-google-beta/pull/4764))

FEATURES:
* **New Resource:**  `google_cloud_ids_endpoint` ([#4765](https://github.com/hashicorp/terraform-provider-google-beta/pull/4765))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_binding` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_member` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_policy` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))

IMPROVEMENTS:
* appengine: added support for `service_account` field to `google_app_engine_standard_app_version` resource ([#4757](https://github.com/hashicorp/terraform-provider-google-beta/pull/4757))
* bigquery: added `avro_options` field to `google_bigquery_table` resource ([#4768](https://github.com/hashicorp/terraform-provider-google-beta/pull/4768))
* cloudrun: added field `startup_probe` to resource `google_cloud_run_service` ([#4773](https://github.com/hashicorp/terraform-provider-google-beta/pull/4773))
* compute: added `node_config.0.guest_accelerator.0.gpu_sharing_config` field to `google_container_node_pool` resource ([#4758](https://github.com/hashicorp/terraform-provider-google-beta/pull/4758))
* datafusion: added `crypto_key_config` field to `google_data_fusion_instance` resource ([#4761](https://github.com/hashicorp/terraform-provider-google-beta/pull/4761))
* filestore: removed constraint that forced multiple `google_filestore_instance` creations to occur serially ([#4770](https://github.com/hashicorp/terraform-provider-google-beta/pull/4770))

BUG FIXES:
* kms: fixed apply failure when `google_kms_crypto_key` is removed after its versions were destroyed earlier ([#4769](https://github.com/hashicorp/terraform-provider-google-beta/pull/4769))
* monitoring: fixed a bug causing a perma-diff in `google_monitoring_alert_policy` when `cross_series_reducer` was set to "REDUCE_NONE" ([#4763](https://github.com/hashicorp/terraform-provider-google-beta/pull/4763))

## 4.40.0 (October 10, 2022)

FEATURES:
* **New Data Source:** `google_cloudfunctions2_function` ([#4732](https://github.com/hashicorp/terraform-provider-google-beta/pull/4732))
* **New Data Source:** `google_compute_snapshot` ([#4731](https://github.com/hashicorp/terraform-provider-google-beta/pull/4731))
* **New Resource:** `google_compute_region_target_tcp_proxy` ([#4749](https://github.com/hashicorp/terraform-provider-google-beta/pull/4749))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))
* **New Resource:** `google_bigquery_datapolicy_data_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_binding` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_member` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_org_policy_custom_constraint` ([#4741](https://github.com/hashicorp/terraform-provider-google-beta/pull/4741))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_feature` ([#4736](https://github.com/hashicorp/terraform-provider-google-beta/pull/4736))

IMPROVEMENTS:
* bigqueryreservation: added `concurrency` and `multiRegionAuxiliary` to `google_bigquery_reservation` ([#4739](https://github.com/hashicorp/terraform-provider-google-beta/pull/4739))
* bigtable: added additional retry GC policy operations with a longer poll interval to avoid quota issues ([#4750](https://github.com/hashicorp/terraform-provider-google-beta/pull/4750))
* bigtable: improved error messaging ([#4746](https://github.com/hashicorp/terraform-provider-google-beta/pull/4746))
* compute: added support for `compression_mode` field in `google_compute_backend_bucket` and `google_compute_backend_service` resource ([#4733](https://github.com/hashicorp/terraform-provider-google-beta/pull/4733))
* dataflow : added support of `labels` to resource `google_dataflow_flextemplate_job` ([#4748](https://github.com/hashicorp/terraform-provider-google-beta/pull/4748))
* datastream: added field `bigquery_profile` to `google_datastream_connection_profile` ([#4742](https://github.com/hashicorp/terraform-provider-google-beta/pull/4742))
* dns: added general field `cloud_logging_config` to `google_dns_managed_zone` ([#4734](https://github.com/hashicorp/terraform-provider-google-beta/pull/4734))
* metastore: added bigquery support for `google_dataproc_metastore_service` ([#4753](https://github.com/hashicorp/terraform-provider-google-beta/pull/4753))
* storage: added `custom_placement_config` field to `google_storage_bucket` resource to support custom dual-region GCS buckets ([#4752](https://github.com/hashicorp/terraform-provider-google-beta/pull/4752))
* sql: added  `password_policy` field to `google_sql_user` resource ([#4730](https://github.com/hashicorp/terraform-provider-google-beta/pull/4730))

BUG FIXES:
* storage: fixed a bug where user specified labels get overwritten by Dataplex auto generated labels ([#4743](https://github.com/hashicorp/terraform-provider-google-beta/pull/4743))
* storagetransfer: fixed a crash in `google_storagetransfer_job` refreshes when `transfer_schedule` was empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))


## 4.39.0 (October 3, 2022)

FEATURES:
* **New Data Source:** `google_artifact_registry_repository` ([#4714](https://github.com/hashicorp/terraform-provider-google-beta/pull/4714))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))

IMPROVEMENTS:
* certificatemanager: added public/private PEM fields `pem_certificate` / `pem_private_key` and deprecated `certificate_pem` / `private_key_pem` ([#4728](https://github.com/hashicorp/terraform-provider-google-beta/pull/4728))
* clouddeploy: added `serial_pipeline.stages.strategy` field to `google_clouddeploy_delivery_pipeline` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* container: added `notification_config.pubsub.filter` field to `google_container_cluster` ([#4718](https://github.com/hashicorp/terraform-provider-google-beta/pull/4718))
* eventarc: added `channels` and `conditions` fields to `google_eventarc_trigger` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* healthcare: added `notification_configs ` field to `google_healthcare_fhir_store` resource ([#4720](https://github.com/hashicorp/terraform-provider-google-beta/pull/4720))
* iap: added ability to import `google_iap_brand` using ID using {{project}}/{{brand_id}} format ([#4712](https://github.com/hashicorp/terraform-provider-google-beta/pull/4712))
* secretmanager: added output field 'version' to resource 'secret_manager_secret_version' ([#4724](https://github.com/hashicorp/terraform-provider-google-beta/pull/4724))
* sql: added `maintenance_version` and `available_maintenance_versions` fields to `google_sql_database_instance` resource ([#4725](https://github.com/hashicorp/terraform-provider-google-beta/pull/4725))
* storagetransfer: added `notification_config` field to `google_storage_transfer_job` resource ([#4709](https://github.com/hashicorp/terraform-provider-google-beta/pull/4709))
* tags: added `purpose` and `purpose_data` properties to `google_tags_tag_key` ([#4721](https://github.com/hashicorp/terraform-provider-google-beta/pull/4721))

BUG FIXES:
* bigquery: fixed a bug where `allow_quoted_newlines` and `allow_jagged_rows` could not be set to false on `google_bigquery_table` ([#4711](https://github.com/hashicorp/terraform-provider-google-beta/pull/4711))
* cloudfunction: fixed inability to update `docker_repository` and `kms_key_name` on `google_cloudfunctions_function` ([#4727](https://github.com/hashicorp/terraform-provider-google-beta/pull/4727))
* compute: fixed inability to manage Cloud Armor `adaptive_protection_config` on `google_compute_security_policy` ([#4726](https://github.com/hashicorp/terraform-provider-google-beta/pull/4726))
* container: fixed a bug where upgrading provider version breaks on `node_pool_auto_config` or `node_pool_defaults` ([#4706](https://github.com/hashicorp/terraform-provider-google-beta/pull/4706))
* iam: fixed diffs between `policy_data` from `google_iam_policy` data source and policy data in API responses ([#4722](https://github.com/hashicorp/terraform-provider-google-beta/pull/4722))
* iam: fixed permadiff resulting from empty fields being sent in requests to set conditional IAM policies ([#4723](https://github.com/hashicorp/terraform-provider-google-beta/pull/4723))
* secretmanager: fixed a bug where `google_secret_manager_secret_version` that was destroyed outside of Terraform would not be recreated on apply ([#4719](https://github.com/hashicorp/terraform-provider-google-beta/pull/4719))
* storagetransfer: fixed a crash in `google_storagetransfer_job` when `transfer_schedule` is empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))
 
## 4.38.0 (September 26, 2022)

FEATURES:
* **New Data Source:** `google_vpc_access_connector` ([#4693](https://github.com/hashicorp/terraform-provider-google-beta/pull/4693))
* **New Resource:** `google_datastream_private_connection` ([#4691](https://github.com/hashicorp/terraform-provider-google-beta/pull/4691))

IMPROVEMENTS:
* appengine: Added `egress_setting` for field `vpc_access_connector` to `google_app_engine_standard_app_version` ([#4701](https://github.com/hashicorp/terraform-provider-google-beta/pull/4701))
* bigquery: added `json_extension` field to the `load` block of `google_bigquery_job` resource ([#4699](https://github.com/hashicorp/terraform-provider-google-beta/pull/4699))
* cloudfunctions: Added `build_worker_pool` to `google_cloudfunctions_function` ([#4696](https://github.com/hashicorp/terraform-provider-google-beta/pull/4696))
* compute: added `json_custom_config` field to `google_compute_security_policy` resource ([#4703](https://github.com/hashicorp/terraform-provider-google-beta/pull/4703))
* redis: Added `persistence_config` field to the `google_redis_instance` resource. ([#4688](https://github.com/hashicorp/terraform-provider-google-beta/pull/4688))
* storage: added support for `overwriteWhen` field to `transfer_options` in `google_storage_transfer_job` resource ([#4690](https://github.com/hashicorp/terraform-provider-google-beta/pull/4690))

BUG FIXES:
* bigtable: added drift detection on `gc_rules` for `google_bigtable_gc_policy` ([#4687](https://github.com/hashicorp/terraform-provider-google-beta/pull/4687))
* compute: fixed the inability to update `most_disruptive_allowed_action` for both `google_compute_per_instance_config` and `google_compute_region_per_instance_config` ([#4685](https://github.com/hashicorp/terraform-provider-google-beta/pull/4685))
* container: fixed allow passing empty list to monitoring_config and logging_config in `google_container_cluster` ([#4700](https://github.com/hashicorp/terraform-provider-google-beta/pull/4700))
* sql: fixed a bug causing a perma-diff on `disk_type` due to API values being downcased ([#4686](https://github.com/hashicorp/terraform-provider-google-beta/pull/4686))
* storage: fixed the inability to set 0 for `lifecycle_rule.condition.age` on `google_storage_bucket` ([#4698](https://github.com/hashicorp/terraform-provider-google-beta/pull/4698))

## 4.37.0 (September 19, 2022)

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#4676](https://github.com/hashicorp/terraform-provider-google-beta/pull/4676))
* **New Resource:** `google_dialogflow_cx_webhook` ([#4667](https://github.com/hashicorp/terraform-provider-google-beta/pull/4667))
* **New Resource:** `google_filestore_snapshot` ([#4661](https://github.com/hashicorp/terraform-provider-google-beta/pull/4661))

IMPROVEMENTS:
* apigee: added read-only field `connection_state` to `google_apigee_endpoint_attachment` ([#4668](https://github.com/hashicorp/terraform-provider-google-beta/pull/4668))
* bigtable: added support for `autoscaling_config.storage_target` to `google_bigtable_instance` ([#4671](https://github.com/hashicorp/terraform-provider-google-beta/pull/4671))
* cloudbuild: added support for `BITBUCKET` option to `git_source.repo_type` in `google_cloudbuild_trigger` ([#4679](https://github.com/hashicorp/terraform-provider-google-beta/pull/4679))
* dns: added in validation for trailing dot at end of DNS record name ([#4674](https://github.com/hashicorp/terraform-provider-google-beta/pull/4674))
* project: added validation for field `project_id` in `google_project` datasource. ([#4684](https://github.com/hashicorp/terraform-provider-google-beta/pull/4684))
* serviceaccount: added `expires_in` attribute for generating `exp` claim to `google_service_account_jwt` datasource. ([#4677](https://github.com/hashicorp/terraform-provider-google-beta/pull/4677))

BUG FIXES:
* notebooks: fixed perma-diff in `google_notebooks_instance` ([#4664](https://github.com/hashicorp/terraform-provider-google-beta/pull/4664))
* privateca: fixed an issue that blocked subordinate CA data sources when `state` was not `AWAITING_USER_ACTIVATION` ([#4672](https://github.com/hashicorp/terraform-provider-google-beta/pull/4672))
* storage: fixed permdiff on the field `versioning` of `google_storage_bucket` ([#4665](https://github.com/hashicorp/terraform-provider-google-beta/pull/4665))

## 4.36.0 (September 12, 2022)

FEATURES:
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_binding` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_member` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_policy` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_datastream_connection_profile` ([#4657](https://github.com/hashicorp/terraform-provider-google-beta/pull/4657))

IMPROVEMENTS:
* appengine: added field `service_account` to `google_app_engine_flexible_app_version` ([#4653](https://github.com/hashicorp/terraform-provider-google-beta/pull/4653))
* bigtable: increased timeout in `google_bigtable_table` creation. ([#4655](https://github.com/hashicorp/terraform-provider-google-beta/pull/4655))
* cloudbuild: added `location` field to `google_cloudbuild_trigger` resource ([#4646](https://github.com/hashicorp/terraform-provider-google-beta/pull/4646))
* compute: added `certificate_map` to `compute_target_ssl_proxy` resource ([#4654](https://github.com/hashicorp/terraform-provider-google-beta/pull/4654))
* compute: added field `chain_name` to `google_compute_resource_policy.snapshot_properties` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* compute: added field `chain_name` to resource `google_compute_snapshot` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_cluster.node_pool` ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_node_pool` resource ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added field `node_pool_defaults` to `resource_container_cluster`. ([#4648](https://github.com/hashicorp/terraform-provider-google-beta/pull/4648))
* dataproc: added option `shielded_instance_config` to resource `google_dataproc_workflow_template`. ([#4647](https://github.com/hashicorp/terraform-provider-google-beta/pull/4647))
* metastore: extended default timeouts for `google_dataproc_metastore_service` from 40m to 60m ([#4652](https://github.com/hashicorp/terraform-provider-google-beta/pull/4652))
* pubsub: made `google_pubsub_subscription.enable_exactly_once_delivery` mutable so that it updates subscription without recreation. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))

## 4.35.0 (September 6, 2022)

IMPROVEMENTS:
* apigee: added support for `nodeConfig` in `google_apigee_environment` ([#4632](https://github.com/hashicorp/terraform-provider-google-beta/pull/4632))
* apigee: added a `properties` field to `google_apigee_organization` ([#4644](https://github.com/hashicorp/terraform-provider-google-beta/pull/4644))
* cloudfunctions2: added `secret_environment_variables` and `secret_volumes` to `google_cloudfunctions2_function` ([#4641](https://github.com/hashicorp/terraform-provider-google-beta/pull/4641))
* compute: added support for param `visible_core_count` in `google_compute_instance` and `google_compute_instance_template` under `advanced_machine_features` ([#4635](https://github.com/hashicorp/terraform-provider-google-beta/pull/4635))
* compute: added support documentation links to error messages for certain Compute Operation errors. ([#4642](https://github.com/hashicorp/terraform-provider-google-beta/pull/4642))
* container: added `service_external_ips_config` support to `cluster_container` resource. ([#4639](https://github.com/hashicorp/terraform-provider-google-beta/pull/4639))
* container: added `enable_cost_allocation` to `google_container_cluster` ([#4640](https://github.com/hashicorp/terraform-provider-google-beta/pull/4640))
* dns: added `behavior` field to `google_dns_response_policy_rule` resource ([#4637](https://github.com/hashicorp/terraform-provider-google-beta/pull/4637))
* monitoring: added `force_delete` field to `google_monitoring_notification_channel` resource ([#4638](https://github.com/hashicorp/terraform-provider-google-beta/pull/4638))
* pubsub: made `enable_exactly_once_delivery` mutable so that it updates subscription in-place and avoids recreation of the subscription. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))
* vertexai: added `encryption_spec` field to `google_vertex_ai_featurestore` resource (beta) ([#4643](https://github.com/hashicorp/terraform-provider-google-beta/pull/4643))

BUG FIXES:
* compute: fixed the `id` format of the data source `google_compute_instance` ([#4636](https://github.com/hashicorp/terraform-provider-google-beta/pull/4636))

## 4.34.0 (August 29, 2022)
NOTES:
* updated Bigtable go client version from 1.13 to 1.16. ([#4613](https://github.com/hashicorp/terraform-provider-google-beta/pull/4613))

IMPROVEMENTS:
* apigee: added support for specifying retention when deleting `google_apigee_organization` ([#4604](https://github.com/hashicorp/terraform-provider-google-beta/pull/4604))
* appengine: added `app_engine_apis` field to `google_app_engine_standard_app_version` resource ([#4607](https://github.com/hashicorp/terraform-provider-google-beta/pull/4607))
* compute: improved error messaging for compute errors ([#4602](https://github.com/hashicorp/terraform-provider-google-beta/pull/4602))
* container: added general field `reservation_affinity` to `google_container_node_pool` ([#4622](https://github.com/hashicorp/terraform-provider-google-beta/pull/4622))
* container: added field `auto_provisioning_network_tags` to `google_container_cluster` (beta) ([#4611](https://github.com/hashicorp/terraform-provider-google-beta/pull/4611))
* sql: added support for major version upgrade to `google_sql_database_instance ` resource  ([#4606](https://github.com/hashicorp/terraform-provider-google-beta/pull/4606))

BUG FIXES:
* bigtable: fixed comparing column family name when reading a GC policy. ([#4624](https://github.com/hashicorp/terraform-provider-google-beta/pull/4624))
* bigtable: passed `isTopeLevel` in getGCPolicyFromJSON() instead of hardcoding it to true. ([#4615](https://github.com/hashicorp/terraform-provider-google-beta/pull/4615))
* cloud iam: made `denial_condition` optional on `google_iam_deny_policy` ([#4617](https://github.com/hashicorp/terraform-provider-google-beta/pull/4617))

## 4.33.0 (August 22, 2022)
IMPROVEMENTS:
* container: added update support for `authenticator_groups_config` in `google_container_cluster` ([#4591](https://github.com/hashicorp/terraform-provider-google-beta/pull/4591))
* dataflow: added ability to import `google_dataflow_job` ([#4595](https://github.com/hashicorp/terraform-provider-google-beta/pull/4595))
* dns: added `managed_zone_id` attribute to `google_dns_managed_zone` data source ([#4593](https://github.com/hashicorp/terraform-provider-google-beta/pull/4593))
* metastore: added `metadata_integration` and `hive_metastore_config.auxiliary_versions` fields to `google_dataproc_metastore_service` resource ([#4598](https://github.com/hashicorp/terraform-provider-google-beta/pull/4598))
* monitoring: added `accepted_response_status_codes` to `google_monitoring_uptime_check_config` ([#4594](https://github.com/hashicorp/terraform-provider-google-beta/pull/4594))
* sql: added `password_validation_policy` field to `google_cloud_sql` resource ([#4597](https://github.com/hashicorp/terraform-provider-google-beta/pull/4597))

BUG FIXES:
* bigquery: removed force replacement for `display_name` on `google_bigquery_data_transfer_config` ([#4592](https://github.com/hashicorp/terraform-provider-google-beta/pull/4592))
* compute: fixed permadiff for `instance_termination_action` in `google_compute_instance_template` ([#4590](https://github.com/hashicorp/terraform-provider-google-beta/pull/4590))

## 4.32.0 (August 15, 2022)

NOTES:
* Updated to Golang 1.18 ([#4564](https://github.com/hashicorp/terraform-provider-google-beta/pull/4564))

FEATURES:
* **New Resource:** `google_dataplex_asset` ([#4543](https://github.com/hashicorp/terraform-provider-google-beta/pull/4543))
* **New Resource:** `google_gke_hub_membership_iam_binding` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_member` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_policy` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))

IMPROVEMENTS:
* certificatemanager: added `state`, `authorization_attempt_info` and `provisioning_issue` output fields to `google_certificate_manager_certificate` ([#4548](https://github.com/hashicorp/terraform-provider-google-beta/pull/4548))
* cloudfunctions2: added field `event_filters` to resource `google_cloudfunctions2_function` ([#4547](https://github.com/hashicorp/terraform-provider-google-beta/pull/4547))
* compute: added `certificate_map` to `compute_target_https_proxy` resource ([#4550](https://github.com/hashicorp/terraform-provider-google-beta/pull/4550))
* compute: added validation for name field on `google_compute_network` ([#4579](https://github.com/hashicorp/terraform-provider-google-beta/pull/4579))
* compute: made `port` optional in `google_compute_network_endpoint` to allow network endpoints to be associated with `GCE_VM_IP` network endpoint groups ([#4575](https://github.com/hashicorp/terraform-provider-google-beta/pull/4575))
* container: added support for additional values `APISERVER`, `CONTROLLER_MANAGER`, and `SCHEDULER` in `google_container_cluster.monitoring_config` ([#4565](https://github.com/hashicorp/terraform-provider-google-beta/pull/4565))
* gkehub: added `monitoring` and `mutation_enabled` fields to resource `feature_membership` ([#4572](https://github.com/hashicorp/terraform-provider-google-beta/pull/4572))
* gkehub: added better support for import for `google_gke_hub_membership` ([#4542](https://github.com/hashicorp/terraform-provider-google-beta/pull/4542))
* pubsub: added `bigquery_config` to `google_pubsub_subscription` ([#4545](https://github.com/hashicorp/terraform-provider-google-beta/pull/4545))
* scheduler: added `paused` field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))
* scheduler: added `state` output field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))

BUG FIXES:
* apigee: fixed an issue where `google_apigee_instance` creation would fail due to multiple concurrent instances ([#4584](https://github.com/hashicorp/terraform-provider-google-beta/pull/4584))
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.services` was not updating. ([#4577](https://github.com/hashicorp/terraform-provider-google-beta/pull/4577))
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dataflow: fixed bug where permadiff would show on `google_dataflow_job.additional_experiments` ([#4576](https://github.com/hashicorp/terraform-provider-google-beta/pull/4576))
* storage: fixed a bug in `google_storage_bucket` where `name` was incorrectly validated. ([#4566](https://github.com/hashicorp/terraform-provider-google-beta/pull/4566))

## 4.31.0 (Aug 1, 2022)

FEATURES:
* **New Resource:** `google_dataplex_zone` ([#4511](https://github.com/hashicorp/terraform-provider-google-beta/pull/4511))

IMPROVEMENTS:
* bucket: added support for `matches_prefix` and `matches_suffix` in `condition` of a `lifecycle_rule` in  `google_storage_bucket` ([#4527](https://github.com/hashicorp/terraform-provider-google-beta/pull/4527))
* compute: added `network` and `subnetwork` fields to `google_compute_region_network_endpoint_group` for PSC. ([#4528](https://github.com/hashicorp/terraform-provider-google-beta/pull/4528))
* container: added field `boot_disk_kms_key` to `auto_provisioning_defaults` in `google_container_cluster` ([#4524](https://github.com/hashicorp/terraform-provider-google-beta/pull/4524))
* notebooks: added `bootDiskType` support for `PD_EXTREME` in `google_notebooks_instance` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* notebooks: added `softwareConfig.upgradeable`, `softwareConfig.postStartupScriptBehavior`, `softwareConfig.kernels` in `google_notebooks_runtime` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* storage: added name validation for `google_storage_bucket` ([#4532](https://github.com/hashicorp/terraform-provider-google-beta/pull/4532))

BUG FIXES:
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dns: fixed a bug where `google_dns_record_set` would create an inconsistent plan when using interpolated values in `rrdatas` ([#4515](https://github.com/hashicorp/terraform-provider-google-beta/pull/4515))
* kms: fixed setting of resource id post-import for `google_kms_crypto_key` ([#4520](https://github.com/hashicorp/terraform-provider-google-beta/pull/4520))
* provider: fixed a bug where user-agent was showing "dev" rather than the provider version ([#4509](https://github.com/hashicorp/terraform-provider-google-beta/pull/4509))

## 4.30.0 (July 25, 2022)

FEATURES:
* **New Data Source:** `google_service_account_jwt` ([#4489](https://github.com/hashicorp/terraform-provider-google-beta/pull/4489))
* **New Resource:** `google_certificate_map_entry` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_certificate_map` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_compute_backend_bucket_iam_binding` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_member` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_policy` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_dataproc_metastore_federation` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_binding` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_member` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_policy` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))

IMPROVEMENTS:
* billingbudget: made `thresholdRules` optional in `google_billing_budget` ([#4480](https://github.com/hashicorp/terraform-provider-google-beta/pull/4480))
* compute: added `instance_termination_action` field to `google_compute_instance_template` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `instance_termination_action` field to `google_compute_instance` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `request_coalescing` and `bypass_cache_on_request_headers` fields to `compute_backend_bucket` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* compute: added field `all_instances_config` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4506](https://github.com/hashicorp/terraform-provider-google-beta/pull/4506))
* compute: added support for `esp` protocol in `google_compute_packet_mirroring.filters.ip_protocols` ([#4496](https://github.com/hashicorp/terraform-provider-google-beta/pull/4496))
* monitoring: added `evaluation_missing_data` field to `google_monitoring_alert_policy` ([#4502](https://github.com/hashicorp/terraform-provider-google-beta/pull/4502))
* notebooks: added field `reserved_ip_range` to `google_notebooks_runtime` ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))

BUG FIXES:
* bigtable: fixed an incorrect diff when adding two or more clusters ([#4490](https://github.com/hashicorp/terraform-provider-google-beta/pull/4490))
* compute: allowed properly updating `adaptive_protection_config` in `compute_security_policy` ([#4478](https://github.com/hashicorp/terraform-provider-google-beta/pull/4478))
* notebooks: fixed a bug where`google_notebooks_runtime` can't be updated ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))
* sql: fixed an issue in `google_sql_database_instance` where updates would fail because of the `collation` field ([#4505](https://github.com/hashicorp/terraform-provider-google-beta/pull/4505))

## 4.29.0 (July 18, 2022)

FEATURES:
* **New Resource:** `google_cloudiot_registry_iam_binding` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_member` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_policy` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_compute_snapshot_iam_binding` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_member` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_policy` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))

IMPROVEMENTS:
* container: added `binauthz_evaluation_mode` field to `resource_container_cluster`. ([#4451](https://github.com/hashicorp/terraform-provider-google-beta/pull/4451))
* kms: added support for MAC value in `google_kms_crypto_key.purpose` ([#4458](https://github.com/hashicorp/terraform-provider-google-beta/pull/4458))
* metastore: added `databaseType`, `releaseChannel`, and `hiveMetastoreConfig.endpointProtocol` arguments ([#4443](https://github.com/hashicorp/terraform-provider-google-beta/pull/4443))

BUG FIXES:
* bigquery: fixed case-sensitivity for `user_by_email` and `group_by_email` on `google_bigquery_dataset_access` ([#4446](https://github.com/hashicorp/terraform-provider-google-beta/pull/4446))
* clouddeploy: fixed permadiff on `execution_configs` in `google_clouddeploy_target` resource ([#4450](https://github.com/hashicorp/terraform-provider-google/pull/4450))
* cloudscheduler: fixed a diff on the last slash of uri on `google_cloud_scheduler_job` ([#4444](https://github.com/hashicorp/terraform-provider-google-beta/pull/4444))
* compute: fixed force recreation on `provisioned_iops` of `google_compute_disk` ([#4464](https://github.com/hashicorp/terraform-provider-google-beta/pull/4464))
* compute: fixed missing `network_interface.0.ipv6_access_config.0.external_ipv6` output on `google_compute_instance` ([#4470](https://github.com/hashicorp/terraform-provider-google-beta/pull/4470))
* documentai: fixed a bug where eu region could not be utilized for documentai resources ([#4472](https://github.com/hashicorp/terraform-provider-google-beta/pull/4472))
* gkehub: fixed a bug where `issuer` can't be updated on `google_gke_hub_membership` ([#4471](https://github.com/hashicorp/terraform-provider-google-beta/pull/4471))

## 4.28.0 (July 11, 2022)

FEATURES:
* **New Resource:** google_bigquery_connection_iam_binding ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_member ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_policy ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_cloud_tasks_queue_iam_binding ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_member ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_policy ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_dataproc_autoscaling_policy_iam_binding ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_member ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_policy ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_metastore_service_iam_binding ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_member ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_policy ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))

IMPROVEMENTS:
* bigquery: fixed a permadiff in `google_bigquery_job.query. destination_table` ([#4401](https://github.com/hashicorp/terraform-provider-google-beta/pull/4401))
* billing: added `calendar_period` and `custom_period` fields to `google_billing_budget` ([#4429](https://github.com/hashicorp/terraform-provider-google-beta/pull/4429))
* cloudsql: added attribute `project` to data source `google_sql_backup_run` ([#4402](https://github.com/hashicorp/terraform-provider-google-beta/pull/4402))
* composer: added CMEK, PUPI and IP_masq_agent support for Composer 2 in `google_composer_environment` resource ([#4430](https://github.com/hashicorp/terraform-provider-google-beta/pull/4430))
* compute: added `max_ports_per_vm` field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `GCE_VM_IP` support to `google_compute_network_endpoint_group` resource. ([#4434](https://github.com/hashicorp/terraform-provider-google-beta/pull/4434))
* privateca: added support to subordinate CA activation ([#4422](https://github.com/hashicorp/terraform-provider-google-beta/pull/4422))
* redis: added CMEK key field `customer_managed_key` in `google_redis_instance ` ([#4435](https://github.com/hashicorp/terraform-provider-google-beta/pull/4435))
* spanner: added field `version_retention_period` to `google_spanner_database` resource ([#4424](https://github.com/hashicorp/terraform-provider-google-beta/pull/4424))
* sql: added `settings.location_preference.secondary_zone` field in `google_sql_database_instance` ([#4433](https://github.com/hashicorp/terraform-provider-google-beta/pull/4433))
* sql: added `sql_server_audit_config` field in `google_sql_database_instance` ([#4403](https://github.com/hashicorp/terraform-provider-google-beta/pull/4403))

BUG FIXES:
* composer: fixed a problem with updating Cloud Composer's `scheduler_count` field (https://github.com/hashicorp/terraform-provider-google/issues/11940) ([#4408](https://github.com/hashicorp/terraform-provider-google-beta/pull/4408))
* composer: fixed permadiff on `private_environment_config.cloud_composer_connection_subnetwork` ([#4411](https://github.com/hashicorp/terraform-provider-google-beta/pull/4411))
* container: fixed an issue where `node_config.min_cpu_platform` could cause a perma-diff in `google_container_cluster` ([#4426](https://github.com/hashicorp/terraform-provider-google-beta/pull/4426))
* filestore: fixed a case where `google_filestore_instance.networks.network` would incorrectly see a diff between state and config when the network `id` format was used ([#4431](https://github.com/hashicorp/terraform-provider-google-beta/pull/4431))
* serviceusage: fixed an issue where `google_project_service_identity` didn't handle service identities without emails correctly ([#4432](https://github.com/hashicorp/terraform-provider-google-beta/pull/4432))


## 4.27.0 (June 27, 2022)

IMPROVEMENTS:
* clouddeploy: added `suspend` field to `google_clouddeploy_delivery_pipeline` resource ([#4394](https://github.com/hashicorp/terraform-provider-google-beta/pull/4394))
* compute: added maxPortsPerVm field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `psc_connection_id` and `psc_connection_status` output fields to `google_compute_forwarding_rule` and `google_compute_global_forwarding_rule` resources ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))
* container: added `tpu_config` to `google_container_cluster` (beta only) ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))
* containeraws: made `config.instance_type` field updatable in `google_container_aws_node_pool` ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))

BUG FIXES:
* compute: fixed default handling for `enable_dynamic_port_allocation ` to be managed by the api ([#4391](https://github.com/hashicorp/terraform-provider-google-beta/pull/4391))
* vertexai: Fixed a bug where terraform crashes when `force_destroy` is set in `google_vertex_ai_featurestore` resource ([#4398](https://github.com/hashicorp/terraform-provider-google-beta/pull/4398))

## 4.26.0 (June 21, 2022)

FEATURES:
* **New Resource:** `google_cloudfunctions2_function_iam_binding` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_member` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_policy` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_compute_region_ssl_policy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* **New Resource:** `google_documentai_processor` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))
* **New Resource:** `google_documentai_processor_default_version` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))

IMPROVEMENTS:
* accesscontextmanager: Added `external_resources` to `egress_to` in `google_access_context_manager_service_perimeter` and `google_access_context_manager_service_perimeters` resource ([#4378](https://github.com/hashicorp/terraform-provider-google-beta/pull/4378))
* apigateway: Added `grpc_services` and `managed_service_configs` to `google_api_gateway_api_config` ([#4388](https://github.com/hashicorp/terraform-provider-google-beta/pull/4388))
* cloudbuild: Added `include_build_logs` to `google_cloudbuild_trigger` ([#4380](https://github.com/hashicorp/terraform-provider-google-beta/pull/4380))
* compute: Added `ssl_policy` field to `google_compute_region_target_https_proxy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* container: Added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: Added `tpu_config` to `google_container_cluster` ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))

BUG FIXES:
* dns: Fixed a bug where `google_dns_record_set` resource can not be changed from default routing to Geo routing policy. ([#4383](https://github.com/hashicorp/terraform-provider-google-beta/pull/4383))
* sql: Fixed a bug where `google_sql_database_instance` would fail if a replica was created, with an encryption key, in a different region than the master instance. ([#4379](https://github.com/hashicorp/terraform-provider-google-beta/pull/4379))

## 4.25.0 (June 15, 2022)

IMPROVEMENTS:
* bigquery: added `connection_id` to `external_data_configuration` for `google_bigquery_table` ([#4365](https://github.com/hashicorp/terraform-provider-google-beta/pull/4365))
* cloudfunctions2: added support for configuring `service_account_email` to `google_cloudfunctions2_function` resource ([#4367](https://github.com/hashicorp/terraform-provider-google-beta/pull/4367))
* compute: added `advanced_options_config` to `google_compute_security_policy` ([#4354](https://github.com/hashicorp/terraform-provider-google-beta/pull/4354))
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))
* compute: added `include_named_cookies` to `cdn_policy` on `compute_backend_service` resource ([#4358](https://github.com/hashicorp/terraform-provider-google-beta/pull/4358))
* compute: added internal IPv6 support on `google_compute_network` and `google_compute_subnetwork` ([#4368](https://github.com/hashicorp/terraform-provider-google-beta/pull/4368))
* container: added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: added `spot` field to `node_config` sub-resource ([#4350](https://github.com/hashicorp/terraform-provider-google-beta/pull/4350))
* gkehub: added `prevent_drift` field to `google_gke_hub_feature_membership` resource ([#4370](https://github.com/hashicorp/terraform-provider-google-beta/pull/4370))
* monitoring: added support for JSONPath content matchers to `google_monitoring_uptime_check_config` resource ([#4361](https://github.com/hashicorp/terraform-provider-google-beta/pull/4361))
* monitoring: added support for `user_labels` to `google_monitoring_slo` resource ([#4363](https://github.com/hashicorp/terraform-provider-google-beta/pull/4363))
* sql: added `sql_server_user_details` field to `google_sql_user` resource ([#4364](https://github.com/hashicorp/terraform-provider-google-beta/pull/4364))

BUG FIXES:
* certificatemanager: fixed bug where `DEFAULT` scope would permadiff and force replace the certificate. ([#4356](https://github.com/hashicorp/terraform-provider-google-beta/pull/4356))
* dns: fixed perma-diff for updated labels in `google_dns_managed_zone` ([#4372](https://github.com/hashicorp/terraform-provider-google-beta/pull/4372))
* storagetransfer: fixed perm diff on transfer_options for `google_storage_transfer_job` ([#4357](https://github.com/hashicorp/terraform-provider-google-beta/pull/4357))

## 4.24.0 (June 6, 2022)

IMPROVEMENTS:
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))

## 4.23.0 (June 1, 2022)

FEATURES:
* **New Data Source:** `google_tags_tag_key` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Data Source:** `google_tags_tag_value` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Resource:** `google_dataplex_lake` ([#4341](https://github.com/hashicorp/terraform-provider-google-beta/pull/4341))

IMPROVEMENTS:
* bigqueryconnection: updated connection types to support v1 ga ([#4323](https://github.com/hashicorp/terraform-provider-google-beta/pull/4323))
* cloudfunctions: added docker registry support for Cloud Functions ([#4324](https://github.com/hashicorp/terraform-provider-google-beta/pull/4324))
* memcache: added `maintenance_policy` and `maintenance_schedule` to `google_memcache_instance` ([#4338](https://github.com/hashicorp/terraform-provider-google-beta/pull/4338))
* service-directory: marked network field immutable in `google_service_directory_endpoint` ([#4334](https://github.com/hashicorp/terraform-provider-google-beta/pull/4334))

BUG FIXES:
* binaryauthorization: fixed permadiff in `google_binary_authorization_attestor` ([#4325](https://github.com/hashicorp/terraform-provider-google-beta/pull/4325))
* service: added re-polling for service account after creation, 404s sometimes due to [eventual consistency](https://cloud.google.com/iam/docs/overview#consistency) ([#4333](https://github.com/hashicorp/terraform-provider-google-beta/pull/4333))

## 4.22.0 (May 24, 2022)

NOTE: Due to technical difficulties encountered in the release process, the `4.22.0` release for `google-beta` occurred several hours after the corresponding `google` provider release.

FEATURES:
* **New Resource:** `google_certificate_manager_certificate` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_certificate_manager_dns_authorization` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_clouddeploy_delivery_pipeline` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))
* **New Resource:** `google_clouddeploy_target` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))

IMPROVEMENTS:
* bigquery: added connection of type cloud_resource for `google_bigquery_connection` ([#4312](https://github.com/hashicorp/terraform-provider-google-beta/pull/4312))
* cloudfunctions: added `https_trigger_security_level` to `google_cloudfunctions_function` ([#4295](https://github.com/hashicorp/terraform-provider-google-beta/pull/4295))
* cloudrun: added `traffic.tag` and `traffic.url` fields to `google_cloud_run_service` ([#4283](https://github.com/hashicorp/terraform-provider-google-beta/pull/4283))
* compute: added `enable_dynamic_port_allocation` to `google_compute_router_nat` ([#4316](https://github.com/hashicorp/terraform-provider-google-beta/pull/4316))
* compute: added field `update_policy.most_disruptive_allowed_action` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* compute: added support for NEG type `PRIVATE_SERVICE_CONNECT` in `NetworkEndpointGroup` ([#4303](https://github.com/hashicorp/terraform-provider-google-beta/pull/4303))
* compute: added support for `domain_names` attribute in `google_compute_service_attachment` ([#4313](https://github.com/hashicorp/terraform-provider-google-beta/pull/4313))
* compute: added value `REFRESH` to field update_policy.minimal_action` in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* container: added field `exclusion_options` to `google_container_cluster` ([#4291](https://github.com/hashicorp/terraform-provider-google-beta/pull/4291))
* monitoring: added `checker_type` field to `google_monitoring_uptime_check_config` resource ([#4302](https://github.com/hashicorp/terraform-provider-google-beta/pull/4302))
* privateca: add a new field `desired_state` to manage CertificateAuthority state. ([#4279](https://github.com/hashicorp/terraform-provider-google-beta/pull/4279))
* sql: added `active_directory_config` field in `google_sql_database_instance` ([#4298](https://github.com/hashicorp/terraform-provider-google-beta/pull/4298))
* sql: removed requirement that Cloud SQL Insight is only allowed for Postgres in `google_sql_database_instance` ([#4310](https://github.com/hashicorp/terraform-provider-google-beta/pull/4310))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed extra diffs generated on `google_security_policy` `rules` when modifying a rule ([#4287](https://github.com/hashicorp/terraform-provider-google-beta/pull/4287))
* container: fixed Autopilot cluster couldn't omit master ipv4 cidr in `google_container_cluster` ([#4280](https://github.com/hashicorp/terraform-provider-google-beta/pull/4280))
* resourcemanager: fixed a bug in wrongly writing to state when creation failed on `google_project_organization_policy` ([#4297](https://github.com/hashicorp/terraform-provider-google-beta/pull/4297))
* storage: not specifying `content` or `source` for `google_storage_bucket_object` now fails at plan-time instead of apply-time. ([#4292](https://github.com/hashicorp/terraform-provider-google-beta/pull/4292))

## 4.21.0 (May 16, 2022)

IMPROVEMENTS:
* cloudfunctions: added CMEK support for Cloud Functions ([#4272](https://github.com/hashicorp/terraform-provider-google-beta/pull/4272))
* compute: added `service_directory_registrations` to `google_compute_forwarding_rule` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* compute: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* container: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* containeraws: added `image_type` and `instance_placement` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `instance_placement` and `logging_config` to `google_container_aws_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `proxy_config` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `image_type` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `logging_config` to `google_container_azure_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `proxy_config` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* dataproc: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* dns: added `routing_policy` to `google_dns_record_set` resource ([#4265](https://github.com/hashicorp/terraform-provider-google-beta/pull/4265))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed a crash in `google_compute_instance` when the instance is deleted outside of Terraform ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))

## 4.20.0 (May 2, 2022)

NOTES:
* `google_privateca_certificate_authority` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))

FEATURES:
* **New Data Source:** `google_compute_disk` ([#4255](https://github.com/hashicorp/terraform-provider-google-beta/pull/4255))

IMPROVEMENTS:
* apigee: `consumer_accept_list` and `service_attachment` to `google_apigee_instance`. ([#4260](https://github.com/hashicorp/terraform-provider-google-beta/pull/4260))
* compute: added `subsetting` field to `google_compute_region_backend_service` ([#4246](https://github.com/hashicorp/terraform-provider-google-beta/pull/4246))
* privateca: added `deletion_protection` for `google_privateca_certificate_authority`. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))
* privateca: added new output fields on `google_privateca_certificate` including `issuer_certificate_authority`, `pem_certificate_chain` and `certificate_description.x509_description` ([#4242](https://github.com/hashicorp/terraform-provider-google-beta/pull/4242))
* redis: added multi read replica field `read_replicas_mode` and `secondary_ip_range` in `google_redis_instance` ([#4259](https://github.com/hashicorp/terraform-provider-google-beta/pull/4259))

BUG FIXES:
* compute: fixed a crash when `compute.instance` is not found ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))
* sql: fixed bug where `encryption_key_name` was not being propagated to the API. ([#4261](https://github.com/hashicorp/terraform-provider-google-beta/pull/4261))

## 4.19.0 (April 25, 2022)

IMPROVEMENTS:
* cloudbuild: made `CLOUD_LOGGING_ONLY` available as a cloud build logging option. ([#4224](https://github.com/hashicorp/terraform-provider-google-beta/pull/4224))
* compute: added `redirect_options` field for `google_compute_security_policy` rules ([#4217](https://github.com/hashicorp/terraform-provider-google-beta/pull/4217))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance_template` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added passing `exceed_redirect_options` field for `google_compute_security_policy` rules ([#4238](https://github.com/hashicorp/terraform-provider-google-beta/pull/4238))
* container: added `gke_backup_agent_config` in `addons_config` to `google_container_cluster` (beta) ([#4231](https://github.com/hashicorp/terraform-provider-google-beta/pull/4231))
* filestore: added `kms_key_name` field to `google_filestore_instance` resource to support CMEK ([#11493](https://github.com/hashicorp/terraform-provider-google/pull/11493))
* logging: made `google_logging_*_bucket_config` deletable ([#4234](https://github.com/hashicorp/terraform-provider-google-beta/pull/4234))
* notebooks: updated `container_images` on `google_notebooks_runtime` to default to the value returned by the API if not set ([#4216](https://github.com/hashicorp/terraform-provider-google-beta/pull/4216))
* provider: modified request retry logic to retry all per-minute quota limits returned with a 403 error code. Previously, only read requests were retried. This will generally affect Google Compute Engine resources. ([#4223](https://github.com/hashicorp/terraform-provider-google-beta/pull/4223))

BUG FIXES:
* bigquery: fixed a bug where `encryption_configuration.kms_key_name` stored the version rather than the key name. ([#4221](https://github.com/hashicorp/terraform-provider-google-beta/pull/4221))
* compute: fixed url_mask required mis-annotation in `google_compute_region_network_endpoint_group`, making it optional ([#4227](https://github.com/hashicorp/terraform-provider-google-beta/pull/4227))
* spanner: fixed escaping of database names with Postgres dialect in `google_spanner_database` ([#4228](https://github.com/hashicorp/terraform-provider-google-beta/pull/4228))

## 4.18.0 (April 18, 2022)

FEATURES:
* **New Resource:** `google_privateca_certificate_template_iam_binding` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_member` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_policy` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))

IMPROVEMENTS:
* bigtable: added `gc_rules` to `google_bigtable_gc_policy` resource. ([#4212](https://github.com/hashicorp/terraform-provider-google-beta/pull/4212))
* dialogflow: added support for location based dialogflow resources ([#4206](https://github.com/hashicorp/terraform-provider-google-beta/pull/4206))
* metastore: added support for encryption_config during service creation. ([#4204](https://github.com/hashicorp/terraform-provider-google-beta/pull/4204))
* privateca: support update on CertificateAuthority and Certificate ([#4207](https://github.com/hashicorp/terraform-provider-google-beta/pull/4207))

BUG FIXES:
* Update mutex on google_apigee_instance_attachment to lock on org_id. ([#4203](https://github.com/hashicorp/terraform-provider-google-beta/pull/4203))
* vpcaccess: fixed an issue where `google_vpc_access_connector` would be repeatedly recreated when `network` was not specified ([#4205](https://github.com/hashicorp/terraform-provider-google-beta/pull/4205))

## 4.17.0 (April 11, 2022)

FEATURES:
* **New Data Source:** `google_access_approval_folder_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_organization_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_project_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Resource:** `google_access_context_manager_access_policy_iam_binding` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_member` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_endpoints_service_consumers_iam_binding` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_member` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_policy` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_iam_deny_policy` ([#4194](https://github.com/hashicorp/terraform-provider-google-beta/pull/4194))

IMPROVEMENTS:
* access approval: added `active_key_version`, `ancestor_has_active_key_version`, and `invalid_key_version` fields to `google_folder_access_approval_settings`, `google_organization_access_approval_settings`, and `google_project_access_approval_settings` resources ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* access context manager: added support for scoped policies in `google_access_context_manager_access_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* apigee: added `deployment_type` and `api_proxy_type` to `google_apigee_environment` ([#4177](https://github.com/hashicorp/terraform-provider-google-beta/pull/4177))
* bigtable: updated the examples to show users can create all 3 different flavors of AppProfile ([#4172](https://github.com/hashicorp/terraform-provider-google-beta/pull/4172))
* cloudbuild: added `approval_config` to `google_cloudbuild_trigger` ([#4162](https://github.com/hashicorp/terraform-provider-google-beta/pull/4162))
* composer: added support for `airflow-1` and `airflow-2` aliases in image version argument ([#4185](https://github.com/hashicorp/terraform-provider-google-beta/pull/4185))
* dataflow: added `skip_wait_on_job_termination` attribute to `google_dataflow_job` and `google_dataflow_flex_template_job` resources (issue #10559) ([#4196](https://github.com/hashicorp/terraform-provider-google-beta/pull/4196))
* dataproc: added `presto_config` to `dataproc_job` ([#4171](https://github.com/hashicorp/terraform-provider-google-beta/pull/4171))
* healthcare: added support V3 parser version for Healthcare HL7 stores. ([#4189](https://github.com/hashicorp/terraform-provider-google-beta/pull/4189))
* healthcare: added support for `ANALYTICS_V2 `and `LOSSLESS` BigQueryDestination schema types to `google_healthcare_fhir_store` ([#4186](https://github.com/hashicorp/terraform-provider-google-beta/pull/4186))
* os-config: added field `migInstancesAllowed` to resource `os_config_patch_deployment` ([#4195](https://github.com/hashicorp/terraform-provider-google-beta/pull/4195))
* privateca: added support for IAM conditions to CaPool ([#4170](https://github.com/hashicorp/terraform-provider-google-beta/pull/4170))
* pubsub: added `enable_exactly_once_delivery` to `google_pubsub_subscription` ([#4166](https://github.com/hashicorp/terraform-provider-google-beta/pull/4166))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))

BUG FIXES:
* redis: fixed an issue where older redis instances had a dangerous diff on the field `read_replicas_mode`, adding a default of `READ_REPLICAS_DISABLED`. Now, if the field is not set in config, the value of the field will keep the old value from state. ([#4184](https://github.com/hashicorp/terraform-provider-google-beta/pull/4184))
* tags: fixed issue where tags could not be applied sequentially to the same parent in `google_tags_tag_binding` ([#4191](https://github.com/hashicorp/terraform-provider-google-beta/pull/4191))

## 4.16.0 (April 4, 2022)

FEATURES:
* **New Data Source:** `google_dataproc_metastore_service` ([#4155](https://github.com/hashicorp/terraform-provider-google-beta/pull/4155))
* **New Resource:** `google_firebaserules_release` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))
* **New Resource:** `google_firebaserules_ruleset` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))

IMPROVEMENTS:
* bigtable: added support for `autoscaling_config` to `google_bigtable_instance` ([#4150](https://github.com/hashicorp/terraform-provider-google-beta/pull/4150))
* composer: Added support for `composer-1` and `composer-2` aliases in image version argument ([#4131](https://github.com/hashicorp/terraform-provider-google-beta/pull/4131))
* compute: added support for attaching a `edge_security_policy` to `google_compute_backend_bucket` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* compute: added support for field `type` to `google_compute_security_policy` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* eventarc: added gke and workflows destination for eventarc trigger resource. ([#4152](https://github.com/hashicorp/terraform-provider-google-beta/pull/4152))
* networkservices: added `included_cookie_names` to cache key policy configuration ([#4147](https://github.com/hashicorp/terraform-provider-google-beta/pull/4147))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))
* storagetransfer: added `repeat_interval` field to `google_storage_transfer_job` resource ([#4144](https://github.com/hashicorp/terraform-provider-google-beta/pull/4144))

BUG FIXES:
* apikeys: fixed a bug where `google_apikeys_key.key_string` was not being set. ([#4139](https://github.com/hashicorp/terraform-provider-google-beta/pull/4139))
* container: fixed a bug where `google_container_cluster.authenticator_groups_config` could not be set in tandem with `enable_autopilot` ([#4140](https://github.com/hashicorp/terraform-provider-google-beta/pull/4140))
* iam: fixed an issue where special identifiers `allAuthenticatedUsers` and `allUsers` were flattened to lower case in IAM members. ([#4156](https://github.com/hashicorp/terraform-provider-google-beta/pull/4156))
* logging: fixed bug where `google_logging_project_bucket_config` would erroneously write to state after it errored out and wasn't actually created. ([#4141](https://github.com/hashicorp/terraform-provider-google-beta/pull/4141))
* monitoring: fixed a permadiff when `google_monitoring_uptime_check_config.http_check.path` does not begin with "/" ([#4135](https://github.com/hashicorp/terraform-provider-google-beta/pull/4135))
* osconfig: fixed a bug where `recurring_schedule.time_of_day` can not be set to 12am exact time in `google_os_config_patch_deployment` resource ([#4127](https://github.com/hashicorp/terraform-provider-google-beta/pull/4127))
* sql: fixed bug where permadiff of `encryption_key_name` would show on `google_sql_database_instance` for replica instances. ([#4130](https://github.com/hashicorp/terraform-provider-google-beta/pull/4130))
* storage: fixed a bug where `google_storage_bucket` data source would retry for 20 min when bucket was not found. ([#4129](https://github.com/hashicorp/terraform-provider-google-beta/pull/4129))
* storage: fixed bug where `google_storage_transfer_job` that was deleted outside of Terraform would not be recreated on apply. ([#4138](https://github.com/hashicorp/terraform-provider-google-beta/pull/4138))


## 4.15.0 (March 21, 2022)

FEATURES:
* **New Resource:** google_logging_log_view ([#4125](https://github.com/hashicorp/terraform-provider-google-beta/pull/4125))

IMPROVEMENTS:
* apigee: added `billing_type` attribute to `google_apigee_organization` resource. ([#4126](https://github.com/hashicorp/terraform-provider-google-beta/pull/4126))
* networkservices: added `disable_http2` property to `google_network_services_edge_cache_service` resource ([#4119](https://github.com/hashicorp/terraform-provider-google-beta/pull/4119))
* networkservices: updated `google_network_services_edge_cache_origin` resource to read and write the `timeout` property, including a new `read_timeout` field. ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))
* networkservices: updated `google_network_services_edge_cache_origin` to retry_conditions to include `FORBIDDEN` ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_workflow_template` ([#4124](https://github.com/hashicorp/terraform-provider-google-beta/pull/4124))
* sql: fixed crash when one of `settings.database_flags` is nil. ([#4123](https://github.com/hashicorp/terraform-provider-google-beta/pull/4123))

## 4.14.0 (March 14, 2022)

FEATURES:
* **New Resource:** `google_bigqueryreservation_assignment` ([#4098](https://github.com/hashicorp/terraform-provider-google-beta/pull/4098))
* **New Resource:** `google_apikeys_key` ([#4114](https://github.com/hashicorp/terraform-provider-google-beta/pull/4114))

IMPROVEMENTS:
* artifactregistry: added maven config for `google_artifact_registry_repository` ([#4112](https://github.com/hashicorp/terraform-provider-google-beta/pull/4112))
* cloudbuild: added support for manual builds, git source for webhook/pubsub triggered builds and filter field ([#4100](https://github.com/hashicorp/terraform-provider-google-beta/pull/4100))
* container: added support for gvnic to `google_container_node_pool` ([#4111](https://github.com/hashicorp/terraform-provider-google-beta/pull/4111))
* dataproc: added `preemptibility` field to the `preemptible_worker_config` of `google_dataproc_cluster` ([#4107](https://github.com/hashicorp/terraform-provider-google-beta/pull/4107))
* serviceusage: supported `force` behavior for deleting consumer quota override ([#4094](https://github.com/hashicorp/terraform-provider-google-beta/pull/4094))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_job` ([#4108](https://github.com/hashicorp/terraform-provider-google-beta/pull/4108))

## 4.13.0 (March 7, 2022)

FEATURES:
* **New Resource:** `google_apigee_endpoint_attachment` ([#4074](https://github.com/hashicorp/terraform-provider-google-beta/pull/4074))
* **New Resource:** `google_cloudfunctions2_function` ([#4093](https://github.com/hashicorp/terraform-provider-google-beta/pull/4093))
* **New Resource:** `google_region_backend_service_iam_*` ([#4088](https://github.com/hashicorp/terraform-provider-google-beta/pull/4088))
* **New Datasource:** `google_dns_record_set` ([#4085](https://github.com/hashicorp/terraform-provider-google-beta/pull/4085))
* **New Datasource:** `google_privateca_certificate_authority` ([#4087](https://github.com/hashicorp/terraform-provider-google-beta/pull/4087))

IMPROVEMENTS:
* compute: added support for `keepalive_interval` to `google_compute_router.bgp` ([#4089](https://github.com/hashicorp/terraform-provider-google-beta/pull/4089))
* compute: added update support for `google_compute_reservation.share_settings` ([#4092](https://github.com/hashicorp/terraform-provider-google-beta/pull/4092))
* storagetransfer: added attribute `subject_id` to data source `google_storage_transfer_project_service_account` ([#4073](https://github.com/hashicorp/terraform-provider-google-beta/pull/4073))

BUG FIXES:
* composer: allow region to be undefined in configuration for `google_composer_environment` ([#4083](https://github.com/hashicorp/terraform-provider-google-beta/pull/4083))
* container: fixed a bug where `vertical_pod_autoscaling` would cause autopilot clusters to recreate ([#4076](https://github.com/hashicorp/terraform-provider-google-beta/pull/4076))

## 4.12.0 (February 28, 2022)

NOTE:
* updated to go 1.16.14 ([#4066](https://github.com/hashicorp/terraform-provider-google-beta/pull/4066))

FEATURES:
* **New Resource:** `dns_response_policy*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))
* **New Resource:** `dns_response_policy_rule*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))

DEPRECATIONS:
* datafusion: deprecated `service_account` in `google_datafusion_instance`. Use `tenant_project_id` instead to extract the tenant project ID (beta) ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))

IMPROVEMENTS:
* bigquery: added support for authorized datasets to `google_bigquery_dataset.access` and `google_bigquery_dataset_access` ([#4047](https://github.com/hashicorp/terraform-provider-google-beta/pull/4047))
* bigtable: added `multi_cluster_routing_cluster_ids` fields to `google_bigtable_app_profile` ([#4051](https://github.com/hashicorp/terraform-provider-google-beta/pull/4051))
* compute: added field `serverless_deployment` to `google_compute_network_endpoint_group` (beta only) for API Gateway resources ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* compute: updated `instance` attribute for `google_compute_network_endpoint` to be optional, as Hybrid connectivity NEGs use network endpoints with just IP and Port. ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `NON_GCP_PRIVATE_IP_PORT` value for `network_endpoint_type` in the `google_compute_network_endpoint_group` resource ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `provisioning_model` field to `google_compute_instance_template ` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* compute: added `provisioning_model` field to `google_compute_instance` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* container: Add support for GKE Compact Placement ([#4043](https://github.com/hashicorp/terraform-provider-google-beta/pull/4043))
* datafusion: added support for `tenant_project_id` and `gcs_bucket` in `google_datafusion_instance` resource. ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))
* provider: added retries for `ReadRequest` errors incorrectly coded as `403` errors, particularly in Google Compute Engine ([#4064](https://github.com/hashicorp/terraform-provider-google-beta/pull/4064))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance` could not be used on the same `google_apigee_organization` ([#4059](https://github.com/hashicorp/terraform-provider-google-beta/pull/4059))
* compute: corrected an issue in `google_compute_security_policy` where only alpha values for certain enums were accepted ([#4049](https://github.com/hashicorp/terraform-provider-google-beta/pull/4049))
* compute: fixed permadiff in `google_compute_instance.scheduling.provisioning_model` ([#4044](https://github.com/hashicorp/terraform-provider-google-beta/pull/4044))
* compute: fixed permadiff in `google_compute_instance_template.scheduling.provisioning_model` ([#4052](https://github.com/hashicorp/terraform-provider-google-beta/pull/4052))

## 4.11.0 (February 16, 2022)

IMPROVEMENTS:
* cloudfunctions: Added SecretManager integration support to `google_cloudfunctions_function`. ([#4040](https://github.com/hashicorp/terraform-provider-google-beta/pull/4040))
* compute: Added field `serverless_deployment` to `google_compute_network_endpoint_group` ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* sql: added field `clone.allocated_ip_range` to support address range picker for clone in resource `google_sql_database_instance` ([#4037](https://github.com/hashicorp/terraform-provider-google-beta/pull/4037))
* storagetransfer: added support for POSIX data source and data sink to `google_storage_transfer_job` via `transfer_spec.posix_data_source` and `transfer_spec.posix_data_sink` fields ([#4029](https://github.com/hashicorp/terraform-provider-google-beta/pull/4029))

BUG FIXES:
* cloudrun: updated `containers.ports.container_port` to be optional instead of required on `google_cloud_run_service` ([#4030](https://github.com/hashicorp/terraform-provider-google-beta/pull/4030))
* compute: marked `project` field optional in `google_compute_instance_template` data source ([#4031](https://github.com/hashicorp/terraform-provider-google-beta/pull/4031))

## 4.10.0 (February 7, 2022)

FEATURES:
* **New Resource:** `google_backend_service_iam_*` ([#4021](https://github.com/hashicorp/terraform-provider-google-beta/pull/4021))

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_global_forwarding_rule` resource ([#4011](https://github.com/hashicorp/terraform-provider-google-beta/pull/4011))
* compute: added field `rate_limit_options` to `google_compute_security_policy` rules ([#4020](https://github.com/hashicorp/terraform-provider-google-beta/pull/4020))
* container: added support for image type configuration on the GKE Node Auto-provisioning ([#4023](https://github.com/hashicorp/terraform-provider-google-beta/pull/4023))
* container: added support for GCPFilestoreCSIDriver addon to `google_container_cluster` resource. ([#4015](https://github.com/hashicorp/terraform-provider-google-beta/pull/4015))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* vpcaccess: updated field `network` in `google_vpc_access_connector` to accept `self_link` or `name` ([#4013](https://github.com/hashicorp/terraform-provider-google-beta/pull/4013))

BUG FIXES:
* storage: fixed bug where the provider crashes when `Object.owner` is missing when using `google_storage_object_acl` ([#4019](https://github.com/hashicorp/terraform-provider-google-beta/pull/4019))

## 4.9.0 (January 31, 2022)

BREAKING CHANGES:
* cloudrun: changed the `location` of `google_cloud_run_service` so that modifying the `location` field will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))

IMPROVEMENTS:
* provider: changed the default timeout for many resources to 20 minutes, the current Terraform default, where it was less than 20 minutes previously ([#4002](https://github.com/hashicorp/terraform-provider-google-beta/pull/4002))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* storage: added field `transfer_spec.aws_s3_data_source.role_arn` to `google_storage_transfer_job` ([#3999](https://github.com/hashicorp/terraform-provider-google-beta/pull/3999))

BUG FIXES:
* cloudrun: fixed a bug where changing the non-updatable `location` of a `google_cloud_run_service` would not force resource recreation ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))
* compute: fixed a bug where `google_compute_firewall` would incorrectly find `source_ranges` to be empty during validation ([#4008](https://github.com/hashicorp/terraform-provider-google-beta/pull/4008))
* notebooks: fixed permadiff in `google_notebooks_runtime.software_config` ([#3997](https://github.com/hashicorp/terraform-provider-google-beta/pull/3997))

## 4.8.0 (January 24, 2022)

BREAKING CHANGES:
* dlp: renamed the `characters_to_ignore.character_to_skip` field to `characters_to_ignore.characters_to_skip` in `google_data_loss_prevention_deidentify_template`. Any affected configurations will have been failing with an error at apply time already. ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))

FEATURES:
* **New Resource:** `google_network_connectivity_spoke` ([#3987](https://github.com/hashicorp/terraform-provider-google-beta/pull/3987))

IMPROVEMENTS:
* apigee: added `ip_range` field to `google_apigee_instance` ([#3989](https://github.com/hashicorp/terraform-provider-google-beta/pull/3989))
* cloudrun: added support for `default_mode` and `mode` settings for created files within `secrets` in `google_cloud_run_service` ([#3984](https://github.com/hashicorp/terraform-provider-google-beta/pull/3984))
* compute: Added `share_settings` in `google_compute_reservation` ([#3980](https://github.com/hashicorp/terraform-provider-google-beta/pull/3980))

BUG FIXES:
* all: Fixed operation polling to support custom endpoints. ([#3986](https://github.com/hashicorp/terraform-provider-google-beta/pull/3986))
* cloudrun: Fixed permadiff in `google_cloud_run_service`'s `template.spec.service_account_name`. ([#3993](https://github.com/hashicorp/terraform-provider-google-beta/pull/3993))
* dlp: Fixed typo in name of `characters_to_ignore.characters_to_skip` field for `google_data_loss_prevention_deidentify_template` ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))
* storagetransfer: fixed bug where `schedule` was required, but really it is optional. ([#3995](https://github.com/hashicorp/terraform-provider-google-beta/pull/3995))

## 4.7.0 (January 19, 2022)

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_backend_service` resource ([#3975](https://github.com/hashicorp/terraform-provider-google-beta/pull/3975))
* container: promoted `dns_config` field of `google_container_cluster` to GA ([#3978](https://github.com/hashicorp/terraform-provider-google-beta/pull/3978))
* monitoring: added `conditionMatchedLog` and `alertStrategy` fields to `google_monitoring_alert_policy` resource ([#3968](https://github.com/hashicorp/terraform-provider-google-beta/pull/3968))

## 4.6.0 (January 10, 2022)

BREAKING CHANGES:
* pubsub: changed `google_pubsub_schema` so that modifiying fields will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#3941](https://github.com/hashicorp/terraform-provider-google-beta/pull/3941))
* **New Resource:** `google_network_connectivity_hub` ([#3947](https://github.com/hashicorp/terraform-provider-google-beta/pull/3947))

IMPROVEMENTS:
* bigquery: added ability to create a table with both a schema and view simultaneously to `google_bigquery_table` ([#3950](https://github.com/hashicorp/terraform-provider-google-beta/pull/3950))
* cloud_composer: Added support for Cloud Composer master authorized networks flag ([#3937](https://github.com/hashicorp/terraform-provider-google-beta/pull/3937))
* container: Added field `identity_service_config` to `google_container_cluster` ([#3957](https://github.com/hashicorp/terraform-provider-google-beta/pull/3957))
* osconfig: Added daily os config patch deployments ([#3945](https://github.com/hashicorp/terraform-provider-google-beta/pull/3945))
* storage: added configurable read timeout to `google_storage_bucket` ([#3938](https://github.com/hashicorp/terraform-provider-google-beta/pull/3938))

BUG FIXES:
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.labels` was not updating. ([#3932](https://github.com/hashicorp/terraform-provider-google-beta/pull/3932))
* compute: fixed scenario where `region_instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3949](https://github.com/hashicorp/terraform-provider-google-beta/pull/3949))
* healthcare: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))
* pubsub: fixed update failure when attempting to change non-updatable resource `google_pubsub_schema` ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))
* storage: fixed a bug where `google_storage_bucket.lifecycle_rule.condition.days_since_custom_time` was not updating. ([#3936](https://github.com/hashicorp/terraform-provider-google-beta/pull/3936))
* vpcaccess: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))

## 4.5.0 (December 20, 2021)

FEATURES:
* **New Data Source:** google_container_aws_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Data Source:** google_container_azure_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_client ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))


IMPROVEMENTS:
* bigquery: added the `return_table_type` field to `google_bigquery_routine` ([#3922](https://github.com/hashicorp/terraform-provider-google-beta/pull/3922))
* cloudbuild: added support for `available_secrets` to `google_cloudbuild_trigger` ([#3907](https://github.com/hashicorp/terraform-provider-google-beta/pull/3907))
* cloudfunctions: added support for `min_instances` to `google_cloudfunctions_function` ([#3904](https://github.com/hashicorp/terraform-provider-google-beta/pull/3904))
* composer: added support for Private Service Connect by adding field `cloud_composer_connection_subnetwork` in `google_composer_environment` ([#3912](https://github.com/hashicorp/terraform-provider-google-beta/pull/3912))
* compute: fixed bug where `google_compute_instance`'s `can_ip_forward` could not be updated without recreating or restarting the instance. ([#3920](https://github.com/hashicorp/terraform-provider-google-beta/pull/3920))
* compute: added field `public_access_prevention` to resource `bucket` ([#3919](https://github.com/hashicorp/terraform-provider-google-beta/pull/3919))
* compute: added support for regional external HTTP(S) load balancer ([#3916](https://github.com/hashicorp/terraform-provider-google-beta/pull/3916))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* provider: enabled gRPC requests and response logging ([#3910](https://github.com/hashicorp/terraform-provider-google-beta/pull/3910))

BUG FIXES:
* assuredworkloads: fixed a bug preventing `google_assured_workloads_workload` from being created in any region other than us-central1 ([#3925](https://github.com/hashicorp/terraform-provider-google-beta/pull/3925))

## 4.4.0 (December 13, 2021)

DEPRECATIONS:
* filestore: deprecated `zone` on `google_filestore_instance` in favor of `location` to allow for regional instances ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))

FEATURES:
* **New Resource:** `google_os_config_os_policy_assignment` ([#3892](https://github.com/hashicorp/terraform-provider-google-beta/pull/3892))
* **New Resource:** `google_recaptcha_enterprise_key` ([#3890](https://github.com/hashicorp/terraform-provider-google-beta/pull/3890))

IMPROVEMENTS:
* filestore: added support for `ENTERPRISE` value on `google_filestore_instance` `tier` ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* sql: added field `allocated_ip_range` to resource `google_sql_database_instance` ([#3897](https://github.com/hashicorp/terraform-provider-google-beta/pull/3897))

BUG FIXES:
* compute: fixed incorrectly failing validation for `INTERNAL_MANAGED` `google_compute_region_backend_service`. ([#3888](https://github.com/hashicorp/terraform-provider-google-beta/pull/3888))
* compute: fixed scenario where `instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3893](https://github.com/hashicorp/terraform-provider-google-beta/pull/3893))
* container: fixed the `ROUTES` value for the `networking_mode` field in `google_container_cluster`. A recent API change unintentionally changed the default to a `VPC_NATIVE` cluster, and removed the ability to create a `ROUTES`-based one. Provider versions prior to this one will default to `VPC_NATIVE` due to this change, and are unable to create `ROUTES` clusters. ([#3896](https://github.com/hashicorp/terraform-provider-google-beta/pull/3896))

## 4.3.0 (December 7, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))
* **New Data Source:** `google_folders` ([#3886](https://github.com/hashicorp/terraform-provider-google-beta/pull/3886))
* **New Resource:** `google_notebooks_runtime` ([#3878](https://github.com/hashicorp/terraform-provider-google-beta/pull/3878))
* **New Resource:** `google_vertex_ai_metadata_store` ([#3885](https://github.com/hashicorp/terraform-provider-google-beta/pull/3885))

IMPROVEMENTS
* apigee: Added IAM support for `google_apigee_environment`. ([#3871](https://github.com/hashicorp/terraform-provider-google-beta/pull/3871)):
* apigee: Added supported values for 'peeringCidrRange' in `google_apigee_instance`. ([#3880](https://github.com/hashicorp/terraform-provider-google-beta/pull/3880))
* cloudbuild: added display_name and annotations to google_cloudbuild_worker_pool for compatibility with new GA. ([#3873](https://github.com/hashicorp/terraform-provider-google-beta/pull/3873))
* container: added `node_group` to `node_config` for container clusters and node pools to support sole tenancy ([#3881](https://github.com/hashicorp/terraform-provider-google-beta/pull/3881))
* container: added `spot` field to `node_config` sub-resource ([#3863](https://github.com/hashicorp/terraform-provider-google-beta/pull/3863))
* redis: Added Multi read replica field `replicaCount `, `nodes`,  `readEndpoint`, `readEndpointPort`, `readReplicasMode` in `google_redis_instance` ([#3870](https://github.com/hashicorp/terraform-provider-google-beta/pull/3870))

BUG FIXES:
* essentialcontacts: marked updating `email` in `google_essential_contacts_contact` as requiring recreation ([#3864](https://github.com/hashicorp/terraform-provider-google-beta/pull/3864))
* privateca: fixed crlAccessUrls in `CertificateAuthority ` ([#3861](https://github.com/hashicorp/terraform-provider-google-beta/pull/3861))

## 4.2.1 (December 3, 2021)

BUG FIXES:
* provider: reverted a requirement in v4.2.0 for Terraform 0.13 and above. This release should be compatible with Terraform 0.12.31

## 4.2.0 (December 2, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))

IMPROVEMENTS:
* compute: added support for `queue_count` to `google_compute_instance.network_interface` and `google_compute_instance_template.network_interface` ([#3857](https://github.com/hashicorp/terraform-provider-google-beta/pull/3857))

BUG FIXES:
* all: fixed an issue where some documentation for new resources was not showing up in the GA provider if it was beta-only. ([#3848](https://github.com/hashicorp/terraform-provider-google-beta/pull/3848))
* bigquery: fixed update failure when attempting to change non-updatable fields in `google_bigquery_routine`. ([#3849](https://github.com/hashicorp/terraform-provider-google-beta/pull/3849))
* compute: fixed a bug that would cause `google_instance_from_machine_image` to fail with a resourceInUseByAnotherResource error ([#3855](https://github.com/hashicorp/terraform-provider-google-beta/pull/3855))
* compute: fixed a bug when `cache_mode` is set to FORCE_CACHE_ALL on `google_compute_backend_bucket` ([#3858](https://github.com/hashicorp/terraform-provider-google-beta/pull/3858))
* compute: fixed a perma-diff on `google_compute_region_health_check` when `log_config.enable` is set to false ([#3853](https://github.com/hashicorp/terraform-provider-google-beta/pull/3853))
* servicedirectory: added support for vpc network configuration in `google_service_directory_endpoint`. ([#3856](https://github.com/hashicorp/terraform-provider-google-beta/pull/3856))

## 4.1.0 (November 15, 2021)

IMPROVEMENTS:
* compute: Added `bfd` to `google_compute_router_peer` ([#3822](https://github.com/hashicorp/terraform-provider-google-beta/pull/3822))
* container: added `gcfs_config` to `node_config` of `google_container_node_pool` resource ([#3828](https://github.com/hashicorp/terraform-provider-google-beta/pull/3828))
* provider: added retries for the `resourceNotReady` error returned when attempting to add resources to a recently-modified subnetwork ([#3827](https://github.com/hashicorp/terraform-provider-google-beta/pull/3827))
* pubsub: added `message_retention_duration` field to `google_pubsub_topic` ([#3831](https://github.com/hashicorp/terraform-provider-google-beta/pull/3831))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance_attachment` could not be used on the same `google_apigee_instance` ([#3838](https://github.com/hashicorp/terraform-provider-google-beta/pull/3838))
* bigquery: fixed a bug following import where schema is empty on `google_bigquery_table` ([#3839](https://github.com/hashicorp/terraform-provider-google-beta/pull/3839))
* billingbudget: fixed unable to provide `labels` on `google_billing_budget` ([#3823](https://github.com/hashicorp/terraform-provider-google-beta/pull/3823))
* compute: allowed `source_disk` to accept full image path on `google_compute_snapshot` ([#3835](https://github.com/hashicorp/terraform-provider-google-beta/pull/3835))
* compute: fixed a bug in `google_compute_firewall` that would cause changes in `source_ranges` to not correctly be applied ([#3834](https://github.com/hashicorp/terraform-provider-google-beta/pull/3834))
* logging: fixed a bug with updating `description` on `google_logging_project_sink`, `google_logging_folder_sink` and `google_logging_organization_sink` ([#3826](https://github.com/hashicorp/terraform-provider-google-beta/pull/3826))

## 4.0.0 (November 02, 2021)

NOTES:
* compute: Google Compute Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3787](https://github.com/hashicorp/terraform-provider-google-beta/pull/3787))
* container: Google Kubernetes Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3788](https://github.com/hashicorp/terraform-provider-google-beta/pull/3788))

BREAKING CHANGES:
* appengine: marked `google_app_engine_standard_app_version` `entrypoint` as required ([#3784](https://github.com/hashicorp/terraform-provider-google-beta/pull/3784))
* compute: removed the ability to specify the `trace-append` or `trace-ro` as scopes in `google_compute_instance`, use `trace` instead ([#3759](https://github.com/hashicorp/terraform-provider-google-beta/pull/3759))
* compute: changed `advanced_machine_features` on `google_compute_instance_template` to track changes when the block is undefined in a user's config ([#3786](https://github.com/hashicorp/terraform-provider-google-beta/pull/3786))
* compute: changed `source_ranges` in `google_compute_firewall_rule` to track changes when it is not set in a config file ([#3791](https://github.com/hashicorp/terraform-provider-google-beta/pull/3791))
* compute: changed the import / drift detection behaviours for `metadata_startup_script`, `metadata.startup-script` in `google_compute_instance`. Now, `metadata.startup-script` will be set by default, and `metadata_startup_script` will only be set if present. ([#3765](https://github.com/hashicorp/terraform-provider-google-beta/pull/3765))
* compute: removed `source_disk_link` field from `google_compute_snapshot` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* container: `instance_group_urls` has been removed in favor of `node_pool.instance_group_urls` ([#3796](https://github.com/hashicorp/terraform-provider-google-beta/pull/3796))
* container: changed default for `enable_shielded_nodes` to true for `google_container_cluster` ([#3773](https://github.com/hashicorp/terraform-provider-google-beta/pull/3773))
* container: made `master_auth.client_certificate_config` required ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `master_auth.username` and `master_auth.password` from `google_container_cluster` ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `workload_metadata_configuration.node_metadata` in favor of `workload_metadata_configuration.mode` in `google_container_cluster` ([#3772](https://github.com/hashicorp/terraform-provider-google-beta/pull/3772))
* container: removed the `workload_identity_config.0.identity_namespace` field from `google_container_cluster`, use `workload_identity_config.0.workload_pool` instead ([#3776](https://github.com/hashicorp/terraform-provider-google-beta/pull/3776))
* kms: removed `self_link` field from `google_kms_crypto_key` and `google_kms_key_ring` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* project: removed ability to specify `bigquery-json.googleapis.com`, the provider will no longer convert it as the upstream API migration is finished. Use `bigquery.googleapis.com` instead. ([#3751](https://github.com/hashicorp/terraform-provider-google-beta/pull/3751))
* provider: changed `credentials`, `access_token` precedence so that `credentials` values in configuration take precedence over `access_token` values assigned through environment variables ([#3766](https://github.com/hashicorp/terraform-provider-google-beta/pull/3766))
* provider: removed redundant default scopes. The provider's default scopes when authenticating with credentials are now exclusively "https://www.googleapis.com/auth/cloud-platform" and "https://www.googleapis.com/auth/userinfo.email". ([#3756](https://github.com/hashicorp/terraform-provider-google-beta/pull/3756))
* pubsub: removed `path` from `google_pubsub_subscription` ([#3777](https://github.com/hashicorp/terraform-provider-google-beta/pull/3777))
* pubsub: removed `path` field from `google_pubsub_subscription` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* resourcemanager: made `google_project` remove `org_id` and `folder_id` from state when they are removed from config ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* resourcemanager: changed the `project` field to `Required` in all `google_project_iam_*` resources ([#3767](https://github.com/hashicorp/terraform-provider-google-beta/pull/3767))
* sql: added drift detection to the following `google_sql_database_instance` fields: `activation_policy` (defaults `ALWAYS`), `availability_type` (defaults `ZONAL`), `disk_type` (defaults `PD_SSD`), `encryption_key_name` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* sql: changed the `database_version` field to `Required` in `google_sql_database_instance` resource ([#3770](https://github.com/hashicorp/terraform-provider-google-beta/pull/3770))
* sql: removed the following `google_sql_database_instance` fields: `authorized_gae_applications`, `crash_safe_replication`, `replication_type` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* storage: removed `bucket_policy_only` from `google_storage_bucket` ([#3769](https://github.com/hashicorp/terraform-provider-google-beta/pull/3769))
* storage: changed the `location` field to required in `google_storage_bucket` ([#3771](https://github.com/hashicorp/terraform-provider-google-beta/pull/3771))

VALIDATION CHANGES:
* bigquery: at least one of `statement_timeout_ms`, `statement_byte_budget`, or `key_result_statement` is required on `google_bigquery_job.query.script_options.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `query`, `load`, `copy` or `extract` is required on `google_bigquery_job` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `source_table` or `source_model` is required on `google_bigquery_job.extract` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* cloudbuild: exactly one of `branch_name`, `commit_sha` or `tag_name` is required on `google_cloudbuild_trigger.build.source.repo_source` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed_delay` or `percentage` is required on `google_compute_url_map.default_route_action.fault_injection_policy.delay` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: required one of `source_tags`, `source_ranges` or `source_service_accounts` on INGRESS `google_compute_firewall` resources ([#3750](https://github.com/hashicorp/terraform-provider-google-beta/pull/3750))
* dlp: at least one of `start_time` or `end_time` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.timespan_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* dlp: exactly one of `url` or `regex_file_set` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.cloud_storage_options.file_set` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* resourcemanager: added conflict between `org_id`, `folder_id` at plan time in `google_project` ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.post_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.pre_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `reboot_config`, `apt`, `yum`, `goo` `zypper`, `windows_update`, `pre_step` or `pre_step` is required on `google_os_config_patch_deployment.patch_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `security`, `minimal`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.yum` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `type`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.apt` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `with_optional`, `with_update`, `categories`, `severities`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.patch_config.zypper` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: exactly one of `classifications`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.inspect_job.patch_config.windows_update` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* spanner: at least one of `num_nodes` or `processing_units` is required on `google_spanner_instance` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))

IMPROVEMENTS:
* container: added `managed_instance_group_urls` to `google_container_node_pool` to replace `instance_group_urls` on `google_container_cluster` ([#3815](https://github.com/hashicorp/terraform-provider-google-beta/pull/3815))
* kms: added support for EKM to `google_kms_crypto_key.protection_level` ([#3763](https://github.com/hashicorp/terraform-provider-google-beta/pull/3763))
* project: added support for `billing_project` on `google_project_service` ([#3768](https://github.com/hashicorp/terraform-provider-google-beta/pull/3768))
* spanner: increased the default timeout on `google_spanner_instance` operations from 4 minutes to 20 minutes, significantly reducing the likelihood that resources will time out ([#3789](https://github.com/hashicorp/terraform-provider-google-beta/pull/3789))

BUG FIXES:
* bigquery: fixed a bug of cannot add required fields to an existing schema on `google_bigquery_table` ([#3781](https://github.com/hashicorp/terraform-provider-google-beta/pull/3781))
* compute: fixed a bug in updating multiple `ttl` fields on `google_compute_backend_bucket` ([#3757](https://github.com/hashicorp/terraform-provider-google-beta/pull/3757))
* compute: fixed a perma-diff on `subnetwork` when it is optional on `google_compute_network_endpoint_group` ([#3780](https://github.com/hashicorp/terraform-provider-google-beta/pull/3780))
* compute: fixed perma-diff bug on `log_config.enable` of both `google_compute_backend_service` and `google_compute_region_backend_service` ([#3760](https://github.com/hashicorp/terraform-provider-google-beta/pull/3760))
* compute: fixed the `google_compute_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* compute: fixed the `google_compute_region_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* spanner: fixed the schema for `data.google_spanner_instance` so that non-configurable fields are considered outputs ([#3804](https://github.com/hashicorp/terraform-provider-google-beta/pull/3804))

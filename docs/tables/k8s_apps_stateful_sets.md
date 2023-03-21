# Table: k8s_apps_stateful_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status_ready_replicas | big_int | X | √ |  | 
| status_current_replicas | big_int | X | √ |  | 
| context | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_persistent_volume_claim_retention_policy | json | X | √ |  | 
| status_replicas | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| spec_replicas | big_int | X | √ |  | 
| status_update_revision | string | X | √ |  | 
| status_collision_count | big_int | X | √ |  | 
| spec_service_name | string | X | √ |  | 
| spec_revision_history_limit | big_int | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_ordinals | json | X | √ |  | 
| status_updated_replicas | big_int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_pod_management_policy | string | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| status_available_replicas | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| api_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_volume_claim_templates | json | X | √ |  | 
| status_current_revision | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_min_ready_seconds | big_int | X | √ |  | 



# Table: k8s_apps_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_references | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| uid | string | X | √ |  | 
| labels | json | X | √ |  | 
| spec_min_ready_seconds | big_int | X | √ |  | 
| status_ready_replicas | big_int | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_progress_deadline_seconds | big_int | X | √ |  | 
| status_updated_replicas | big_int | X | √ |  | 
| status_collision_count | big_int | X | √ |  | 
| name | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_strategy | json | X | √ |  | 
| kind | string | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_paused | bool | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| status_replicas | big_int | X | √ |  | 
| status_available_replicas | big_int | X | √ |  | 
| api_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| generation | big_int | X | √ |  | 
| spec_replicas | big_int | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_revision_history_limit | big_int | X | √ |  | 
| status_unavailable_replicas | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| context | string | X | √ |  | 
| namespace | string | X | √ |  | 



# Table: k8s_apps_daemon_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| finalizers | string_array | X | √ |  | 
| status_collision_count | big_int | X | √ |  | 
| status_number_available | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_min_ready_seconds | big_int | X | √ |  | 
| spec_revision_history_limit | big_int | X | √ |  | 
| resource_version | string | X | √ |  | 
| status_desired_number_scheduled | big_int | X | √ |  | 
| uid | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_updated_number_scheduled | big_int | X | √ |  | 
| context | string | X | √ |  | 
| namespace | string | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| status_number_unavailable | big_int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| labels | json | X | √ |  | 
| status_current_number_scheduled | big_int | X | √ |  | 
| generation | big_int | X | √ |  | 
| status_number_ready | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_number_misscheduled | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 



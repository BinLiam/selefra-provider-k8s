# Table: k8s_apps_replica_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| status_replicas | big_int | X | √ |  | 
| status_ready_replicas | big_int | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| labels | json | X | √ |  | 
| spec_min_ready_seconds | big_int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_fully_labeled_replicas | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_available_replicas | big_int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_replicas | big_int | X | √ |  | 



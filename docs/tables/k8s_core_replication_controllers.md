# Table: k8s_core_replication_controllers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status_ready_replicas | big_int | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| status_replicas | big_int | X | √ |  | 
| status_available_replicas | big_int | X | √ |  | 
| spec_template | json | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_min_ready_seconds | big_int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| api_version | string | X | √ |  | 
| spec_replicas | big_int | X | √ |  | 
| status_fully_labeled_replicas | big_int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 



# Table: k8s_autoscaling_hpas

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| finalizers | string_array | X | √ |  | 
| spec_min_replicas | big_int | X | √ |  | 
| status_last_scale_time | timestamp | X | √ |  | 
| status_current_replicas | big_int | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_scale_target_ref | json | X | √ |  | 
| spec_target_cpu_utilization_percentage | big_int | X | √ |  | 
| status_current_cpu_utilization_percentage | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_max_replicas | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| status_observed_generation | big_int | X | √ |  | 
| status_desired_replicas | big_int | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 



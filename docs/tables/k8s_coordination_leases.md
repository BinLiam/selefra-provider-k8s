# Table: k8s_coordination_leases

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_holder_identity | string | X | √ |  | 
| spec_acquire_time | json | X | √ |  | 
| spec_lease_transitions | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| spec_lease_duration_seconds | big_int | X | √ |  | 
| spec_renew_time | json | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| resource_version | string | X | √ |  | 



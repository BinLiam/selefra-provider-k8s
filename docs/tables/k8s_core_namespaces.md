# Table: k8s_core_namespaces

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_finalizers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| status_phase | string | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| status_conditions | json | X | √ |  | 



# Table: k8s_core_resource_quotas

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_scope_selector | json | X | √ |  | 
| uid | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| status_hard | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| spec_scopes | string_array | X | √ |  | 
| status_used | json | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_hard | json | X | √ |  | 



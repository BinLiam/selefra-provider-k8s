# Table: k8s_rbac_role_bindings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| subjects | json | X | √ |  | 
| role_ref | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| uid | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 



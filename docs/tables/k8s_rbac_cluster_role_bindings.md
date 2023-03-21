# Table: k8s_rbac_cluster_role_bindings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| subjects | json | X | √ |  | 
| uid | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| role_ref | json | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 



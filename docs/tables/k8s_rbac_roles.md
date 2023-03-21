# Table: k8s_rbac_roles

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| rules | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| resource_version | string | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| context | string | X | √ |  | 
| namespace | string | X | √ |  | 
| api_version | string | X | √ |  | 



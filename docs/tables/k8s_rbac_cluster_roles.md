# Table: k8s_rbac_cluster_roles

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rules | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| annotations | json | X | √ |  | 
| aggregation_rule | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 



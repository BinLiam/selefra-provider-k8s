# Table: k8s_core_endpoints

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| annotations | json | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| uid | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| subsets | json | X | √ |  | 



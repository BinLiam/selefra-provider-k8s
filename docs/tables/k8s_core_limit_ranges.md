# Table: k8s_core_limit_ranges

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| name | string | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| uid | string | X | √ |  | 
| spec_limits | json | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| context | string | X | √ |  | 



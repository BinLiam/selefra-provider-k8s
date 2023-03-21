# Table: k8s_core_config_maps

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| finalizers | string_array | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| immutable | bool | X | √ |  | 
| namespace | string | X | √ |  | 
| data | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| context | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| binary_data | json | X | √ |  | 



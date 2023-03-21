# Table: k8s_core_component_statuses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| conditions | json | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 



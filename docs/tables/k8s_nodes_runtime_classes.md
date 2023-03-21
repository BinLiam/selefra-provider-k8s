# Table: k8s_nodes_runtime_classes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_references | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| generation | big_int | X | √ |  | 
| overhead | json | X | √ |  | 
| context | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| handler | string | X | √ |  | 
| scheduling | json | X | √ |  | 



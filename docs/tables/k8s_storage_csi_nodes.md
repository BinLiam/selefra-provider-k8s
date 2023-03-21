# Table: k8s_storage_csi_nodes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_drivers | json | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| resource_version | string | X | √ |  | 



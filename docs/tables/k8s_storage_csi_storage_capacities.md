# Table: k8s_storage_csi_storage_capacities

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| node_topology | json | X | √ |  | 
| uid | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| storage_class_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| capacity | json | X | √ |  | 
| maximum_volume_size | json | X | √ |  | 



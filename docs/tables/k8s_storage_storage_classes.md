# Table: k8s_storage_storage_classes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| allow_volume_expansion | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| provisioner | string | X | √ |  | 
| parameters | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| namespace | string | X | √ |  | 
| reclaim_policy | string | X | √ |  | 
| mount_options | string_array | X | √ |  | 
| volume_binding_mode | string | X | √ |  | 
| allowed_topologies | json | X | √ |  | 



# Table: k8s_storage_csi_drivers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_requires_republish | bool | X | √ |  | 
| spec_se_linux_mount | bool | X | √ |  | 
| context | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_token_requests | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_pod_info_on_mount | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_storage_capacity | bool | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_attach_required | bool | X | √ |  | 
| spec_volume_lifecycle_modes | string_array | X | √ |  | 
| spec_fs_group_policy | string | X | √ |  | 



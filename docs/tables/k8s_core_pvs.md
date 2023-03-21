# Table: k8s_core_pvs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_volume_mode | string | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_capacity | json | X | √ |  | 
| spec_persistent_volume_reclaim_policy | string | X | √ |  | 
| status_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_persistent_volume_source | json | X | √ |  | 
| spec_claim_ref | json | X | √ |  | 
| spec_storage_class_name | string | X | √ |  | 
| spec_mount_options | string_array | X | √ |  | 
| status_phase | string | X | √ |  | 
| spec_access_modes | string_array | X | √ |  | 
| status_message | string | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_node_affinity | json | X | √ |  | 



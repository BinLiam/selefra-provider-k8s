# Table: k8s_core_pvcs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_data_source_ref | json | X | √ |  | 
| context | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_volume_mode | string | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_volume_name | string | X | √ |  | 
| spec_data_source | json | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_access_modes | string_array | X | √ |  | 
| name | string | X | √ |  | 
| status_resize_status | string | X | √ |  | 
| spec_resources | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_storage_class_name | string | X | √ |  | 
| status_capacity | json | X | √ |  | 
| api_version | string | X | √ |  | 
| status_allocated_resources | json | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| uid | string | X | √ |  | 
| status_phase | string | X | √ |  | 
| status_access_modes | string_array | X | √ |  | 
| kind | string | X | √ |  | 



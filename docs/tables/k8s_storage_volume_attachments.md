# Table: k8s_storage_volume_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_attacher | string | X | √ |  | 
| status_detach_error | json | X | √ |  | 
| spec_source | json | X | √ |  | 
| spec_node_name | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| status_attach_error | json | X | √ |  | 
| uid | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| status_attached | bool | X | √ |  | 
| status_attachment_metadata | json | X | √ |  | 



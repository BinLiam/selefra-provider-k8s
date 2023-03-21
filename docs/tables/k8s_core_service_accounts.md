# Table: k8s_core_service_accounts

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| image_pull_secrets | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| automount_service_account_token | bool | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| secrets | json | X | √ |  | 
| api_version | string | X | √ |  | 



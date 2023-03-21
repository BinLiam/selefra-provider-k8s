# Table: k8s_admissionregistration_mutating_webhook_configurations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| webhooks | json | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 



# Table: k8s_admissionregistration_validating_webhook_configurations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| webhooks | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| context | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| name | string | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| kind | string | X | √ |  | 



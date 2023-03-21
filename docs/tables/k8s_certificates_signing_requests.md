# Table: k8s_certificates_signing_requests

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_request | int_array | X | √ |  | 
| spec_expiration_seconds | big_int | X | √ |  | 
| spec_uid | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_extra | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_certificate | int_array | X | √ |  | 
| uid | string | X | √ |  | 
| spec_signer_name | string | X | √ |  | 
| spec_usages | string_array | X | √ |  | 
| spec_username | string | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_groups | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 



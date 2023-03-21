# Table: k8s_networking_ingress_classes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_parameters | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| api_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| spec_controller | string | X | √ |  | 
| name | string | X | √ |  | 
| annotations | json | X | √ |  | 



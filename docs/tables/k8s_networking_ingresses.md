# Table: k8s_networking_ingresses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| spec_default_backend | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_ingress_class_name | string | X | √ |  | 
| spec_tls | json | X | √ |  | 
| spec_rules | json | X | √ |  | 
| uid | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| status_load_balancer | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| finalizers | string_array | X | √ |  | 



# Table: k8s_networking_network_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| spec_ingress | json | X | √ |  | 
| context | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| name | string | X | √ |  | 
| api_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| spec_egress | json | X | √ |  | 
| spec_policy_types | string_array | X | √ |  | 
| kind | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_pod_selector | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| uid | string | X | √ |  | 



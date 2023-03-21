# Table: k8s_discovery_endpoint_slices

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| address_type | string | X | √ |  | 
| endpoints | json | X | √ |  | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| ports | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| uid | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 



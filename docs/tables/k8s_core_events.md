# Table: k8s_core_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| first_timestamp | timestamp | X | √ |  | 
| related | json | X | √ |  | 
| uid | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| message | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| kind | string | X | √ |  | 
| source | json | X | √ |  | 
| last_timestamp | timestamp | X | √ |  | 
| count | big_int | X | √ |  | 
| api_version | string | X | √ |  | 
| action | string | X | √ |  | 
| reporting_instance | string | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| reason | string | X | √ |  | 
| type | string | X | √ |  | 
| event_time | json | X | √ |  | 
| reporting_component | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| involved_object | json | X | √ |  | 
| series | json | X | √ |  | 



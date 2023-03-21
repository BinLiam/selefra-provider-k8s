# Table: k8s_core_nodes

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| uid | string | √ | √ |  | 
| spec_pod_cidr | cidr | X | √ |  | 
| spec | json | X | √ |  | 
| spec_pod_cidrs | cidr_array | X | √ |  | 
| type_meta | json | X | √ |  | 
| metadata | json | X | √ |  | 
| status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 



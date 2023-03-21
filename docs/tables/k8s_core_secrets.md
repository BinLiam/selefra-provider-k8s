# Table: k8s_core_secrets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| type_meta | json | X | √ |  | 
| immutable | bool | X | √ |  | 
| data | json | X | √ |  | 
| string_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uid | string | √ | √ |  | 
| metadata | json | X | √ |  | 
| type | string | X | √ |  | 



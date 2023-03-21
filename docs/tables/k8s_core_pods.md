# Table: k8s_core_pods

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | √ | √ |  | 
| status_host_ip | ip | X | √ |  | 
| status_pod_ips | ip_array | X | √ |  | 
| metadata | json | X | √ |  | 
| context | string | X | √ |  | 
| status_pod_ip | ip | X | √ |  | 
| type_meta | json | X | √ |  | 
| spec | json | X | √ |  | 
| status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 



# Table: k8s_core_services

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_load_balancer_source_ranges | string_array | X | √ |  | 
| spec_external_name | string | X | √ |  | 
| status_load_balancer | json | X | √ |  | 
| spec_load_balancer_ip | string | X | √ |  | 
| spec_session_affinity | string | X | √ |  | 
| spec_ip_families | string_array | X | √ |  | 
| spec_load_balancer_class | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| namespace | string | X | √ |  | 
| spec_internal_traffic_policy | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_ports | json | X | √ |  | 
| spec_session_affinity_config | json | X | √ |  | 
| spec_allocate_load_balancer_node_ports | bool | X | √ |  | 
| spec_cluster_ip | string | X | √ |  | 
| kind | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_ip_family_policy | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| spec_external_ips | string_array | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| spec_health_check_node_port | big_int | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | X | √ |  | 
| name | string | X | √ |  | 
| spec_external_traffic_policy | string | X | √ |  | 
| spec_publish_not_ready_addresses | bool | X | √ |  | 
| spec_cluster_ips | string_array | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| spec_type | string | X | √ |  | 



# Table: k8s_batch_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| spec_manual_selector | bool | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_completion_time | timestamp | X | √ |  | 
| generation | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_succeeded | big_int | X | √ |  | 
| status_failed | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_pod_failure_policy | json | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| status_start_time | timestamp | X | √ |  | 
| namespace | string | X | √ |  | 
| spec_parallelism | big_int | X | √ |  | 
| spec_ttl_seconds_after_finished | big_int | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| spec_completion_mode | string | X | √ |  | 
| status_ready | big_int | X | √ |  | 
| name | string | X | √ |  | 
| status_active | big_int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_completed_indexes | string | X | √ |  | 
| status_uncounted_terminated_pods | json | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| spec_active_deadline_seconds | big_int | X | √ |  | 
| spec_backoff_limit | big_int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| uid | string | X | √ |  | 
| spec_completions | big_int | X | √ |  | 



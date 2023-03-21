# Table: k8s_batch_cron_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| generation | big_int | X | √ |  | 
| deletion_grace_period_seconds | big_int | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_active | json | X | √ |  | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| status_last_schedule_time | timestamp | X | √ |  | 
| spec_time_zone | string | X | √ |  | 
| spec_concurrency_policy | string | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| spec_job_template | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| namespace | string | X | √ |  | 
| spec_schedule | string | X | √ |  | 
| labels | json | X | √ |  | 
| uid | string | X | √ |  | 
| api_version | string | X | √ |  | 
| spec_starting_deadline_seconds | big_int | X | √ |  | 
| spec_successful_jobs_history_limit | big_int | X | √ |  | 
| spec_failed_jobs_history_limit | big_int | X | √ |  | 
| status_last_successful_time | timestamp | X | √ |  | 
| kind | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 



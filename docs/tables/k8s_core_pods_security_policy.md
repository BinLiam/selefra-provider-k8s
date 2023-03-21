# Table: k8s_core_pods_security_policy

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Name of the object.  Name must be unique within a namespace. | 
| uid | string | √ | √ | UID is the unique in time and space value for this object. | 
| allow_privilege_escalation | bool | X | √ | Determines if a pod can request to allow privilege escalation. If unspecified, defaults to true. | 
| default_allow_privilege_escalation | bool | X | √ | Controls the default setting for whether a process can gain more privileges than its parent process. | 
| host_network | bool | X | √ | Determines if the policy allows the use of HostNetwork in the pod spec. | 
| host_ports | json | X | √ | Determines which host port ranges are allowed to be exposed. | 
| host_pid | bool | X | √ | Determines if the policy allows the use of HostPID in the pod spec. | 
| host_ipc | bool | X | √ | Determines if the policy allows the use of HostIPC in the pod spec. | 
| privileged | bool | X | √ | privileged determines if a pod can request to be run as privileged. | 
| read_only_root_filesystem | bool | X | √ | If set to true will force containers to run with a read only root file system. If set to false the container may run with a read only root file system if it wishes but it will not be forced to. | 
| allowed_csi_drivers | json | X | √ | An allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. | 
| allowed_host_paths | json | X | √ | An allowlist of host paths. Empty indicates that all host paths may be used. | 
| allowed_flex_volumes | json | X | √ | An allowlist of Flexvolumes. Empty or nil indicates that all Flexvolumes may be used. | 
| allowed_proc_mount_types | json | X | √ | An allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. | 
| allowed_unsafe_sysctls | json | X | √ | List of explicitly allowed unsafe sysctls, defaults to none. | 
| default_add_capabilities | json | X | √ | List of the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability. | 
| forbidden_sysctls | json | X | √ | List of explicitly forbidden sysctls, defaults to none. | 
| fs_group | json | X | √ | The strategy that will dictate what fs group is used by the SecurityContext. | 
| required_drop_capabilities | json | X | √ | List of the capabilities that will be dropped from the container. These are required to be dropped and cannot be added. | 
| run_as_group | json | X | √ | The strategy that will dictate the allowable RunAsGroup values that may be set. | 
| run_as_user | json | X | √ | The strategy that will dictate the allowable RunAsUser values that may be set. | 
| runtime_class | json | X | √ | The strategy that will dictate the allowable RuntimeClasses for a pod. | 
| se_linux | json | X | √ | The strategy that will dictate the allowable labels that may be set. | 
| supplemental_groups | json | X | √ | The strategy that will dictate what supplemental groups are used by the SecurityContext. | 
| volumes | json | X | √ | An allowlist of volume plugins. Empty indicates that no volumes may be used. | 
| title | string | X | √ | Title of the resource. | 
| tags | json | X | √ | A map of tags for the resource. This includes both labels and annotations. | 
| generate_name | string | X | √ | GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. | 
| resource_version | string | X | √ | An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. | 
| generation | big_int | X | √ | A sequence number representing a specific generation of the desired state. | 
| creation_timestamp | timestamp | X | √ | CreationTimestamp is a timestamp representing the server time when this object was created. | 
| deletion_timestamp | timestamp | X | √ | DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. | 
| deletion_grace_period_seconds | big_int | X | √ | Number of seconds allowed for this object to gracefully terminate before it will be removed from the system.  Only set when deletionTimestamp is also set. | 
| labels | json | X | √ | Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. | 
| annotations | json | X | √ | Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. | 
| owner_references | json | X | √ | List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller. | 
| finalizers | json | X | √ | Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. | 
| context_name | string | X | √ | Kubectl config context name. | 



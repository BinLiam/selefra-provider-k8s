package core

import (
	"context"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCorePodsSecurityPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCorePodsSecurityPolicyGenerator{}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetTableName() string {
	return "k8s_core_pods_security_policy"
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"uid",
		},
	}
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			diagnostics := schema.NewDiagnostics()
			k8sClient := client.(*k8s_client.Client).Client()

			var err error
			var response *v1beta1.PodSecurityPolicyList
			pageLeft := true

			input := metav1.ListOptions{
				Limit: 500,
			}

			for pageLeft {
				response, err = k8sClient.PolicyV1beta1().PodSecurityPolicies().List(ctx, input)
				if err != nil {
					return diagnostics.AddErrorMsg("pull security policy error: %s", err)
				}

				if response.GetContinue() != "" {
					input.Continue = response.Continue
				} else {
					pageLeft = false
				}

				for _, podSecurityPolicy := range response.Items {
					resultChannel <- podSecurityPolicy
				}
			}
			return nil
		},
	}
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func mergeTags(labels map[string]string, annotations map[string]string) map[string]string {
	tags := make(map[string]string)
	for k, v := range annotations {
		tags[k] = v
	}
	for k, v := range labels {
		tags[k] = v
	}
	return tags
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{

		// commons
		{
			ColumnName:  "name",
			Type:        schema.ColumnTypeString,
			Description: "Name of the object.  Name must be unique within a namespace.",
			Extractor:   column_value_extractor.StructSelector("ObjectMeta.Name"),
		},
		{
			ColumnName:  "uid",
			Type:        schema.ColumnTypeString,
			Description: "UID is the unique in time and space value for this object.",
			Extractor:   column_value_extractor.StructSelector("ObjectMeta.UID"),
		},

		// PodSecurityPolicySpec
		{
			ColumnName:  "allow_privilege_escalation",
			Type:        schema.ColumnTypeBool,
			Description: "Determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.",
			Extractor:   column_value_extractor.StructSelector("Spec.AllowPrivilegeEscalation"),
		},
		{
			ColumnName:  "default_allow_privilege_escalation",
			Type:        schema.ColumnTypeBool,
			Description: "Controls the default setting for whether a process can gain more privileges than its parent process.",
			Extractor:   column_value_extractor.StructSelector("Spec.DefaultAllowPrivilegeEscalation"),
		},
		{
			ColumnName:  "host_network",
			Type:        schema.ColumnTypeBool,
			Description: "Determines if the policy allows the use of HostNetwork in the pod spec.",
			Extractor:   column_value_extractor.StructSelector("Spec.HostNetwork"),
		},
		{
			ColumnName:  "host_ports",
			Type:        schema.ColumnTypeJSON,
			Description: "Determines which host port ranges are allowed to be exposed.",
			Extractor:   column_value_extractor.StructSelector("Spec.HostPorts"),
		},
		{
			ColumnName:  "host_pid",
			Type:        schema.ColumnTypeBool,
			Description: "Determines if the policy allows the use of HostPID in the pod spec.",
			Extractor:   column_value_extractor.StructSelector("Spec.HostPID"),
		},
		{
			ColumnName:  "host_ipc",
			Type:        schema.ColumnTypeBool,
			Description: "Determines if the policy allows the use of HostIPC in the pod spec.",
			Extractor:   column_value_extractor.StructSelector("Spec.HostIPC"),
		},
		{
			ColumnName:  "privileged",
			Type:        schema.ColumnTypeBool,
			Description: "privileged determines if a pod can request to be run as privileged.",
			Extractor:   column_value_extractor.StructSelector("Spec.Privileged"),
		},
		{
			ColumnName:  "read_only_root_filesystem",
			Type:        schema.ColumnTypeBool,
			Description: "If set to true will force containers to run with a read only root file system. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
			Extractor:   column_value_extractor.StructSelector("Spec.ReadOnlyRootFilesystem"),
		},

		// JSON Fields
		{
			ColumnName:  "allowed_csi_drivers",
			Type:        schema.ColumnTypeJSON,
			Description: "An allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec.",
			Extractor:   column_value_extractor.StructSelector("Spec.allowedCSIDrivers"),
		},
		{
			ColumnName:  "allowed_host_paths",
			Type:        schema.ColumnTypeJSON,
			Description: "An allowlist of host paths. Empty indicates that all host paths may be used.",
			Extractor:   column_value_extractor.StructSelector("Spec.AllowedHostPaths"),
		},
		{
			ColumnName:  "allowed_flex_volumes",
			Type:        schema.ColumnTypeJSON,
			Description: "An allowlist of Flexvolumes. Empty or nil indicates that all Flexvolumes may be used.",
			Extractor:   column_value_extractor.StructSelector("Spec.AllowedFlexVolumes"),
		},
		{
			ColumnName:  "allowed_proc_mount_types",
			Type:        schema.ColumnTypeJSON,
			Description: "An allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used.",
			Extractor:   column_value_extractor.StructSelector("Spec.AllowedProcMountTypes"),
		},
		{
			ColumnName:  "allowed_unsafe_sysctls",
			Type:        schema.ColumnTypeJSON,
			Description: "List of explicitly allowed unsafe sysctls, defaults to none.",
			Extractor:   column_value_extractor.StructSelector("Spec.AllowedUnsafeSysctls"),
		},
		{
			ColumnName:  "default_add_capabilities",
			Type:        schema.ColumnTypeJSON,
			Description: "List of the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.",
			Extractor:   column_value_extractor.StructSelector("Spec.DefaultAddCapabilities"),
		},
		{
			ColumnName:  "forbidden_sysctls",
			Type:        schema.ColumnTypeJSON,
			Description: "List of explicitly forbidden sysctls, defaults to none.",
			Extractor:   column_value_extractor.StructSelector("Spec.ForbiddenSysctls"),
		},
		{
			ColumnName:  "fs_group",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate what fs group is used by the SecurityContext.",
			Extractor:   column_value_extractor.StructSelector("Spec.FSGroup"),
		},
		{
			ColumnName:  "required_drop_capabilities",
			Type:        schema.ColumnTypeJSON,
			Description: "List of the capabilities that will be dropped from the container. These are required to be dropped and cannot be added.",
			Extractor:   column_value_extractor.StructSelector("Spec.RequiredDropCapabilities"),
		},
		{
			ColumnName:  "run_as_group",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate the allowable RunAsGroup values that may be set.",
			Extractor:   column_value_extractor.StructSelector("Spec.RunAsGroup"),
		},
		{
			ColumnName:  "run_as_user",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate the allowable RunAsUser values that may be set.",
			Extractor:   column_value_extractor.StructSelector("Spec.RunAsUser"),
		},
		{
			ColumnName:  "runtime_class",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate the allowable RuntimeClasses for a pod.",
			Extractor:   column_value_extractor.StructSelector("Spec.RuntimeClass"),
		},
		{
			ColumnName:  "se_linux",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate the allowable labels that may be set.",
			Extractor:   column_value_extractor.StructSelector("Spec.SELinux"),
		},
		{
			ColumnName:  "supplemental_groups",
			Type:        schema.ColumnTypeJSON,
			Description: "The strategy that will dictate what supplemental groups are used by the SecurityContext.",
			Extractor:   column_value_extractor.StructSelector("Spec.SupplementalGroups"),
		},
		{
			ColumnName:  "volumes",
			Type:        schema.ColumnTypeJSON,
			Description: "An allowlist of volume plugins. Empty indicates that no volumes may be used.",
			Extractor:   column_value_extractor.StructSelector("Spec.Volumes"),
		},

		{
			ColumnName:  "title",
			Type:        schema.ColumnTypeString,
			Description: "Title of the resource.",
			Extractor:   column_value_extractor.StructSelector("ObjectMeta.Name"),
		},
		{
			ColumnName:  "tags",
			Type:        schema.ColumnTypeJSON,
			Description: "A map of tags for the resource. This includes both labels and annotations.",
			//Transform:   transform.From(transformPodSecurityPolicyTags),
			Extractor: column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				if reflect_util.IsNil(result) {
					return nil, nil
				}
				obj := result.(v1beta1.PodSecurityPolicy)
				return mergeTags(obj.Labels, obj.Annotations), nil
			}),
		},
		{
			ColumnName:  "generate_name",
			Type:        schema.ColumnTypeString,
			Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided.",
		},
		{
			ColumnName:  "resource_version",
			Type:        schema.ColumnTypeString,
			Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed.",
		},
		{
			ColumnName:  "generation",
			Type:        schema.ColumnTypeBigInt,
			Description: "A sequence number representing a specific generation of the desired state.",
		},
		{
			ColumnName:  "creation_timestamp",
			Type:        schema.ColumnTypeTimestamp,
			Description: "CreationTimestamp is a timestamp representing the server time when this object was created."},
		{
			ColumnName:  "deletion_timestamp",
			Type:        schema.ColumnTypeTimestamp,
			Description: "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted."},
		{
			ColumnName:  "deletion_grace_period_seconds",
			Type:        schema.ColumnTypeBigInt,
			Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system.  Only set when deletionTimestamp is also set."},
		{
			ColumnName:  "labels",
			Type:        schema.ColumnTypeJSON,
			Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services."},
		{
			ColumnName:  "annotations",
			Type:        schema.ColumnTypeJSON,
			Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata."},
		{
			ColumnName:  "owner_references",
			Type:        schema.ColumnTypeJSON,
			Description: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller."},
		{
			ColumnName:  "finalizers",
			Type:        schema.ColumnTypeJSON,
			Description: "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed."},

		{
			ColumnName:  "context_name",
			Type:        schema.ColumnTypeString,
			Description: "Kubectl config context name.",
			Extractor: column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				config, err := client.(*k8s_client.Client).KubeConfig.RawConfig()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				}
				return config.CurrentContext, nil
			}),
		},
	}
}

func (x *TableK8sCorePodsSecurityPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}

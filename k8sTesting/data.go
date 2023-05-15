package k8sTesting

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	"github.com/selefra/selefra-provider-k8s/faker"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FakeThroughPointers(t *testing.T, ptrs ...interface{}) {
	for i, ptr := range ptrs {
		if err := faker.FakeObject(ptr); err != nil {
			t.Fatalf(constants.Vv, i, ptr)
		}
	}
}

func FakeDaemonSet(t *testing.T) appsv1.DaemonSet {
	var ds appsv1.DaemonSet
	ds.Spec.Template.Spec.Volumes = []corev1.Volume{FakeVolume(t)}
	FakeThroughPointers(t,
		&ds.TypeMeta,
		&ds.ObjectMeta,
		&ds.Status,
		&ds.ManagedFields,
		&ds.Annotations,
		&ds.Labels,
		&ds.OwnerReferences,
		&ds.Status,
		&ds.Spec.Selector,
		&ds.Spec.RevisionHistoryLimit,
	)
	ds.ManagedFields = []metav1.ManagedFieldsEntry{FakeManagedFields(t)}
	ds.Spec.Template = FakePodTemplateSpec(t)
	ds.Spec.Template.ManagedFields = []metav1.ManagedFieldsEntry{FakeManagedFields(t)}

	return ds
}

func FakeManagedFields(t *testing.T) metav1.ManagedFieldsEntry {
	m := metav1.ManagedFieldsEntry{}
	if err := faker.FakeObject(&m); err != nil {
		t.Fatal(err)
	}
	m.FieldsV1 = &metav1.FieldsV1{
		Raw: []byte(constants.Test),
	}
	return m
}

func FakePodTemplateSpec(t *testing.T) corev1.PodTemplateSpec {
	var templateSpec corev1.PodTemplateSpec
	if err := faker.FakeObject(&templateSpec); err != nil {
		t.Fatal(err)
	}
	templateSpec.Spec = FakePodSpec(t)
	templateSpec.ManagedFields = []metav1.ManagedFieldsEntry{FakeManagedFields(t)}
	return templateSpec
}

func FakeNode(t *testing.T) corev1.Node {

	var node corev1.Node
	FakeThroughPointers(t,
		&node.TypeMeta,
		&node.ObjectMeta,
		&node.Spec,
		&node.Status.Phase,
		&node.Status.Conditions,
		&node.Status.Addresses,
		&node.Status.DaemonEndpoints,
		&node.Status.NodeInfo,
		&node.Status.Images,
		&node.Status.VolumesInUse,
		&node.Status.VolumesAttached,
		&node.Status.Config,
	)
	node.Status.Capacity = *FakeResourceList(t)
	node.Status.Allocatable = *FakeResourceList(t)
	node.Spec.PodCIDR = constants.Constants_1
	node.Spec.PodCIDRs = []string{constants.Constants_2}
	node.Status.Addresses = []corev1.NodeAddress{
		{
			Type:    corev1.NodeHostName,
			Address: constants.Testname,
		},
		{
			Type:    corev1.NodeInternalIP,
			Address: constants.Fd,
		},
		{
			Type:    corev1.NodeExternalIP,
			Address: constants.Constants_3,
		},
	}
	return node
}

func FakeResourceList(t *testing.T) *corev1.ResourceList {
	rl := make(corev1.ResourceList)
	rq := apiresource.Quantity{}
	if err := faker.FakeObject(&rq); err != nil {
		t.Fatal(err)
	}
	var rn corev1.ResourceName
	if err := faker.FakeObject(&rn); err != nil {
		t.Fatal(err)
	}

	rl[rn] = rq
	return &rl
}

func FakeVolume(t *testing.T) corev1.Volume {

	var volume corev1.Volume
	FakeThroughPointers(t,
		&volume.Name,
		&volume.VolumeSource.HostPath,
		&volume.VolumeSource.EmptyDir,
		&volume.VolumeSource.GCEPersistentDisk,
		&volume.VolumeSource.AWSElasticBlockStore,
		&volume.VolumeSource.GitRepo,
		&volume.VolumeSource.Secret,
		&volume.VolumeSource.NFS,
		&volume.VolumeSource.ISCSI,
		&volume.VolumeSource.Glusterfs,
		&volume.VolumeSource.PersistentVolumeClaim,
		&volume.VolumeSource.RBD,
		&volume.VolumeSource.FlexVolume,
		&volume.VolumeSource.Cinder,
		&volume.VolumeSource.CephFS,
		&volume.VolumeSource.Flocker,
		&volume.VolumeSource.DownwardAPI,
		&volume.VolumeSource.FC,
		&volume.VolumeSource.AzureFile,
		&volume.VolumeSource.ConfigMap,
		&volume.VolumeSource.VsphereVolume,
		&volume.VolumeSource.Quobyte,
		&volume.VolumeSource.AzureDisk,
		&volume.VolumeSource.PhotonPersistentDisk,
		&volume.VolumeSource.Projected,
		&volume.VolumeSource.PortworxVolume,
		&volume.VolumeSource.ScaleIO,
		&volume.VolumeSource.StorageOS,
		&volume.VolumeSource.CSI,
	)
	volume.Ephemeral = &corev1.EphemeralVolumeSource{}
	return volume
}

func FakeContainer(t *testing.T) corev1.Container {
	var c corev1.Container
	FakeThroughPointers(t,
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,

		&c.VolumeMounts,
		&c.VolumeDevices,

		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	)

	c.Resources.Limits = *FakeResourceList(t)
	c.Resources.Requests = *FakeResourceList(t)
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func FakeEphemeralContainer(t *testing.T) corev1.EphemeralContainer {
	var c corev1.EphemeralContainer
	FakeThroughPointers(t,
		&c.TargetContainerName,
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,

		&c.VolumeMounts,
		&c.VolumeDevices,

		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	)

	c.Resources.Limits = *FakeResourceList(t)
	c.Resources.Requests = *FakeResourceList(t)
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func FakePod(t *testing.T) corev1.Pod {
	var pod corev1.Pod
	pod.Spec.Volumes = []corev1.Volume{FakeVolume(t)}
	FakeThroughPointers(t,
		&pod.TypeMeta,
		&pod.ObjectMeta,
		&pod.Status,
	)
	pod.Spec = FakePodSpec(t)

	pod.Status.HostIP = constants.Constants_4
	pod.Status.PodIP = constants.Constants_5
	pod.Status.PodIPs = []corev1.PodIP{{IP: constants.Constants_6}}
	return pod
}

func FakePodSpec(t *testing.T) corev1.PodSpec {
	var podSpec corev1.PodSpec
	if err := faker.FakeObject(&podSpec); err != nil {
		t.Fatal(err)
	}
	podSpec.Overhead = *FakeResourceList(t)
	podSpec.InitContainers = []corev1.Container{FakeContainer(t)}
	podSpec.Containers = []corev1.Container{FakeContainer(t)}
	podSpec.EphemeralContainers = []corev1.EphemeralContainer{FakeEphemeralContainer(t)}
	podSpec.Volumes = []corev1.Volume{FakeVolume(t)}
	podSpec.RestartPolicy = "test"

	return podSpec
}

func FakeSelector(_ *testing.T) *metav1.LabelSelector {
	return &metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "test",
				Operator: "test",
				Values: []string{
					"test1", "test2",
				},
			},
		},
		MatchLabels: map[string]string{
			"test": "test",
		},
	}
}

func FakePersistentVolumeClaim(t *testing.T) *corev1.PersistentVolumeClaim {
	claim := corev1.PersistentVolumeClaim{}
	if err := faker.FakeObject(&claim); err != nil {
		t.Fatal(err)
	}
	claim.ManagedFields = []metav1.ManagedFieldsEntry{FakeManagedFields(t)}
	claim.Status.Capacity = *FakeResourceList(t)
	claim.Spec.Resources.Requests = *FakeResourceList(t)
	claim.Spec.Resources.Limits = *FakeResourceList(t)

	return &claim
}

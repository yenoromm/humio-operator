package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HumioClusterStateBoostrapping is the Bootstrapping state of the cluster
	HumioClusterStateBoostrapping = "Bootstrapping"
	// HumioClusterStateRunning is the Running state of the cluster
	HumioClusterStateRunning = "Running"
)

// HumioClusterSpec defines the desired state of HumioCluster
type HumioClusterSpec struct {
	// Desired container image including the image tag
	Image string `json:"image,omitempty"`
	// Desired number of replicas of both storage and ingest partitions
	TargetReplicationFactor int `json:"targetReplicationFactor,omitempty"`
	// Desired number of storage partitions
	StoragePartitionsCount int `json:"storagePartitionsCount,omitempty"`
	// Desired number of digest partitions
	DigestPartitionsCount int `json:"digestPartitionsCount,omitempty"`
	// Desired number of nodes
	NodeCount int `json:"nodeCount,omitempty"`
	// Extra environment variables
	EnvironmentVariables []corev1.EnvVar `json:"environmentVariables,omitempty"`
}

// HumioClusterStatus defines the observed state of HumioCluster
type HumioClusterStatus struct {
	StateLastUpdatedUnix int64 `json:"stateLastUpdated,omitempty"`
	// Current state set by operator.
	AllDataAvailable string `json:"allDataAvailable,omitempty"`
	// ClusterState will be empty before the cluster is bootstrapped. From there it can be "Bootstrapping" or "Running"
	// TODO: other states?
	ClusterState string `json:"clusterState,omitempty"`
	// ClusterVersion is the version of humio running
	ClusterVersion string `json:"clusterVersion,omitempty"`
	// ClusterNodeCount is the number of nodes of humio running
	ClusterNodeCount int `json:"clusterNodeCount,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HumioCluster is the Schema for the humioclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=humioclusters,scope=Namespaced
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.clusterState",description="The state of the cluster"
// +kubebuilder:printcolumn:name="Nodes",type="string",JSONPath=".status.clusterNodeCount",description="The number of nodes in the cluster"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".status.clusterVersion",description="The version of humior"
type HumioCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HumioClusterSpec   `json:"spec,omitempty"`
	Status HumioClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HumioClusterList contains a list of HumioCluster
type HumioClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HumioCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HumioCluster{}, &HumioClusterList{})
}

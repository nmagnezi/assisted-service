/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MacInterfaceMapItems struct {
	// nic name used in the yaml, which relates 1:1 to the mac address
	LogicalNicName string `json:"logicalNICName"`
	// mac address present on the host.
	// Pattern: ^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$
	MacAddress string `json:"macAddress"`
}

// +kubebuilder:validation:Type=object
type RawState []byte

type State struct {
	Raw RawState `json:""`
}

type NMStateConfigSpec struct {
	// mapping of host macs to logical interfaces used in the network yaml
	MacInterfaceMap []*MacInterfaceMapItems `json:"macInterfaceMap,omitempty"`
	// +kubebuilder:validation:XPreserveUnknownFields
	// yaml string that can be processed by nmstate
	NetworkConfig State `json:"networkConfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type NMStateConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NMStateConfigSpec `json:"spec,omitempty"`
	// No status
}

// +kubebuilder:object:root=true

// NMStateConfigList contains a list of NMStateConfigs
type NMStateConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NMStateConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NMStateConfig{}, &NMStateConfigList{})
}

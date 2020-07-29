/*
Copyright 2020 The OpenEBS Authors.

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

package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=cstorrestore

// CStorRestore describes a cstor restore resource created as a custom resource
type CStorRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"` // set name to restore name + volume name + something like cspi tag
	Spec              CStorRestoreSpec            `json:"spec"`
	Status            CStorRestoreStatus          `json:"status"`
}

// CStorRestoreSpec is the spec for a CStorRestore resource
type CStorRestoreSpec struct {
	// RestoreName holds restore name
	RestoreName string `json:"restoreName,omitempty"`
	// VolumeName is used to restore the data to corresponding volume
	VolumeName string `json:"volumeName,omitempty"`
	// RestoreSrc can be ip:port in case of restore from remote or volumeName
	// in case of local restore
	RestoreSrc string `json:"restoreSrc,omitempty"`
	// MaxRestoreRetryCount is the maximum number of attempt, will be performed to restore
	MaxRetryCount int `json:"maxretrycount,omitempty"`
	// RetryCount represents the number of restore attempts performed for the restore
	RetryCount int `json:"retrycount,omitempty"`
	// StorageClass represents name of StorageClass of restore volume
	StorageClass string `json:"storageClass,omitempty"`
	// Size represents the size of a snapshot to restore
	Size resource.Quantity `json:"size,omitempty"`
	// Local defines whether restore is from local/remote
	Local bool `json:"localRestore,omitempty"` // if restore is from local backup/snapshot
}

// CStorRestoreStatus is to hold result of action.
type CStorRestoreStatus string

// Status written onto CStrorRestore object.
const (
	// RSTCStorStatusEmpty ensures the create operation is to be done, if import fails.
	RSTCStorStatusEmpty CStorRestoreStatus = ""

	// RSTCStorStatusDone , restore operation is completed.
	RSTCStorStatusDone CStorRestoreStatus = "Done"

	// RSTCStorStatusFailed , restore operation is failed.
	RSTCStorStatusFailed CStorRestoreStatus = "Failed"

	// RSTCStorStatusInit , restore operation is initialized.
	RSTCStorStatusInit CStorRestoreStatus = "Init"

	// RSTCStorStatusPending , restore operation is pending.
	RSTCStorStatusPending CStorRestoreStatus = "Pending"

	// RSTCStorStatusInProgress , restore operation is in progress.
	RSTCStorStatusInProgress CStorRestoreStatus = "InProgress"

	// RSTCStorStatusInvalid , restore operation is invalid.
	RSTCStorStatusInvalid CStorRestoreStatus = "Invalid"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=cstorrestore

// CStorRestoreList is a list of CStorRestore resources
type CStorRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CStorRestore `json:"items"`
}

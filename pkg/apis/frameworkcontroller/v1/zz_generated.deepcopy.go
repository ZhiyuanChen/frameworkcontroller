// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionCodeInfo) DeepCopyInto(out *CompletionCodeInfo) {
	*out = *in
	in.Type.DeepCopyInto(&out.Type)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionCodeInfo.
func (in *CompletionCodeInfo) DeepCopy() *CompletionCodeInfo {
	if in == nil {
		return nil
	}
	out := new(CompletionCodeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionPolicySpec) DeepCopyInto(out *CompletionPolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionPolicySpec.
func (in *CompletionPolicySpec) DeepCopy() *CompletionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CompletionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionStatus) DeepCopyInto(out *CompletionStatus) {
	*out = *in
	in.Type.DeepCopyInto(&out.Type)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionStatus.
func (in *CompletionStatus) DeepCopy() *CompletionStatus {
	if in == nil {
		return nil
	}
	out := new(CompletionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionType) DeepCopyInto(out *CompletionType) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]CompletionTypeAttribute, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionType.
func (in *CompletionType) DeepCopy() *CompletionType {
	if in == nil {
		return nil
	}
	out := new(CompletionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfig) DeepCopyInto(out *ControllerConfig) {
	*out = *in
	if in.KubeApiServerAddress != nil {
		in, out := &in.KubeApiServerAddress, &out.KubeApiServerAddress
		*out = new(string)
		**out = **in
	}
	if in.KubeConfigFilePath != nil {
		in, out := &in.KubeConfigFilePath, &out.KubeConfigFilePath
		*out = new(string)
		**out = **in
	}
	if in.WorkerNumber != nil {
		in, out := &in.WorkerNumber, &out.WorkerNumber
		*out = new(int32)
		**out = **in
	}
	if in.CRDEstablishedCheckIntervalSec != nil {
		in, out := &in.CRDEstablishedCheckIntervalSec, &out.CRDEstablishedCheckIntervalSec
		*out = new(int64)
		**out = **in
	}
	if in.CRDEstablishedCheckTimeoutSec != nil {
		in, out := &in.CRDEstablishedCheckTimeoutSec, &out.CRDEstablishedCheckTimeoutSec
		*out = new(int64)
		**out = **in
	}
	if in.ObjectLocalCacheCreationTimeoutSec != nil {
		in, out := &in.ObjectLocalCacheCreationTimeoutSec, &out.ObjectLocalCacheCreationTimeoutSec
		*out = new(int64)
		**out = **in
	}
	if in.FrameworkMinRetryDelaySecForTransientConflictFailed != nil {
		in, out := &in.FrameworkMinRetryDelaySecForTransientConflictFailed, &out.FrameworkMinRetryDelaySecForTransientConflictFailed
		*out = new(int64)
		**out = **in
	}
	if in.FrameworkMaxRetryDelaySecForTransientConflictFailed != nil {
		in, out := &in.FrameworkMaxRetryDelaySecForTransientConflictFailed, &out.FrameworkMaxRetryDelaySecForTransientConflictFailed
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfig.
func (in *ControllerConfig) DeepCopy() *ControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Framework) DeepCopyInto(out *Framework) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(FrameworkStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Framework.
func (in *Framework) DeepCopy() *Framework {
	if in == nil {
		return nil
	}
	out := new(Framework)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Framework) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkAttemptStatus) DeepCopyInto(out *FrameworkAttemptStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.InstanceUID != nil {
		in, out := &in.InstanceUID, &out.InstanceUID
		*out = new(types.UID)
		**out = **in
	}
	if in.ConfigMapUID != nil {
		in, out := &in.ConfigMapUID, &out.ConfigMapUID
		*out = new(types.UID)
		**out = **in
	}
	if in.CompletionStatus != nil {
		in, out := &in.CompletionStatus, &out.CompletionStatus
		*out = new(CompletionStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskRoleStatuses != nil {
		in, out := &in.TaskRoleStatuses, &out.TaskRoleStatuses
		*out = make([]TaskRoleStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkAttemptStatus.
func (in *FrameworkAttemptStatus) DeepCopy() *FrameworkAttemptStatus {
	if in == nil {
		return nil
	}
	out := new(FrameworkAttemptStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkList) DeepCopyInto(out *FrameworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Framework, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkList.
func (in *FrameworkList) DeepCopy() *FrameworkList {
	if in == nil {
		return nil
	}
	out := new(FrameworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrameworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkSpec) DeepCopyInto(out *FrameworkSpec) {
	*out = *in
	out.RetryPolicy = in.RetryPolicy
	if in.TaskRoles != nil {
		in, out := &in.TaskRoles, &out.TaskRoles
		*out = make([]TaskRoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkSpec.
func (in *FrameworkSpec) DeepCopy() *FrameworkSpec {
	if in == nil {
		return nil
	}
	out := new(FrameworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkStatus) DeepCopyInto(out *FrameworkStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	in.TransitionTime.DeepCopyInto(&out.TransitionTime)
	in.RetryPolicyStatus.DeepCopyInto(&out.RetryPolicyStatus)
	in.AttemptStatus.DeepCopyInto(&out.AttemptStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkStatus.
func (in *FrameworkStatus) DeepCopy() *FrameworkStatus {
	if in == nil {
		return nil
	}
	out := new(FrameworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryDecision) DeepCopyInto(out *RetryDecision) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryDecision.
func (in *RetryDecision) DeepCopy() *RetryDecision {
	if in == nil {
		return nil
	}
	out := new(RetryDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryPolicySpec) DeepCopyInto(out *RetryPolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicySpec.
func (in *RetryPolicySpec) DeepCopy() *RetryPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RetryPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryPolicyStatus) DeepCopyInto(out *RetryPolicyStatus) {
	*out = *in
	if in.RetryDelaySec != nil {
		in, out := &in.RetryDelaySec, &out.RetryDelaySec
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicyStatus.
func (in *RetryPolicyStatus) DeepCopy() *RetryPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RetryPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskAttemptStatus) DeepCopyInto(out *TaskAttemptStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.InstanceUID != nil {
		in, out := &in.InstanceUID, &out.InstanceUID
		*out = new(types.UID)
		**out = **in
	}
	if in.PodUID != nil {
		in, out := &in.PodUID, &out.PodUID
		*out = new(types.UID)
		**out = **in
	}
	if in.PodIP != nil {
		in, out := &in.PodIP, &out.PodIP
		*out = new(string)
		**out = **in
	}
	if in.PodHostIP != nil {
		in, out := &in.PodHostIP, &out.PodHostIP
		*out = new(string)
		**out = **in
	}
	if in.CompletionStatus != nil {
		in, out := &in.CompletionStatus, &out.CompletionStatus
		*out = new(CompletionStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskAttemptStatus.
func (in *TaskAttemptStatus) DeepCopy() *TaskAttemptStatus {
	if in == nil {
		return nil
	}
	out := new(TaskAttemptStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRoleSpec) DeepCopyInto(out *TaskRoleSpec) {
	*out = *in
	out.FrameworkAttemptCompletionPolicy = in.FrameworkAttemptCompletionPolicy
	in.Task.DeepCopyInto(&out.Task)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRoleSpec.
func (in *TaskRoleSpec) DeepCopy() *TaskRoleSpec {
	if in == nil {
		return nil
	}
	out := new(TaskRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRoleStatus) DeepCopyInto(out *TaskRoleStatus) {
	*out = *in
	if in.TaskStatuses != nil {
		in, out := &in.TaskStatuses, &out.TaskStatuses
		*out = make([]TaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRoleStatus.
func (in *TaskRoleStatus) DeepCopy() *TaskRoleStatus {
	if in == nil {
		return nil
	}
	out := new(TaskRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	out.RetryPolicy = in.RetryPolicy
	in.Pod.DeepCopyInto(&out.Pod)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	in.TransitionTime.DeepCopyInto(&out.TransitionTime)
	in.RetryPolicyStatus.DeepCopyInto(&out.RetryPolicyStatus)
	in.AttemptStatus.DeepCopyInto(&out.AttemptStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

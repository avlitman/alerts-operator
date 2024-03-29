//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtAlert) DeepCopyInto(out *KubevirtAlert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtAlert.
func (in *KubevirtAlert) DeepCopy() *KubevirtAlert {
	if in == nil {
		return nil
	}
	out := new(KubevirtAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubevirtAlert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtAlertList) DeepCopyInto(out *KubevirtAlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubevirtAlert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtAlertList.
func (in *KubevirtAlertList) DeepCopy() *KubevirtAlertList {
	if in == nil {
		return nil
	}
	out := new(KubevirtAlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubevirtAlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtAlertSpec) DeepCopyInto(out *KubevirtAlertSpec) {
	*out = *in
	out.Metric = in.Metric
	in.RecordRule.DeepCopyInto(&out.RecordRule)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtAlertSpec.
func (in *KubevirtAlertSpec) DeepCopy() *KubevirtAlertSpec {
	if in == nil {
		return nil
	}
	out := new(KubevirtAlertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubevirtAlertStatus) DeepCopyInto(out *KubevirtAlertStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubevirtAlertStatus.
func (in *KubevirtAlertStatus) DeepCopy() *KubevirtAlertStatus {
	if in == nil {
		return nil
	}
	out := new(KubevirtAlertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordRuleSpec) DeepCopyInto(out *RecordRuleSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordRuleSpec.
func (in *RecordRuleSpec) DeepCopy() *RecordRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RecordRuleSpec)
	in.DeepCopyInto(out)
	return out
}

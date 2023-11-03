/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "bytedeath.com/zflow/pkg/apis/zflowcontroller/v1"
	zflowcontrollerv1 "bytedeath.com/zflow/pkg/generated/applyconfiguration/zflowcontroller/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=zflowcontroller.bytedeath.com, Version=v1
	case v1.SchemeGroupVersion.WithKind("Ticket"):
		return &zflowcontrollerv1.TicketApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TicketSpec"):
		return &zflowcontrollerv1.TicketSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TicketStatus"):
		return &zflowcontrollerv1.TicketStatusApplyConfiguration{}

	}
	return nil
}

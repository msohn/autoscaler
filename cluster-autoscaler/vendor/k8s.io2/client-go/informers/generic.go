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

// This file was automatically generated by informer-gen

package informers

import (
	"fmt"
	v1alpha1 "k8s.io2/api/admissionregistration/v1alpha1"
	v1beta1 "k8s.io2/api/admissionregistration/v1beta1"
	v1 "k8s.io2/api/apps/v1"
	apps_v1beta1 "k8s.io2/api/apps/v1beta1"
	v1beta2 "k8s.io2/api/apps/v1beta2"
	autoscaling_v1 "k8s.io2/api/autoscaling/v1"
	v2beta1 "k8s.io2/api/autoscaling/v2beta1"
	batch_v1 "k8s.io2/api/batch/v1"
	batch_v1beta1 "k8s.io2/api/batch/v1beta1"
	v2alpha1 "k8s.io2/api/batch/v2alpha1"
	certificates_v1beta1 "k8s.io2/api/certificates/v1beta1"
	core_v1 "k8s.io2/api/core/v1"
	events_v1beta1 "k8s.io2/api/events/v1beta1"
	extensions_v1beta1 "k8s.io2/api/extensions/v1beta1"
	networking_v1 "k8s.io2/api/networking/v1"
	policy_v1beta1 "k8s.io2/api/policy/v1beta1"
	rbac_v1 "k8s.io2/api/rbac/v1"
	rbac_v1alpha1 "k8s.io2/api/rbac/v1alpha1"
	rbac_v1beta1 "k8s.io2/api/rbac/v1beta1"
	scheduling_v1alpha1 "k8s.io2/api/scheduling/v1alpha1"
	settings_v1alpha1 "k8s.io2/api/settings/v1alpha1"
	storage_v1 "k8s.io2/api/storage/v1"
	storage_v1alpha1 "k8s.io2/api/storage/v1alpha1"
	storage_v1beta1 "k8s.io2/api/storage/v1beta1"
	schema "k8s.io2/apimachinery/pkg/runtime/schema"
	cache "k8s.io2/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=admissionregistration.k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("initializerconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().V1alpha1().InitializerConfigurations().Informer()}, nil

		// Group=admissionregistration.k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().V1beta1().MutatingWebhookConfigurations().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("validatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().V1beta1().ValidatingWebhookConfigurations().Informer()}, nil

		// Group=apps, Version=v1
	case v1.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().ControllerRevisions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().DaemonSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().Deployments().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().ReplicaSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1().StatefulSets().Informer()}, nil

		// Group=apps, Version=v1beta1
	case apps_v1beta1.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ControllerRevisions().Informer()}, nil
	case apps_v1beta1.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Deployments().Informer()}, nil
	case apps_v1beta1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().StatefulSets().Informer()}, nil

		// Group=apps, Version=v1beta2
	case v1beta2.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta2().ControllerRevisions().Informer()}, nil
	case v1beta2.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta2().DaemonSets().Informer()}, nil
	case v1beta2.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta2().Deployments().Informer()}, nil
	case v1beta2.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta2().ReplicaSets().Informer()}, nil
	case v1beta2.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta2().StatefulSets().Informer()}, nil

		// Group=autoscaling, Version=v1
	case autoscaling_v1.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().V1().HorizontalPodAutoscalers().Informer()}, nil

		// Group=autoscaling, Version=v2beta1
	case v2beta1.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().V2beta1().HorizontalPodAutoscalers().Informer()}, nil

		// Group=batch, Version=v1
	case batch_v1.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().V1().Jobs().Informer()}, nil

		// Group=batch, Version=v1beta1
	case batch_v1beta1.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().V1beta1().CronJobs().Informer()}, nil

		// Group=batch, Version=v2alpha1
	case v2alpha1.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().V2alpha1().CronJobs().Informer()}, nil

		// Group=certificates.k8s.io, Version=v1beta1
	case certificates_v1beta1.SchemeGroupVersion.WithResource("certificatesigningrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certificates().V1beta1().CertificateSigningRequests().Informer()}, nil

		// Group=core, Version=v1
	case core_v1.SchemeGroupVersion.WithResource("componentstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ComponentStatuses().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ConfigMaps().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Endpoints().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Events().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("limitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().LimitRanges().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Namespaces().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Nodes().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("persistentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumes().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("persistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumeClaims().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Pods().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("podtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PodTemplates().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("replicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ReplicationControllers().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ResourceQuotas().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Secrets().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Services().Informer()}, nil
	case core_v1.SchemeGroupVersion.WithResource("serviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ServiceAccounts().Informer()}, nil

		// Group=events.k8s.io, Version=v1beta1
	case events_v1beta1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Events().V1beta1().Events().Informer()}, nil

		// Group=extensions, Version=v1beta1
	case extensions_v1beta1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1beta1().DaemonSets().Informer()}, nil
	case extensions_v1beta1.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1beta1().Deployments().Informer()}, nil
	case extensions_v1beta1.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1beta1().Ingresses().Informer()}, nil
	case extensions_v1beta1.SchemeGroupVersion.WithResource("podsecuritypolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1beta1().PodSecurityPolicies().Informer()}, nil
	case extensions_v1beta1.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1beta1().ReplicaSets().Informer()}, nil

		// Group=networking.k8s.io, Version=v1
	case networking_v1.SchemeGroupVersion.WithResource("networkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1().NetworkPolicies().Informer()}, nil

		// Group=policy, Version=v1beta1
	case policy_v1beta1.SchemeGroupVersion.WithResource("poddisruptionbudgets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().V1beta1().PodDisruptionBudgets().Informer()}, nil

		// Group=rbac.authorization.k8s.io, Version=v1
	case rbac_v1.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().ClusterRoles().Informer()}, nil
	case rbac_v1.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().ClusterRoleBindings().Informer()}, nil
	case rbac_v1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().Roles().Informer()}, nil
	case rbac_v1.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1().RoleBindings().Informer()}, nil

		// Group=rbac.authorization.k8s.io, Version=v1alpha1
	case rbac_v1alpha1.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().ClusterRoles().Informer()}, nil
	case rbac_v1alpha1.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().ClusterRoleBindings().Informer()}, nil
	case rbac_v1alpha1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().Roles().Informer()}, nil
	case rbac_v1alpha1.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().RoleBindings().Informer()}, nil

		// Group=rbac.authorization.k8s.io, Version=v1beta1
	case rbac_v1beta1.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1beta1().ClusterRoles().Informer()}, nil
	case rbac_v1beta1.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1beta1().ClusterRoleBindings().Informer()}, nil
	case rbac_v1beta1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1beta1().Roles().Informer()}, nil
	case rbac_v1beta1.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1beta1().RoleBindings().Informer()}, nil

		// Group=scheduling.k8s.io, Version=v1alpha1
	case scheduling_v1alpha1.SchemeGroupVersion.WithResource("priorityclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().PriorityClasses().Informer()}, nil

		// Group=settings.k8s.io, Version=v1alpha1
	case settings_v1alpha1.SchemeGroupVersion.WithResource("podpresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Settings().V1alpha1().PodPresets().Informer()}, nil

		// Group=storage.k8s.io, Version=v1
	case storage_v1.SchemeGroupVersion.WithResource("storageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().StorageClasses().Informer()}, nil

		// Group=storage.k8s.io, Version=v1alpha1
	case storage_v1alpha1.SchemeGroupVersion.WithResource("volumeattachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().VolumeAttachments().Informer()}, nil

		// Group=storage.k8s.io, Version=v1beta1
	case storage_v1beta1.SchemeGroupVersion.WithResource("storageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1beta1().StorageClasses().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

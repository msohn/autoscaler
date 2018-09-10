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

package fake

import (
	apps_v1 "k8s.io2/api/apps/v1"
	v1 "k8s.io2/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io2/apimachinery/pkg/labels"
	schema "k8s.io2/apimachinery/pkg/runtime/schema"
	types "k8s.io2/apimachinery/pkg/types"
	watch "k8s.io2/apimachinery/pkg/watch"
	testing "k8s.io2/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeAppsV1
	ns   string
}

var deploymentsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

var deploymentsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(name string, options v1.GetOptions) (result *apps_v1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentsResource, c.ns, name), &apps_v1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(opts v1.ListOptions) (result *apps_v1.DeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentsResource, deploymentsKind, c.ns, opts), &apps_v1.DeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apps_v1.DeploymentList{}
	for _, item := range obj.(*apps_v1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentsResource, c.ns, opts))

}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(deployment *apps_v1.Deployment) (result *apps_v1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentsResource, c.ns, deployment), &apps_v1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(deployment *apps_v1.Deployment) (result *apps_v1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentsResource, c.ns, deployment), &apps_v1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.Deployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployments) UpdateStatus(deployment *apps_v1.Deployment) (*apps_v1.Deployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentsResource, "status", c.ns, deployment), &apps_v1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentsResource, c.ns, name), &apps_v1.Deployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &apps_v1.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps_v1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, name, data, subresources...), &apps_v1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.Deployment), err
}

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	cluster "github.com/gardener/machine-controller-manager/pkg/apis/cluster"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io2/apimachinery/pkg/apis/meta/v1"
	types "k8s.io2/apimachinery/pkg/types"
	watch "k8s.io2/apimachinery/pkg/watch"
	rest "k8s.io2/client-go/rest"
)

// ClustersGetter has a method to return a ClusterInterface.
// A group's client should implement this interface.
type ClustersGetter interface {
	Clusters(namespace string) ClusterInterface
}

// ClusterInterface has methods to work with Cluster resources.
type ClusterInterface interface {
	Create(*cluster.Cluster) (*cluster.Cluster, error)
	Update(*cluster.Cluster) (*cluster.Cluster, error)
	UpdateStatus(*cluster.Cluster) (*cluster.Cluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*cluster.Cluster, error)
	List(opts v1.ListOptions) (*cluster.ClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cluster.Cluster, err error)
	ClusterExpansion
}

// clusters implements ClusterInterface
type clusters struct {
	client rest.Interface
	ns     string
}

// newClusters returns a Clusters
func newClusters(c *ClusterClient, namespace string) *clusters {
	return &clusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *clusters) Get(name string, options v1.GetOptions) (result *cluster.Cluster, err error) {
	result = &cluster.Cluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *clusters) List(opts v1.ListOptions) (result *cluster.ClusterList, err error) {
	result = &cluster.ClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *clusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *clusters) Create(cluster *cluster.Cluster) (result *cluster.Cluster, err error) {
	result = &cluster.Cluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusters").
		Body(cluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *clusters) Update(cluster *cluster.Cluster) (result *cluster.Cluster, err error) {
	result = &cluster.Cluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusters").
		Name(cluster.Name).
		Body(cluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusters) UpdateStatus(cluster *cluster.Cluster) (result *cluster.Cluster, err error) {
	result = &cluster.Cluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusters").
		Name(cluster.Name).
		SubResource("status").
		Body(cluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *clusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cluster.
func (c *clusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cluster.Cluster, err error) {
	result = &cluster.Cluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

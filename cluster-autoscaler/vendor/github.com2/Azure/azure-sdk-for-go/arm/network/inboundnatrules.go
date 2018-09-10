package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com2/Azure/go-autorest/autorest"
	"github.com2/Azure/go-autorest/autorest/azure"
	"github.com2/Azure/go-autorest/autorest/validation"
	"net/http"
)

// InboundNatRulesClient is the network Client
type InboundNatRulesClient struct {
	ManagementClient
}

// NewInboundNatRulesClient creates an instance of the InboundNatRulesClient client.
func NewInboundNatRulesClient(subscriptionID string) InboundNatRulesClient {
	return NewInboundNatRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInboundNatRulesClientWithBaseURI creates an instance of the InboundNatRulesClient client.
func NewInboundNatRulesClientWithBaseURI(baseURI string, subscriptionID string) InboundNatRulesClient {
	return InboundNatRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a load balancer inbound nat rule. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
// inboundNatRuleName is the name of the inbound nat rule. inboundNatRuleParameters is parameters supplied to the
// create or update inbound nat rule operation.
func (client InboundNatRulesClient) CreateOrUpdate(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, cancel <-chan struct{}) (<-chan InboundNatRule, <-chan error) {
	resultChan := make(chan InboundNatRule, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: inboundNatRuleParameters,
			Constraints: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "inboundNatRuleParameters.InboundNatRulePropertiesFormat.BackendIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
									}},
								}},
							}},
						}},
					}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "network.InboundNatRulesClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result InboundNatRule
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InboundNatRulesClient) CreateOrUpdatePreparer(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"inboundNatRuleName": autorest.Encode("path", inboundNatRuleName),
		"loadBalancerName":   autorest.Encode("path", loadBalancerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", pathParameters),
		autorest.WithJSON(inboundNatRuleParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InboundNatRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InboundNatRulesClient) CreateOrUpdateResponder(resp *http.Response) (result InboundNatRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified load balancer inbound nat rule. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
// inboundNatRuleName is the name of the inbound nat rule.
func (client InboundNatRulesClient) Delete(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, loadBalancerName, inboundNatRuleName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client InboundNatRulesClient) DeletePreparer(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"inboundNatRuleName": autorest.Encode("path", inboundNatRuleName),
		"loadBalancerName":   autorest.Encode("path", loadBalancerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InboundNatRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InboundNatRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified load balancer inbound nat rule.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
// inboundNatRuleName is the name of the inbound nat rule. expand is expands referenced resources.
func (client InboundNatRulesClient) Get(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, expand string) (result InboundNatRule, err error) {
	req, err := client.GetPreparer(resourceGroupName, loadBalancerName, inboundNatRuleName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InboundNatRulesClient) GetPreparer(resourceGroupName string, loadBalancerName string, inboundNatRuleName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"inboundNatRuleName": autorest.Encode("path", inboundNatRuleName),
		"loadBalancerName":   autorest.Encode("path", loadBalancerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InboundNatRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InboundNatRulesClient) GetResponder(resp *http.Response) (result InboundNatRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the inbound nat rules in a load balancer.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
func (client InboundNatRulesClient) List(resourceGroupName string, loadBalancerName string) (result InboundNatRuleListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InboundNatRulesClient) ListPreparer(resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InboundNatRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InboundNatRulesClient) ListResponder(resp *http.Response) (result InboundNatRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client InboundNatRulesClient) ListNextResults(lastResults InboundNatRuleListResult) (result InboundNatRuleListResult, err error) {
	req, err := lastResults.InboundNatRuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundNatRulesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client InboundNatRulesClient) ListComplete(resourceGroupName string, loadBalancerName string, cancel <-chan struct{}) (<-chan InboundNatRule, <-chan error) {
	resultChan := make(chan InboundNatRule)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, loadBalancerName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

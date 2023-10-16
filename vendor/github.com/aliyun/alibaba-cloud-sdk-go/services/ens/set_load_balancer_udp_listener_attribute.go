package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetLoadBalancerUDPListenerAttribute invokes the ens.SetLoadBalancerUDPListenerAttribute API synchronously
func (client *Client) SetLoadBalancerUDPListenerAttribute(request *SetLoadBalancerUDPListenerAttributeRequest) (response *SetLoadBalancerUDPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerUDPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerUDPListenerAttributeWithChan invokes the ens.SetLoadBalancerUDPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerUDPListenerAttributeWithChan(request *SetLoadBalancerUDPListenerAttributeRequest) (<-chan *SetLoadBalancerUDPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerUDPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerUDPListenerAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetLoadBalancerUDPListenerAttributeWithCallback invokes the ens.SetLoadBalancerUDPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerUDPListenerAttributeWithCallback(request *SetLoadBalancerUDPListenerAttributeRequest, callback func(response *SetLoadBalancerUDPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerUDPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerUDPListenerAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetLoadBalancerUDPListenerAttributeRequest is the request struct for api SetLoadBalancerUDPListenerAttribute
type SetLoadBalancerUDPListenerAttributeRequest struct {
	*requests.RpcRequest
	Protocol                  string           `position:"Query" name:"Protocol"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	HealthCheckReq            string           `position:"Query" name:"HealthCheckReq"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"HealthCheckInterval"`
	HealthCheckExp            string           `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
	Description               string           `position:"Query" name:"Description"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	EipTransmit               string           `position:"Query" name:"EipTransmit"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
}

// SetLoadBalancerUDPListenerAttributeResponse is the response struct for api SetLoadBalancerUDPListenerAttribute
type SetLoadBalancerUDPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerUDPListenerAttributeRequest creates a request to invoke SetLoadBalancerUDPListenerAttribute API
func CreateSetLoadBalancerUDPListenerAttributeRequest() (request *SetLoadBalancerUDPListenerAttributeRequest) {
	request = &SetLoadBalancerUDPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "SetLoadBalancerUDPListenerAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLoadBalancerUDPListenerAttributeResponse creates a response to parse from SetLoadBalancerUDPListenerAttribute response
func CreateSetLoadBalancerUDPListenerAttributeResponse() (response *SetLoadBalancerUDPListenerAttributeResponse) {
	response = &SetLoadBalancerUDPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// CreateEipInstance invokes the ens.CreateEipInstance API synchronously
func (client *Client) CreateEipInstance(request *CreateEipInstanceRequest) (response *CreateEipInstanceResponse, err error) {
	response = CreateCreateEipInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEipInstanceWithChan invokes the ens.CreateEipInstance API asynchronously
func (client *Client) CreateEipInstanceWithChan(request *CreateEipInstanceRequest) (<-chan *CreateEipInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateEipInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEipInstance(request)
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

// CreateEipInstanceWithCallback invokes the ens.CreateEipInstance API asynchronously
func (client *Client) CreateEipInstanceWithCallback(request *CreateEipInstanceRequest, callback func(response *CreateEipInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEipInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateEipInstance(request)
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

// CreateEipInstanceRequest is the request struct for api CreateEipInstance
type CreateEipInstanceRequest struct {
	*requests.RpcRequest
	Isp                string           `position:"Query" name:"Isp"`
	EnsRegionId        string           `position:"Query" name:"EnsRegionId"`
	InstanceChargeType string           `position:"Query" name:"InstanceChargeType"`
	Bandwidth          requests.Integer `position:"Query" name:"Bandwidth"`
	InternetChargeType string           `position:"Query" name:"InternetChargeType"`
	Name               string           `position:"Query" name:"Name"`
}

// CreateEipInstanceResponse is the response struct for api CreateEipInstance
type CreateEipInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AllocationId string `json:"AllocationId" xml:"AllocationId"`
}

// CreateCreateEipInstanceRequest creates a request to invoke CreateEipInstance API
func CreateCreateEipInstanceRequest() (request *CreateEipInstanceRequest) {
	request = &CreateEipInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateEipInstance", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEipInstanceResponse creates a response to parse from CreateEipInstance response
func CreateCreateEipInstanceResponse() (response *CreateEipInstanceResponse) {
	response = &CreateEipInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

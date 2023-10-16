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

// CreateSecurityGroup invokes the ens.CreateSecurityGroup API synchronously
func (client *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
	response = CreateCreateSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSecurityGroupWithChan invokes the ens.CreateSecurityGroup API asynchronously
func (client *Client) CreateSecurityGroupWithChan(request *CreateSecurityGroupRequest) (<-chan *CreateSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *CreateSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSecurityGroup(request)
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

// CreateSecurityGroupWithCallback invokes the ens.CreateSecurityGroup API asynchronously
func (client *Client) CreateSecurityGroupWithCallback(request *CreateSecurityGroupRequest, callback func(response *CreateSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateSecurityGroup(request)
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

// CreateSecurityGroupRequest is the request struct for api CreateSecurityGroup
type CreateSecurityGroupRequest struct {
	*requests.RpcRequest
	Description       string `position:"Query" name:"Description"`
	SecurityGroupName string `position:"Query" name:"SecurityGroupName"`
}

// CreateSecurityGroupResponse is the response struct for api CreateSecurityGroup
type CreateSecurityGroupResponse struct {
	*responses.BaseResponse
	SecurityGroupId string `json:"SecurityGroupId" xml:"SecurityGroupId"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateSecurityGroupRequest creates a request to invoke CreateSecurityGroup API
func CreateCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateSecurityGroup", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSecurityGroupResponse creates a response to parse from CreateSecurityGroup response
func CreateCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

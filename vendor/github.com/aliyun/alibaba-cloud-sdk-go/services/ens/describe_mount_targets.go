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

// DescribeMountTargets invokes the ens.DescribeMountTargets API synchronously
func (client *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
	response = CreateDescribeMountTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMountTargetsWithChan invokes the ens.DescribeMountTargets API asynchronously
func (client *Client) DescribeMountTargetsWithChan(request *DescribeMountTargetsRequest) (<-chan *DescribeMountTargetsResponse, <-chan error) {
	responseChan := make(chan *DescribeMountTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMountTargets(request)
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

// DescribeMountTargetsWithCallback invokes the ens.DescribeMountTargets API asynchronously
func (client *Client) DescribeMountTargetsWithCallback(request *DescribeMountTargetsRequest, callback func(response *DescribeMountTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMountTargetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMountTargets(request)
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

// DescribeMountTargetsRequest is the request struct for api DescribeMountTargets
type DescribeMountTargetsRequest struct {
	*requests.RpcRequest
	MountTargetName string           `position:"Query" name:"MountTargetName"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	EnsRegionId     string           `position:"Query" name:"EnsRegionId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	FileSystemId    string           `position:"Query" name:"FileSystemId"`
}

// DescribeMountTargetsResponse is the response struct for api DescribeMountTargets
type DescribeMountTargetsResponse struct {
	*responses.BaseResponse
	PageNumber   int                `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                `json:"PageSize" xml:"PageSize"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	MountTargets []MountTargetsItem `json:"MountTargets" xml:"MountTargets"`
}

// CreateDescribeMountTargetsRequest creates a request to invoke DescribeMountTargets API
func CreateDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
	request = &DescribeMountTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeMountTargets", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMountTargetsResponse creates a response to parse from DescribeMountTargets response
func CreateDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
	response = &DescribeMountTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

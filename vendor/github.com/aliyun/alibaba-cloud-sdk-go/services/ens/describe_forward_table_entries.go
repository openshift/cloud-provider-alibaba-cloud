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

// DescribeForwardTableEntries invokes the ens.DescribeForwardTableEntries API synchronously
func (client *Client) DescribeForwardTableEntries(request *DescribeForwardTableEntriesRequest) (response *DescribeForwardTableEntriesResponse, err error) {
	response = CreateDescribeForwardTableEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeForwardTableEntriesWithChan invokes the ens.DescribeForwardTableEntries API asynchronously
func (client *Client) DescribeForwardTableEntriesWithChan(request *DescribeForwardTableEntriesRequest) (<-chan *DescribeForwardTableEntriesResponse, <-chan error) {
	responseChan := make(chan *DescribeForwardTableEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeForwardTableEntries(request)
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

// DescribeForwardTableEntriesWithCallback invokes the ens.DescribeForwardTableEntries API asynchronously
func (client *Client) DescribeForwardTableEntriesWithCallback(request *DescribeForwardTableEntriesRequest, callback func(response *DescribeForwardTableEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeForwardTableEntriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeForwardTableEntries(request)
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

// DescribeForwardTableEntriesRequest is the request struct for api DescribeForwardTableEntries
type DescribeForwardTableEntriesRequest struct {
	*requests.RpcRequest
	InternalIp       string           `position:"Query" name:"InternalIp"`
	ExternalIp       string           `position:"Query" name:"ExternalIp"`
	IpProtocol       string           `position:"Query" name:"IpProtocol"`
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	ForwardEntryId   string           `position:"Query" name:"ForwardEntryId"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	NatGatewayId     string           `position:"Query" name:"NatGatewayId"`
	ForwardEntryName string           `position:"Query" name:"ForwardEntryName"`
}

// DescribeForwardTableEntriesResponse is the response struct for api DescribeForwardTableEntries
type DescribeForwardTableEntriesResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PageNumber          string              `json:"PageNumber" xml:"PageNumber"`
	TotalCount          string              `json:"TotalCount" xml:"TotalCount"`
	PageSize            string              `json:"PageSize" xml:"PageSize"`
	ForwardTableEntries []ForwardTableEntry `json:"ForwardTableEntries" xml:"ForwardTableEntries"`
}

// CreateDescribeForwardTableEntriesRequest creates a request to invoke DescribeForwardTableEntries API
func CreateDescribeForwardTableEntriesRequest() (request *DescribeForwardTableEntriesRequest) {
	request = &DescribeForwardTableEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeForwardTableEntries", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeForwardTableEntriesResponse creates a response to parse from DescribeForwardTableEntries response
func CreateDescribeForwardTableEntriesResponse() (response *DescribeForwardTableEntriesResponse) {
	response = &DescribeForwardTableEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

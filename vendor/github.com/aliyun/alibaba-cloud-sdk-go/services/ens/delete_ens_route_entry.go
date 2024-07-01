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

// DeleteEnsRouteEntry invokes the ens.DeleteEnsRouteEntry API synchronously
func (client *Client) DeleteEnsRouteEntry(request *DeleteEnsRouteEntryRequest) (response *DeleteEnsRouteEntryResponse, err error) {
	response = CreateDeleteEnsRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEnsRouteEntryWithChan invokes the ens.DeleteEnsRouteEntry API asynchronously
func (client *Client) DeleteEnsRouteEntryWithChan(request *DeleteEnsRouteEntryRequest) (<-chan *DeleteEnsRouteEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteEnsRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEnsRouteEntry(request)
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

// DeleteEnsRouteEntryWithCallback invokes the ens.DeleteEnsRouteEntry API asynchronously
func (client *Client) DeleteEnsRouteEntryWithCallback(request *DeleteEnsRouteEntryRequest, callback func(response *DeleteEnsRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEnsRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteEnsRouteEntry(request)
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

// DeleteEnsRouteEntryRequest is the request struct for api DeleteEnsRouteEntry
type DeleteEnsRouteEntryRequest struct {
	*requests.RpcRequest
	RouteEntryId string `position:"Query" name:"RouteEntryId"`
}

// DeleteEnsRouteEntryResponse is the response struct for api DeleteEnsRouteEntry
type DeleteEnsRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEnsRouteEntryRequest creates a request to invoke DeleteEnsRouteEntry API
func CreateDeleteEnsRouteEntryRequest() (request *DeleteEnsRouteEntryRequest) {
	request = &DeleteEnsRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DeleteEnsRouteEntry", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEnsRouteEntryResponse creates a response to parse from DeleteEnsRouteEntry response
func CreateDeleteEnsRouteEntryResponse() (response *DeleteEnsRouteEntryResponse) {
	response = &DeleteEnsRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
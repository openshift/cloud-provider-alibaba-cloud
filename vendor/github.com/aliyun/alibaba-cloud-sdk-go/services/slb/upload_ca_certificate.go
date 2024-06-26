package slb

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

// UploadCACertificate invokes the slb.UploadCACertificate API synchronously
func (client *Client) UploadCACertificate(request *UploadCACertificateRequest) (response *UploadCACertificateResponse, err error) {
	response = CreateUploadCACertificateResponse()
	err = client.DoAction(request, response)
	return
}

// UploadCACertificateWithChan invokes the slb.UploadCACertificate API asynchronously
func (client *Client) UploadCACertificateWithChan(request *UploadCACertificateRequest) (<-chan *UploadCACertificateResponse, <-chan error) {
	responseChan := make(chan *UploadCACertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadCACertificate(request)
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

// UploadCACertificateWithCallback invokes the slb.UploadCACertificate API asynchronously
func (client *Client) UploadCACertificateWithCallback(request *UploadCACertificateRequest, callback func(response *UploadCACertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadCACertificateResponse
		var err error
		defer close(result)
		response, err = client.UploadCACertificate(request)
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

// UploadCACertificateRequest is the request struct for api UploadCACertificate
type UploadCACertificateRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	CACertificateName    string           `position:"Query" name:"CACertificateName"`
	CACertificate        string           `position:"Query" name:"CACertificate"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	StandardType         string           `position:"Query" name:"StandardType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UploadCACertificateResponse is the response struct for api UploadCACertificate
type UploadCACertificateResponse struct {
	*responses.BaseResponse
	CreateTimeStamp   int64  `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ExpireTime        string `json:"ExpireTime" xml:"ExpireTime"`
	Fingerprint       string `json:"Fingerprint" xml:"Fingerprint"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	CommonName        string `json:"CommonName" xml:"CommonName"`
	ResourceGroupId   string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	CACertificateName string `json:"CACertificateName" xml:"CACertificateName"`
	ExpireTimeStamp   int64  `json:"ExpireTimeStamp" xml:"ExpireTimeStamp"`
	CACertificateId   string `json:"CACertificateId" xml:"CACertificateId"`
}

// CreateUploadCACertificateRequest creates a request to invoke UploadCACertificate API
func CreateUploadCACertificateRequest() (request *UploadCACertificateRequest) {
	request = &UploadCACertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "UploadCACertificate", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUploadCACertificateResponse creates a response to parse from UploadCACertificate response
func CreateUploadCACertificateResponse() (response *UploadCACertificateResponse) {
	response = &UploadCACertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

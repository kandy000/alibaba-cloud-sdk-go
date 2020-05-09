package ddoscoo

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

// DescribleCertList invokes the ddoscoo.DescribleCertList API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describlecertlist.html
func (client *Client) DescribleCertList(request *DescribleCertListRequest) (response *DescribleCertListResponse, err error) {
	response = CreateDescribleCertListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribleCertListWithChan invokes the ddoscoo.DescribleCertList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describlecertlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribleCertListWithChan(request *DescribleCertListRequest) (<-chan *DescribleCertListResponse, <-chan error) {
	responseChan := make(chan *DescribleCertListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribleCertList(request)
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

// DescribleCertListWithCallback invokes the ddoscoo.DescribleCertList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describlecertlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribleCertListWithCallback(request *DescribleCertListRequest, callback func(response *DescribleCertListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribleCertListResponse
		var err error
		defer close(result)
		response, err = client.DescribleCertList(request)
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

// DescribleCertListRequest is the request struct for api DescribleCertList
type DescribleCertListRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// DescribleCertListResponse is the response struct for api DescribleCertList
type DescribleCertListResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	CertList  []CertItem `json:"CertList" xml:"CertList"`
}

// CreateDescribleCertListRequest creates a request to invoke DescribleCertList API
func CreateDescribleCertListRequest() (request *DescribleCertListRequest) {
	request = &DescribleCertListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DescribleCertList", "ddoscoo", "openAPI")
	return
}

// CreateDescribleCertListResponse creates a response to parse from DescribleCertList response
func CreateDescribleCertListResponse() (response *DescribleCertListResponse) {
	response = &DescribleCertListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

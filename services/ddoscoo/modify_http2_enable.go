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

// ModifyHttp2Enable invokes the ddoscoo.ModifyHttp2Enable API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyhttp2enable.html
func (client *Client) ModifyHttp2Enable(request *ModifyHttp2EnableRequest) (response *ModifyHttp2EnableResponse, err error) {
	response = CreateModifyHttp2EnableResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyHttp2EnableWithChan invokes the ddoscoo.ModifyHttp2Enable API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyhttp2enable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHttp2EnableWithChan(request *ModifyHttp2EnableRequest) (<-chan *ModifyHttp2EnableResponse, <-chan error) {
	responseChan := make(chan *ModifyHttp2EnableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHttp2Enable(request)
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

// ModifyHttp2EnableWithCallback invokes the ddoscoo.ModifyHttp2Enable API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifyhttp2enable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHttp2EnableWithCallback(request *ModifyHttp2EnableRequest, callback func(response *ModifyHttp2EnableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHttp2EnableResponse
		var err error
		defer close(result)
		response, err = client.ModifyHttp2Enable(request)
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

// ModifyHttp2EnableRequest is the request struct for api ModifyHttp2Enable
type ModifyHttp2EnableRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Enable          requests.Integer `position:"Query" name:"Enable"`
	Domain          string           `position:"Query" name:"Domain"`
}

// ModifyHttp2EnableResponse is the response struct for api ModifyHttp2Enable
type ModifyHttp2EnableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyHttp2EnableRequest creates a request to invoke ModifyHttp2Enable API
func CreateModifyHttp2EnableRequest() (request *ModifyHttp2EnableRequest) {
	request = &ModifyHttp2EnableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyHttp2Enable", "ddoscoo", "openAPI")
	return
}

// CreateModifyHttp2EnableResponse creates a response to parse from ModifyHttp2Enable response
func CreateModifyHttp2EnableResponse() (response *ModifyHttp2EnableResponse) {
	response = &ModifyHttp2EnableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

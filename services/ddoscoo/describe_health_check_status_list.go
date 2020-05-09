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

// DescribeHealthCheckStatusList invokes the ddoscoo.DescribeHealthCheckStatusList API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describehealthcheckstatuslist.html
func (client *Client) DescribeHealthCheckStatusList(request *DescribeHealthCheckStatusListRequest) (response *DescribeHealthCheckStatusListResponse, err error) {
	response = CreateDescribeHealthCheckStatusListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHealthCheckStatusListWithChan invokes the ddoscoo.DescribeHealthCheckStatusList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describehealthcheckstatuslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHealthCheckStatusListWithChan(request *DescribeHealthCheckStatusListRequest) (<-chan *DescribeHealthCheckStatusListResponse, <-chan error) {
	responseChan := make(chan *DescribeHealthCheckStatusListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHealthCheckStatusList(request)
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

// DescribeHealthCheckStatusListWithCallback invokes the ddoscoo.DescribeHealthCheckStatusList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describehealthcheckstatuslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHealthCheckStatusListWithCallback(request *DescribeHealthCheckStatusListRequest, callback func(response *DescribeHealthCheckStatusListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHealthCheckStatusListResponse
		var err error
		defer close(result)
		response, err = client.DescribeHealthCheckStatusList(request)
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

// DescribeHealthCheckStatusListRequest is the request struct for api DescribeHealthCheckStatusList
type DescribeHealthCheckStatusListRequest struct {
	*requests.RpcRequest
	Listeners string `position:"Query" name:"Listeners"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// DescribeHealthCheckStatusListResponse is the response struct for api DescribeHealthCheckStatusList
type DescribeHealthCheckStatusListResponse struct {
	*responses.BaseResponse
	RequestId             string              `json:"RequestId" xml:"RequestId"`
	HealthCheckStatusList []HealthCheckStatus `json:"HealthCheckStatusList" xml:"HealthCheckStatusList"`
}

// CreateDescribeHealthCheckStatusListRequest creates a request to invoke DescribeHealthCheckStatusList API
func CreateDescribeHealthCheckStatusListRequest() (request *DescribeHealthCheckStatusListRequest) {
	request = &DescribeHealthCheckStatusListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DescribeHealthCheckStatusList", "ddoscoo", "openAPI")
	return
}

// CreateDescribeHealthCheckStatusListResponse creates a response to parse from DescribeHealthCheckStatusList response
func CreateDescribeHealthCheckStatusListResponse() (response *DescribeHealthCheckStatusListResponse) {
	response = &DescribeHealthCheckStatusListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package videosearch

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

// GetTaskStatus invokes the videosearch.GetTaskStatus API synchronously
// api document: https://help.aliyun.com/api/videosearch/gettaskstatus.html
func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
	response = CreateGetTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskStatusWithChan invokes the videosearch.GetTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/videosearch/gettaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTaskStatusWithChan(request *GetTaskStatusRequest) (<-chan *GetTaskStatusResponse, <-chan error) {
	responseChan := make(chan *GetTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskStatus(request)
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

// GetTaskStatusWithCallback invokes the videosearch.GetTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/videosearch/gettaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTaskStatusWithCallback(request *GetTaskStatusRequest, callback func(response *GetTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.GetTaskStatus(request)
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

// GetTaskStatusRequest is the request struct for api GetTaskStatus
type GetTaskStatusRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	InstanceId  string `position:"Body" name:"InstanceId"`
	TaskId      string `position:"Body" name:"TaskId"`
}

// GetTaskStatusResponse is the response struct for api GetTaskStatus
type GetTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int    `json:"Data" xml:"Data"`
}

// CreateGetTaskStatusRequest creates a request to invoke GetTaskStatus API
func CreateGetTaskStatusRequest() (request *GetTaskStatusRequest) {
	request = &GetTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videosearch", "2020-02-25", "GetTaskStatus", "", "")
	return
}

// CreateGetTaskStatusResponse creates a response to parse from GetTaskStatus response
func CreateGetTaskStatusResponse() (response *GetTaskStatusResponse) {
	response = &GetTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

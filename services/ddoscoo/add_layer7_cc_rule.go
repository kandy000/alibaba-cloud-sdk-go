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

// AddLayer7CCRule invokes the ddoscoo.AddLayer7CCRule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/addlayer7ccrule.html
func (client *Client) AddLayer7CCRule(request *AddLayer7CCRuleRequest) (response *AddLayer7CCRuleResponse, err error) {
	response = CreateAddLayer7CCRuleResponse()
	err = client.DoAction(request, response)
	return
}

// AddLayer7CCRuleWithChan invokes the ddoscoo.AddLayer7CCRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/addlayer7ccrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLayer7CCRuleWithChan(request *AddLayer7CCRuleRequest) (<-chan *AddLayer7CCRuleResponse, <-chan error) {
	responseChan := make(chan *AddLayer7CCRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLayer7CCRule(request)
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

// AddLayer7CCRuleWithCallback invokes the ddoscoo.AddLayer7CCRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/addlayer7ccrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLayer7CCRuleWithCallback(request *AddLayer7CCRuleRequest, callback func(response *AddLayer7CCRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLayer7CCRuleResponse
		var err error
		defer close(result)
		response, err = client.AddLayer7CCRule(request)
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

// AddLayer7CCRuleRequest is the request struct for api AddLayer7CCRule
type AddLayer7CCRuleRequest struct {
	*requests.RpcRequest
	Mode            string           `position:"Query" name:"Mode"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Act             string           `position:"Query" name:"Act"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Count           requests.Integer `position:"Query" name:"Count"`
	Ttl             requests.Integer `position:"Query" name:"Ttl"`
	Uri             string           `position:"Query" name:"Uri"`
	Domain          string           `position:"Query" name:"Domain"`
	Name            string           `position:"Query" name:"Name"`
	Interval        requests.Integer `position:"Query" name:"Interval"`
}

// AddLayer7CCRuleResponse is the response struct for api AddLayer7CCRule
type AddLayer7CCRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLayer7CCRuleRequest creates a request to invoke AddLayer7CCRule API
func CreateAddLayer7CCRuleRequest() (request *AddLayer7CCRuleRequest) {
	request = &AddLayer7CCRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "AddLayer7CCRule", "ddoscoo", "openAPI")
	return
}

// CreateAddLayer7CCRuleResponse creates a response to parse from AddLayer7CCRule response
func CreateAddLayer7CCRuleResponse() (response *AddLayer7CCRuleResponse) {
	response = &AddLayer7CCRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

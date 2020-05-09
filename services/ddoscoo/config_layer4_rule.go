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

// ConfigLayer4Rule invokes the ddoscoo.ConfigLayer4Rule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/configlayer4rule.html
func (client *Client) ConfigLayer4Rule(request *ConfigLayer4RuleRequest) (response *ConfigLayer4RuleResponse, err error) {
	response = CreateConfigLayer4RuleResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigLayer4RuleWithChan invokes the ddoscoo.ConfigLayer4Rule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/configlayer4rule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigLayer4RuleWithChan(request *ConfigLayer4RuleRequest) (<-chan *ConfigLayer4RuleResponse, <-chan error) {
	responseChan := make(chan *ConfigLayer4RuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigLayer4Rule(request)
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

// ConfigLayer4RuleWithCallback invokes the ddoscoo.ConfigLayer4Rule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/configlayer4rule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigLayer4RuleWithCallback(request *ConfigLayer4RuleRequest, callback func(response *ConfigLayer4RuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigLayer4RuleResponse
		var err error
		defer close(result)
		response, err = client.ConfigLayer4Rule(request)
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

// ConfigLayer4RuleRequest is the request struct for api ConfigLayer4Rule
type ConfigLayer4RuleRequest struct {
	*requests.RpcRequest
	Listeners string `position:"Query" name:"Listeners"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// ConfigLayer4RuleResponse is the response struct for api ConfigLayer4Rule
type ConfigLayer4RuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigLayer4RuleRequest creates a request to invoke ConfigLayer4Rule API
func CreateConfigLayer4RuleRequest() (request *ConfigLayer4RuleRequest) {
	request = &ConfigLayer4RuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "ConfigLayer4Rule", "ddoscoo", "openAPI")
	return
}

// CreateConfigLayer4RuleResponse creates a response to parse from ConfigLayer4Rule response
func CreateConfigLayer4RuleResponse() (response *ConfigLayer4RuleResponse) {
	response = &ConfigLayer4RuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

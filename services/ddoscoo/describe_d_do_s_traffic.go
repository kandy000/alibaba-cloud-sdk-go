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

// DescribeDDoSTraffic invokes the ddoscoo.DescribeDDoSTraffic API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddostraffic.html
func (client *Client) DescribeDDoSTraffic(request *DescribeDDoSTrafficRequest) (response *DescribeDDoSTrafficResponse, err error) {
	response = CreateDescribeDDoSTrafficResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDDoSTrafficWithChan invokes the ddoscoo.DescribeDDoSTraffic API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddostraffic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDDoSTrafficWithChan(request *DescribeDDoSTrafficRequest) (<-chan *DescribeDDoSTrafficResponse, <-chan error) {
	responseChan := make(chan *DescribeDDoSTrafficResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDDoSTraffic(request)
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

// DescribeDDoSTrafficWithCallback invokes the ddoscoo.DescribeDDoSTraffic API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddostraffic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDDoSTrafficWithCallback(request *DescribeDDoSTrafficRequest, callback func(response *DescribeDDoSTrafficResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDDoSTrafficResponse
		var err error
		defer close(result)
		response, err = client.DescribeDDoSTraffic(request)
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

// DescribeDDoSTrafficRequest is the request struct for api DescribeDDoSTraffic
type DescribeDDoSTrafficRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	Eip             string           `position:"Query" name:"Eip"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	Interval        requests.Integer `position:"Query" name:"Interval"`
}

// DescribeDDoSTrafficResponse is the response struct for api DescribeDDoSTraffic
type DescribeDDoSTrafficResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	DefenseInBytes    int64              `json:"DefenseInBytes" xml:"DefenseInBytes"`
	SourceInBytes     int64              `json:"SourceInBytes" xml:"SourceInBytes"`
	DDoSTrafficPoints []DDoSTrafficPoint `json:"DDoSTrafficPoints" xml:"DDoSTrafficPoints"`
}

// CreateDescribeDDoSTrafficRequest creates a request to invoke DescribeDDoSTraffic API
func CreateDescribeDDoSTrafficRequest() (request *DescribeDDoSTrafficRequest) {
	request = &DescribeDDoSTrafficRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DescribeDDoSTraffic", "ddoscoo", "openAPI")
	return
}

// CreateDescribeDDoSTrafficResponse creates a response to parse from DescribeDDoSTraffic response
func CreateDescribeDDoSTrafficResponse() (response *DescribeDDoSTrafficResponse) {
	response = &DescribeDDoSTrafficResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

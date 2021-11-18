package smartag

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

// ModifySagUserDns invokes the smartag.ModifySagUserDns API synchronously
func (client *Client) ModifySagUserDns(request *ModifySagUserDnsRequest) (response *ModifySagUserDnsResponse, err error) {
	response = CreateModifySagUserDnsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagUserDnsWithChan invokes the smartag.ModifySagUserDns API asynchronously
func (client *Client) ModifySagUserDnsWithChan(request *ModifySagUserDnsRequest) (<-chan *ModifySagUserDnsResponse, <-chan error) {
	responseChan := make(chan *ModifySagUserDnsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagUserDns(request)
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

// ModifySagUserDnsWithCallback invokes the smartag.ModifySagUserDns API asynchronously
func (client *Client) ModifySagUserDnsWithCallback(request *ModifySagUserDnsRequest, callback func(response *ModifySagUserDnsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagUserDnsResponse
		var err error
		defer close(result)
		response, err = client.ModifySagUserDns(request)
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

// ModifySagUserDnsRequest is the request struct for api ModifySagUserDns
type ModifySagUserDnsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	SlaveDns             string           `position:"Query"`
	MasterDns            string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	SmartAGId            string           `position:"Query"`
	SmartAGSn            string           `position:"Query"`
}

// ModifySagUserDnsResponse is the response struct for api ModifySagUserDns
type ModifySagUserDnsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagUserDnsRequest creates a request to invoke ModifySagUserDns API
func CreateModifySagUserDnsRequest() (request *ModifySagUserDnsRequest) {
	request = &ModifySagUserDnsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagUserDns", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagUserDnsResponse creates a response to parse from ModifySagUserDns response
func CreateModifySagUserDnsResponse() (response *ModifySagUserDnsResponse) {
	response = &ModifySagUserDnsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
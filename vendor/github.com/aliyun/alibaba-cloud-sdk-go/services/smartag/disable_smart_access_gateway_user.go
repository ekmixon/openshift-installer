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

// DisableSmartAccessGatewayUser invokes the smartag.DisableSmartAccessGatewayUser API synchronously
func (client *Client) DisableSmartAccessGatewayUser(request *DisableSmartAccessGatewayUserRequest) (response *DisableSmartAccessGatewayUserResponse, err error) {
	response = CreateDisableSmartAccessGatewayUserResponse()
	err = client.DoAction(request, response)
	return
}

// DisableSmartAccessGatewayUserWithChan invokes the smartag.DisableSmartAccessGatewayUser API asynchronously
func (client *Client) DisableSmartAccessGatewayUserWithChan(request *DisableSmartAccessGatewayUserRequest) (<-chan *DisableSmartAccessGatewayUserResponse, <-chan error) {
	responseChan := make(chan *DisableSmartAccessGatewayUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableSmartAccessGatewayUser(request)
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

// DisableSmartAccessGatewayUserWithCallback invokes the smartag.DisableSmartAccessGatewayUser API asynchronously
func (client *Client) DisableSmartAccessGatewayUserWithCallback(request *DisableSmartAccessGatewayUserRequest, callback func(response *DisableSmartAccessGatewayUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableSmartAccessGatewayUserResponse
		var err error
		defer close(result)
		response, err = client.DisableSmartAccessGatewayUser(request)
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

// DisableSmartAccessGatewayUserRequest is the request struct for api DisableSmartAccessGatewayUser
type DisableSmartAccessGatewayUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	SmartAGId            string           `position:"Query"`
	UserName             string           `position:"Query"`
}

// DisableSmartAccessGatewayUserResponse is the response struct for api DisableSmartAccessGatewayUser
type DisableSmartAccessGatewayUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableSmartAccessGatewayUserRequest creates a request to invoke DisableSmartAccessGatewayUser API
func CreateDisableSmartAccessGatewayUserRequest() (request *DisableSmartAccessGatewayUserRequest) {
	request = &DisableSmartAccessGatewayUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DisableSmartAccessGatewayUser", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableSmartAccessGatewayUserResponse creates a response to parse from DisableSmartAccessGatewayUser response
func CreateDisableSmartAccessGatewayUserResponse() (response *DisableSmartAccessGatewayUserResponse) {
	response = &DisableSmartAccessGatewayUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
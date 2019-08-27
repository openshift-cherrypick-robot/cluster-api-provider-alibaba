package vpc

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

// UpdateNatGatewayNatType invokes the vpc.UpdateNatGatewayNatType API synchronously
func (client *Client) UpdateNatGatewayNatType(request *UpdateNatGatewayNatTypeRequest) (response *UpdateNatGatewayNatTypeResponse, err error) {
	response = CreateUpdateNatGatewayNatTypeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNatGatewayNatTypeWithChan invokes the vpc.UpdateNatGatewayNatType API asynchronously
func (client *Client) UpdateNatGatewayNatTypeWithChan(request *UpdateNatGatewayNatTypeRequest) (<-chan *UpdateNatGatewayNatTypeResponse, <-chan error) {
	responseChan := make(chan *UpdateNatGatewayNatTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNatGatewayNatType(request)
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

// UpdateNatGatewayNatTypeWithCallback invokes the vpc.UpdateNatGatewayNatType API asynchronously
func (client *Client) UpdateNatGatewayNatTypeWithCallback(request *UpdateNatGatewayNatTypeRequest, callback func(response *UpdateNatGatewayNatTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNatGatewayNatTypeResponse
		var err error
		defer close(result)
		response, err = client.UpdateNatGatewayNatType(request)
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

// UpdateNatGatewayNatTypeRequest is the request struct for api UpdateNatGatewayNatType
type UpdateNatGatewayNatTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	NatType              string           `position:"Query" name:"NatType"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
}

// UpdateNatGatewayNatTypeResponse is the response struct for api UpdateNatGatewayNatType
type UpdateNatGatewayNatTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateNatGatewayNatTypeRequest creates a request to invoke UpdateNatGatewayNatType API
func CreateUpdateNatGatewayNatTypeRequest() (request *UpdateNatGatewayNatTypeRequest) {
	request = &UpdateNatGatewayNatTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UpdateNatGatewayNatType", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateNatGatewayNatTypeResponse creates a response to parse from UpdateNatGatewayNatType response
func CreateUpdateNatGatewayNatTypeResponse() (response *UpdateNatGatewayNatTypeResponse) {
	response = &UpdateNatGatewayNatTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

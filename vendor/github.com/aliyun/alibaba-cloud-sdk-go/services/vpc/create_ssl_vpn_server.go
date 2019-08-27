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

// CreateSslVpnServer invokes the vpc.CreateSslVpnServer API synchronously
func (client *Client) CreateSslVpnServer(request *CreateSslVpnServerRequest) (response *CreateSslVpnServerResponse, err error) {
	response = CreateCreateSslVpnServerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSslVpnServerWithChan invokes the vpc.CreateSslVpnServer API asynchronously
func (client *Client) CreateSslVpnServerWithChan(request *CreateSslVpnServerRequest) (<-chan *CreateSslVpnServerResponse, <-chan error) {
	responseChan := make(chan *CreateSslVpnServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSslVpnServer(request)
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

// CreateSslVpnServerWithCallback invokes the vpc.CreateSslVpnServer API asynchronously
func (client *Client) CreateSslVpnServerWithCallback(request *CreateSslVpnServerRequest, callback func(response *CreateSslVpnServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSslVpnServerResponse
		var err error
		defer close(result)
		response, err = client.CreateSslVpnServer(request)
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

// CreateSslVpnServerRequest is the request struct for api CreateSslVpnServer
type CreateSslVpnServerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	LocalSubnet           string           `position:"Query" name:"LocalSubnet"`
	IDaaSRegionId         string           `position:"Query" name:"IDaaSRegionId"`
	EnableMultiFactorAuth requests.Boolean `position:"Query" name:"EnableMultiFactorAuth"`
	IDaaSInstanceId       string           `position:"Query" name:"IDaaSInstanceId"`
	Cipher                string           `position:"Query" name:"Cipher"`
	ClientIpPool          string           `position:"Query" name:"ClientIpPool"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	Compress              requests.Boolean `position:"Query" name:"Compress"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	VpnGatewayId          string           `position:"Query" name:"VpnGatewayId"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Port                  requests.Integer `position:"Query" name:"Port"`
	Proto                 string           `position:"Query" name:"Proto"`
	Name                  string           `position:"Query" name:"Name"`
}

// CreateSslVpnServerResponse is the response struct for api CreateSslVpnServer
type CreateSslVpnServerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	SslVpnServerId string `json:"SslVpnServerId" xml:"SslVpnServerId"`
	Name           string `json:"Name" xml:"Name"`
}

// CreateCreateSslVpnServerRequest creates a request to invoke CreateSslVpnServer API
func CreateCreateSslVpnServerRequest() (request *CreateSslVpnServerRequest) {
	request = &CreateSslVpnServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateSslVpnServer", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSslVpnServerResponse creates a response to parse from CreateSslVpnServer response
func CreateCreateSslVpnServerResponse() (response *CreateSslVpnServerResponse) {
	response = &CreateSslVpnServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

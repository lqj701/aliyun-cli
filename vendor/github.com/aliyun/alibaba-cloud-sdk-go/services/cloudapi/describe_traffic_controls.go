package cloudapi

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

// DescribeTrafficControls invokes the cloudapi.DescribeTrafficControls API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrols.html
func (client *Client) DescribeTrafficControls(request *DescribeTrafficControlsRequest) (response *DescribeTrafficControlsResponse, err error) {
	response = CreateDescribeTrafficControlsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrafficControlsWithChan invokes the cloudapi.DescribeTrafficControls API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrols.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTrafficControlsWithChan(request *DescribeTrafficControlsRequest) (<-chan *DescribeTrafficControlsResponse, <-chan error) {
	responseChan := make(chan *DescribeTrafficControlsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrafficControls(request)
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

// DescribeTrafficControlsWithCallback invokes the cloudapi.DescribeTrafficControls API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrols.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTrafficControlsWithCallback(request *DescribeTrafficControlsRequest, callback func(response *DescribeTrafficControlsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrafficControlsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrafficControls(request)
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

// DescribeTrafficControlsRequest is the request struct for api DescribeTrafficControls
type DescribeTrafficControlsRequest struct {
	*requests.RpcRequest
	TrafficControlId   string           `position:"Query" name:"TrafficControlId"`
	GroupId            string           `position:"Query" name:"GroupId"`
	ApiId              string           `position:"Query" name:"ApiId"`
	StageName          string           `position:"Query" name:"StageName"`
	TrafficControlName string           `position:"Query" name:"TrafficControlName"`
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeTrafficControlsResponse is the response struct for api DescribeTrafficControls
type DescribeTrafficControlsResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	TrafficControls TrafficControls `json:"TrafficControls" xml:"TrafficControls"`
}

// CreateDescribeTrafficControlsRequest creates a request to invoke DescribeTrafficControls API
func CreateDescribeTrafficControlsRequest() (request *DescribeTrafficControlsRequest) {
	request = &DescribeTrafficControlsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeTrafficControls", "apigateway", "openAPI")
	return
}

// CreateDescribeTrafficControlsResponse creates a response to parse from DescribeTrafficControls response
func CreateDescribeTrafficControlsResponse() (response *DescribeTrafficControlsResponse) {
	response = &DescribeTrafficControlsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
